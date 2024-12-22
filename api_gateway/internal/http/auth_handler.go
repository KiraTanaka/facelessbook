package http

import (
	"api_gateway/internal/services"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(routes *gin.RouterGroup, authService services.AuthService) {
	authHandler := &authHandler{}
	authHandler.authService = authService
	authRoutes := routes.Group("/users")
	{
		// POST
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/login", authHandler.Login)

	}

}

func (h *authHandler) Register(c *gin.Context) {
	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "недопустимое тело запроса для регистрации")
	}

	jsonBody := make(map[string]string)
	if err = json.Unmarshal(bodyAsByteArray, &jsonBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "неудалось обработать входные данные для регистрации")
	}

	userId, err := h.authService.Register(jsonBody["phone"], jsonBody["password"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось зарегистрировать пользователя")
	}

	c.JSON(http.StatusOK, userId)
}

func (h *authHandler) Login(c *gin.Context) {
	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "недопустимое тело запроса для входа")
	}

	jsonBody := make(map[string]string)
	if err = json.Unmarshal(bodyAsByteArray, &jsonBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "неудалось обработать входные данные для регистрации")
	}

	token, err := h.authService.Login(jsonBody["phone"], jsonBody["password"])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось авторизовать пользователя")
	}
	c.JSON(http.StatusOK, token)
}
