package osm

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/paulmach/osm"
	"io"
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

	MainOverpassInterpreter = "https://overpass-api.de/api/interpreter"
	//KumiSystemsOverpassInterpreter = "https://overpass.kumi.systems/api/interpreter"

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
			Lat:   node.Lat,
			Lng:   node.Lon,
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

		var lat, lng float64
		for _, node := range nodes {
			lat += node.Lat
			lng += node.Lon
		}
		lat /= float64(len(nodes))
		lng /= float64(len(nodes))

		kebabShops[osmID] = ent.KebabShop{
			OsmID: &osmID,
			Name:  way.Tags.Find("name"),
			Lat:   lat,
			Lng:   lng,
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

		var lat, lng float64
		for _, node := range nodes {
			lat += node.Lat
			lng += node.Lon
		}
		lat /= float64(len(nodes))
		lng /= float64(len(nodes))

		kebabShops[osmID] = ent.KebabShop{
			OsmID: &osmID,
			Name:  rel.Tags.Find("name"),
			Lat:   lat,
			Lng:   lng,
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
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

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
