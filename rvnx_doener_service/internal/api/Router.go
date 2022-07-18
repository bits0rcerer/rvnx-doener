package api

import (
	"github.com/gin-gonic/gin"
	v1 "rvnx_doener_service/internal/api/v1"
	"rvnx_doener_service/internal/services"
)

func BuildRouter(env *services.ServiceEnvironment) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	v1.RouteV1(engine.Group("/v1"), env)

	return engine
}
