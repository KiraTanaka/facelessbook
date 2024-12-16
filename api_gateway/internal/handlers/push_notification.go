package handlers

import (
	"notification_service/internal/db"
	"notification_service/internal/services"

	"github.com/gin-gonic/gin"
)

type notificationHandler struct {
	notificationService services.NewNotificationService
}

func InitPostLikesRoutes(routes *gin.RouterGroup, repository *db.Repository) {
	postLikesHandler := &notificationHandler{}
	postLikesHandler.repository = repository
	//postRoutes := routes.Group("/posts")
	//GET
	//postRoutes.GET("/", postLikesHandler.GetListPosts)
	//postRoutes.GET("/:id", postLikesHandler.GetPost)

}

/*
func (h *postLikesHandler) GetListPosts(c *gin.Context) {
	posts, err := h.repository.GetListPosts()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
	}
	c.JSON(http.StatusOK, posts)
}*/
/*
func (h *postHandler) GetPost(c *gin.Context) {
	postId := c.Param("id")
	post, err := h.repository.GetPost(postId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
	}
	log.Info(post)
	c.JSON(http.StatusOK, post)
}
*/
