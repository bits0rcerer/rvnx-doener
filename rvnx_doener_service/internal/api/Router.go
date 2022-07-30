package api

import (
	"github.com/gin-gonic/gin"
	v1 "rvnx_doener_service/internal/api/v1"
	"rvnx_doener_service/internal/services"
)

func BuildEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	return engine
}

func RouteAPI(router *gin.RouterGroup, env *services.ServiceEnvironment) {
	v1.RouteV1(router.Group("/v1"), env)
}
