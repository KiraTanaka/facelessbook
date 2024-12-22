package http

import (
	"api_gateway/internal/services"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoutes(services *services.Services) *gin.Engine {
	routes := gin.Default()
	routes.GET("/", hello)
	routeGroup := routes.Group("/api")
	{
		NewAuthHandler(routeGroup, services.Auth)
		NewPostHandler(routeGroup, services.Post)
		NewSubscriberHandler(routeGroup, services.Subscriber)
	}

	return routes
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "hello")
}

func ReadBody(request *http.Request) (map[string]string, error) {
	bodyAsByteArray, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, fmt.Errorf("недопустимое тело запроса: %w", err)
	}

	jsonBody := map[string]string{}
	if err = json.Unmarshal(bodyAsByteArray, &jsonBody); err != nil {
		return nil, fmt.Errorf("неудалось обработать входные данные: %w", err)
	}
	return jsonBody, nil
}
