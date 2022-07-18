package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rvnx_doener_service/internal/services"
	"strconv"
)

func RouteKebabShops(r *gin.RouterGroup, env *services.ServiceEnvironment) {
	r.GET("/box", getBoundingBoxHandler(env.KebabShopService))
}

func getBoundingBoxHandler(service *services.KebabShopService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var latMin, latMax, lngMin, lngMax float64

		latMin, err := strconv.ParseFloat(c.Query("ltm"), 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		latMax, err = strconv.ParseFloat(c.Query("ltx"), 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		lngMin, err = strconv.ParseFloat(c.Query("lnm"), 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		lngMax, err = strconv.ParseFloat(c.Query("lnx"), 64)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		shops, err := service.Within(latMin, latMax, lngMin, lngMax, "id", "lat", "lng")
		if err != nil {
			log.Panic(err)
		}

		shopsJSON := make([]gin.H, len(shops))
		for i := 0; i < len(shopsJSON); i++ {
			shopsJSON[i] = gin.H{
				"id":  strconv.Itoa(shops[i].ID),
				"lat": strconv.FormatFloat(shops[i].Lat, 'g', -1, 64),
				"lng": strconv.FormatFloat(shops[i].Lng, 'g', -1, 64),
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"cords": shopsJSON,
		})
	}
}
