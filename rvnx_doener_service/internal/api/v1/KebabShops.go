package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/internal/services"
	"strconv"
)

const defaultClusterCount = 25
const maxClusterCount = 50

const defaultClusterThreshold = 100
const maxClusterThreshold = 1000

const defaultMinNorm = 0.2

func clusterNormScaling(linearNorm float64) float64 {
	return (1-defaultMinNorm)*math.Pow(linearNorm, 2) + defaultMinNorm
}

func RouteKebabShops(r *gin.RouterGroup, env *services.ServiceEnvironment) {
	r.GET("/box", getBoundingBoxHandler(env.KebabShopService))
	r.GET("/clusters", getClustersHandler(env.KebabShopService))
	r.GET("/auto", getAutoHandler(env.KebabShopService))
	r.GET("/:shop_id", getShopByID(env.KebabShopService))
}

type boundingBox struct {
	latMin, latMax, lngMin, lngMax float64
}

func getShopsInBox(c *gin.Context, service *services.KebabShopService) ([]*ent.KebabShop, *boundingBox, bool) {
	box := boundingBox{}
	var err error

	box.latMin, err = strconv.ParseFloat(c.Query("ltm"), 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return nil, nil, true
	}
	box.latMax, err = strconv.ParseFloat(c.Query("ltx"), 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return nil, nil, true
	}
	box.lngMin, err = strconv.ParseFloat(c.Query("lnm"), 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return nil, nil, true
	}
	box.lngMax, err = strconv.ParseFloat(c.Query("lnx"), 64)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return nil, nil, true
	}

	shops, err := service.Within(box.latMin, box.latMax, box.lngMin, box.lngMax, "id", "lat", "lng")
	if err != nil {
		log.Panic(err)
	}

	return shops, &box, false
}

func createShopResponse(c *gin.Context, shops []*ent.KebabShop) {
	shopsJSON := make([]gin.H, len(shops))
	for i := 0; i < len(shopsJSON); i++ {
		shopsJSON[i] = gin.H{
			"id":  strconv.Itoa(shops[i].ID),
			"lat": shops[i].Lat,
			"lng": shops[i].Lng,
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"cords": shopsJSON,
	})
}

func getBoundingBoxHandler(service *services.KebabShopService) func(c *gin.Context) {
	return func(c *gin.Context) {
		shops, _, abort := getShopsInBox(c, service)
		if abort {
			return
		}
		createShopResponse(c, shops)
	}
}

type cluster struct {
	MaxLat, MinLat, Lat float64
	MaxLng, MinLng, Lng float64
	norm                float64
	ShopsCount          int
}

func createClusteredResponse(c *gin.Context, shops []*ent.KebabShop, box *boundingBox) {
	clusterCount, err := strconv.ParseInt(c.Query("cc"), 10, 32)
	if err != nil ||
		clusterCount < 1 ||
		clusterCount > maxClusterCount {
		clusterCount = defaultClusterCount
	}

	clusterWidth := (box.latMax - box.latMin) / float64(clusterCount)
	clusterHeight := (box.lngMax - box.lngMin) / float64(clusterCount)

	clusters := make([][]cluster, clusterCount)
	for i := 0; i < int(clusterCount); i++ {
		clusters[i] = make([]cluster, clusterCount)

		for j := 0; j < int(clusterCount); j++ {
			clusters[i][j].MaxLat = -math.MaxFloat64
			clusters[i][j].MaxLng = -math.MaxFloat64
			clusters[i][j].MinLat = math.MaxFloat64
			clusters[i][j].MinLng = math.MaxFloat64
		}
	}

	biggestShopsCount := 0
	for _, shop := range shops {
		i := int(math.Abs((shop.Lat - box.latMin) / clusterWidth))
		j := int(math.Abs((shop.Lng - box.lngMin) / clusterHeight))

		clusters[i][j].MinLat = math.Min(clusters[i][j].MinLat, shop.Lat)
		clusters[i][j].MaxLat = math.Max(clusters[i][j].MaxLat, shop.Lat)
		clusters[i][j].MinLng = math.Min(clusters[i][j].MinLng, shop.Lng)
		clusters[i][j].MaxLng = math.Max(clusters[i][j].MaxLng, shop.Lng)

		clusters[i][j].Lat += shop.Lat
		clusters[i][j].Lng += shop.Lng
		clusters[i][j].ShopsCount++

		if clusters[i][j].ShopsCount > biggestShopsCount {
			biggestShopsCount = clusters[i][j].ShopsCount
		}
	}

	var clustersJSON []gin.H
	for i := 0; i < int(clusterCount); i++ {
		for j := 0; j < int(clusterCount); j++ {
			if clusters[i][j].ShopsCount == 0 {
				continue
			}

			clusters[i][j].Lat /= float64(clusters[i][j].ShopsCount)
			clusters[i][j].Lng /= float64(clusters[i][j].ShopsCount)
			clusters[i][j].norm = clusterNormScaling(float64(clusters[i][j].ShopsCount) / float64(biggestShopsCount))

			clustersJSON = append(clustersJSON, gin.H{
				"bounds": gin.H{
					"min_lat": clusters[i][j].MinLat,
					"max_lat": clusters[i][j].MaxLat,
					"min_lng": clusters[i][j].MinLng,
					"max_lng": clusters[i][j].MaxLng,
				},
				"shops": clusters[i][j].ShopsCount,
				"norm":  clusters[i][j].norm,
				"lat":   clusters[i][j].Lat,
				"lng":   clusters[i][j].Lng,
			})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"clusters": clustersJSON,
	})
}

func getClustersHandler(service *services.KebabShopService) func(c *gin.Context) {
	return func(c *gin.Context) {
		shops, box, abort := getShopsInBox(c, service)
		if abort {
			return
		}

		createClusteredResponse(c, shops, box)
	}
}

func getAutoHandler(service *services.KebabShopService) func(c *gin.Context) {
	return func(c *gin.Context) {
		clusterThreshold, err := strconv.ParseInt(c.Query("ct"), 10, 32)
		if err != nil ||
			clusterThreshold < 1 ||
			clusterThreshold > maxClusterThreshold {
			clusterThreshold = defaultClusterThreshold
		}

		shops, box, abort := getShopsInBox(c, service)
		if abort {
			return
		}

		if len(shops) > int(clusterThreshold) {
			createClusteredResponse(c, shops, box)
		} else {
			createShopResponse(c, shops)
		}
	}
}

func getShopByID(service *services.KebabShopService) func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("shop_id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		shop, exists, err := service.KebabShop(int(id))
		if err != nil {
			log.Panic(err)
		}

		if !exists {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"shop": gin.H{
				"id":   strconv.Itoa(shop.ID),
				"name": shop.Name,
				"lat":  shop.Lat,
				"lng":  shop.Lng,
			},
		})
	}
}
