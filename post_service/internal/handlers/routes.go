package handlers

import (
	"net/http"
	"post_service/internal/db"

	"github.com/gin-gonic/gin"
)

func InitRoutes(repository *db.Repository) *gin.Engine {
	routes := gin.Default()
	routes.GET("/", hello)
	routeGroup := routes.Group("/api")

	InitPostRoutes(routeGroup, repository)
	return routes
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
