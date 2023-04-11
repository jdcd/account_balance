package server

import (
	"github.com/gin-gonic/gin"
)

// RouterDependencies keeps dependencies to allow dependencies injection
type RouterDependencies struct {
	CheckController *PingController
}

// SetupRouter returns a configured router Engine
func SetupRouter(d *RouterDependencies) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1Check := v1.Group("/ping")
	v1Check.GET("", d.CheckController.GetPing)

	return router
}
