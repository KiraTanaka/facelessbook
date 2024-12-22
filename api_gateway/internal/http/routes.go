package http

import (
	"api_gateway/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutes(services *services.Services) *gin.Engine {
	routes := gin.Default()
	routes.GET("/", hello)
	routeGroup := routes.Group("/api")
	{
		InitAuthHandler(routeGroup, services.Auth)
		InitPostHandler(routeGroup, services.Post)
	}

	return routes
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
