package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingController basic controller to check the app connection
type PingController struct{}

func (r *PingController) GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
