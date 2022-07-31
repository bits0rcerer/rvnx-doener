package twitch

import (
	"github.com/gin-gonic/gin"
	"rvnx_doener_service/internal/services"
)

func RouteTwitch(r *gin.RouterGroup, env *services.ServiceEnvironment) {
	RouteTwitchAuth(r, env)
}
