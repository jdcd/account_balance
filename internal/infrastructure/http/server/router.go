package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jdcd/account_balance/internal/infrastructure/http/server/controller"
)

// RouterDependencies keeps dependencies to allow dependencies injection
type RouterDependencies struct {
	CheckController  *controller.PingController
	ReportController *controller.ReportController
}

// SetupRouter returns a configured router Engine
func SetupRouter(d *RouterDependencies) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1Check := v1.Group("/ping")
	v1Check.GET("", d.CheckController.GetPing)

	v1Report := v1.Group("/report")
	v1Report.POST("", d.ReportController.PostReport)

	return router
}
