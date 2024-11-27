package handlers

import (
	"facelessbook/post_service/app/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(repository *db.Repository) *gin.Engine {
	routes := gin.Default()
	routes.GET("/", hello)
	return routes
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}
