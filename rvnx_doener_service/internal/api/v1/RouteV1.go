package v1

import (
	"github.com/gin-gonic/gin"
	"rvnx_doener_service/internal/services"
)

func RouteV1(r *gin.RouterGroup, env *services.ServiceEnvironment) {
	RouteKebabShops(r.Group("/kebabshops"), env)
}
