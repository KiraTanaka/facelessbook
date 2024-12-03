package handlers

import (
	"net/http"
	"post_service/internal/services"

	"github.com/gin-gonic/gin"
)

func InitRoutes(postService services.PostService) *gin.Engine {
	routes := gin.Default()
	routes.GET("/", hello)
	routeGroup := routes.Group("/api")

	InitPostRoutes(routeGroup, postService)
	return routes
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
