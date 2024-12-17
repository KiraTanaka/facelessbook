package handlers

import (
	"api_gateway/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(authService services.AuthService) *gin.Engine {
	routes := gin.Default()
	routes.GET("/", hello)
	routeGroup := routes.Group("/api")
	{
		InitAuthHandler(routeGroup, authService)
	}

	return routes
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
