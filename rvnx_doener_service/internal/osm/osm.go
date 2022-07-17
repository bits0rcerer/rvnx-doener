package osm

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/jackc/pgtype"
	"github.com/paulmach/osm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"rvnx_doener_service/ent"
	"strconv"
	"strings"
)

const (
	EnableOSMRequestCachingKey = "OSM_REQUEST_CACHING"

	MainOverpassInterpreter        = "https://overpass-api.de/api/interpreter"
	KumiSystemsOverpassInterpreter = "https://overpass.kumi.systems/api/interpreter"

	OverpassKebabQuery = `
	[out:xml][timeout:600][maxsize:33554432];
	(
	  node["cuisine"="kebab"];
	  way["cuisine"="kebab"];
	  relation["cuisine"="kebab"];
	);
	>>->.x;
	out;
	.x out;
	`
)

func GetOSMKebabShops() ([]ent.KebabShop, error) {
	const interpreter = MainOverpassInterpreter

	log.Println("[*] Requesting kebab shops from OpenStreetMap via " + interpreter)
	osmData, err := doOSMRequest(interpreter, OverpassKebabQuery)
	if err != nil {
		return nil, err
	}

	nodeMap := make(map[osm.NodeID]*osm.Node)
	kebabShops := make(map[int]ent.KebabShop)
	for _, node := range osmData.Nodes {
		nodeMap[node.ID] = node

		if node.Tags.Find("name") == "" {
			continue
		}

		osmID := int(node.ID)

		kebabShops[osmID] = ent.KebabShop{
			OsmID: &osmID,
			Name:  node.Tags.Find("name"),
			Point: &pgtype.Point{
				P: pgtype.Vec2{
					X: node.Lat,
					Y: node.Lon,
				},
				Status: pgtype.Present,
			},
		}
	}

	wayMap := make(map[osm.WayID]*osm.Way)
	for _, way := range osmData.Ways {
		wayMap[way.ID] = way

		if way.Tags.Find("name") == "" {
			continue
		}

		osmID := int(way.ID)

		var nodes osm.Nodes
		for _, node := range way.Nodes {
			nodes = append(nodes, nodeMap[node.ID])
		}

		p := pgtype.Point{Status: pgtype.Present}
		for _, node := range nodes {
			p.P.X += node.Lat
			p.P.Y += node.Lon
		}
		p.P.X /= float64(len(nodes))
		p.P.Y /= float64(len(nodes))

		kebabShops[osmID] = ent.KebabShop{
			OsmID: &osmID,
			Name:  way.Tags.Find("name"),
			Point: &p,
		}
	}

	for _, rel := range osmData.Relations {
		if rel.Tags.Find("name") == "" {
			continue
		}

		osmID := int(rel.ID)

		var nodes []*osm.Node
		for _, m := range rel.Members {
			switch m.Type {
			case "way":
				for _, node := range wayMap[osm.WayID(m.Ref)].Nodes {
					nodes = append(nodes, nodeMap[node.ID])
				}
			case "node":
				nodes = append(nodes, nodeMap[osm.NodeID(m.Ref)])
			default:
				// ignore
			}
		}

		p := pgtype.Point{Status: pgtype.Present}
		for _, node := range nodes {
			p.P.X += node.Lat
			p.P.Y += node.Lon
		}
		p.P.X /= float64(len(nodes))
		p.P.Y /= float64(len(nodes))

		kebabShops[osmID] = ent.KebabShop{
			OsmID: &osmID,
			Name:  rel.Tags.Find("name"),
			Point: &p,
		}
	}

	res := make([]ent.KebabShop, len(kebabShops))
	i := 0
	for _, shop := range kebabShops {
		res[i] = shop
		i++
	}

	return res, nil
}

func requestElements(interpreter, elementKey string, ids []int64) (*osm.OSM, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	log.Println("[*] Requesting " + strconv.Itoa(len(ids)) + " " + elementKey + "(s) from OpenStreetMap via " + interpreter)

	queryBuilder := strings.Builder{}
	queryBuilder.WriteString(
		"[out:xml][timeout:600][maxsize:33554432];\n" +
			"(\n")

	for _, id := range ids {
		queryBuilder.WriteString(elementKey)
		queryBuilder.WriteString("(")
		queryBuilder.WriteString(strconv.FormatInt(id, 10))
		queryBuilder.WriteString(");\n")
	}

	queryBuilder.WriteString(
		");\n" +
			"out body;")

	osmData, err := doOSMRequest(interpreter, queryBuilder.String())
	if err != nil {
		return nil, err
	}

	return osmData, nil
}

var osmCounter = 0

func doOSMRequest(interpreter, query string) (*osm.OSM, error) {
	caching := strings.ToLower(os.Getenv(EnableOSMRequestCachingKey)) == "true"

	var body []byte
	var err error
	if caching {
		body, err = ioutil.ReadFile("/tmp/osmData" + strconv.Itoa(osmCounter))
	}

	if !caching || err != nil {
		resp, err := http.Post(interpreter, "text/plain", bytes.NewReader([]byte(query)))
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, errors.New(interpreter + " returned " + strconv.Itoa(resp.StatusCode))
		}

		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		_ = ioutil.WriteFile("/tmp/osmData"+strconv.Itoa(osmCounter), body, 0666)
	}
	osmCounter++

	var osmData osm.OSM
	err = xml.Unmarshal(body, &osmData)
	if err != nil {
		return nil, err
	}

	return &osmData, nil
}
