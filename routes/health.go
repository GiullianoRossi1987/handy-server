package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.String(http.StatusOK, "server is working owo")
}

func SetRouter(r gin.IRouter) {
	r.GET("/health", GetHealth)
}
