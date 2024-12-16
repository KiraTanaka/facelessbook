package handlers

import (
	"assessment_service/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(repository *db.Repository) *gin.Engine {
	routes := gin.Default()
	routes.GET("/", hello)
	routeGroup := routes.Group("/api")

	InitPostLikesRoutes(routeGroup, repository)
	return routes
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
