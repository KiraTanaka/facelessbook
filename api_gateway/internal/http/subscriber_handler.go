package http

import (
	"api_gateway/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type subscriberHandler struct {
	subscriberService services.SubscriberService
}

func NewSubscriberHandler(routes *gin.RouterGroup, subscriberService services.SubscriberService) {
	handler := &subscriberHandler{}
	handler.subscriberService = subscriberService
	subscriberRoutes := routes.Group("/users")
	{
		subscriberRoutes.POST("/:userId/subscribe", handler.Subscribe)

		subscriberRoutes.GET("/:userId/subscribers", handler.ListSubscribers)

		subscriberRoutes.DELETE("/:userId/unsubscribe", handler.Unsubscribe)

	}

}

func (h *subscriberHandler) Subscribe(c *gin.Context) {
	publisherId := c.Param("userId")
	if publisherId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "указание ид пользователя, на которого хотите подписаться, обязательно")
	}

	jsonBody, err := ReadBody(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err = h.subscriberService.Subscribe(publisherId, jsonBody["subscriber_id"]); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось подписаться на пользователя")
		return

	}

	c.JSON(http.StatusOK, "ok")
}

func (h *subscriberHandler) Unsubscribe(c *gin.Context) {
	publisherId := c.Param("userId")
	if publisherId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "указание ид пользователя, от которого хотите отписаться, обязательно")
	}

	jsonBody, err := ReadBody(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err = h.subscriberService.Unsubscribe(publisherId, jsonBody["subscriber_id"]); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось отписаться от пользователя")
		return

	}

	c.JSON(http.StatusOK, "ok")
}

func (h *subscriberHandler) ListSubscribers(c *gin.Context) {
	publisherId := c.Param("userId")
	if publisherId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "указание ид пользователя, для которого хотите найти подписчиков, обязательно")
	}

	subscribers, err := h.subscriberService.ListSubscribers(publisherId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось найти подписчиков")
		return
	}

	c.JSON(http.StatusOK, subscribers)
}
