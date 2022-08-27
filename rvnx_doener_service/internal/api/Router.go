package api

import (
	"rvnx_doener_service/internal/api/session"
	"rvnx_doener_service/internal/api/twitch"
	v1 "rvnx_doener_service/internal/api/v1"
	"rvnx_doener_service/internal/log"
	"rvnx_doener_service/internal/services"

	"github.com/gin-gonic/gin"
)

func BuildEngine() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.LoggerWithFormatter(log.CustomLogFormatter))
	return engine
}

func RouteAPI(router *gin.RouterGroup, env *services.ServiceEnvironment) {
	session.InitSessions(router)

	v1.RouteV1(router.Group("/v1"), env)
	twitch.RouteTwitch(router.Group("/twitch"), env)
}
