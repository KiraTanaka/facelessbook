package handlers

import (
	"net/http"
	"post_service/internal/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type postHandler struct {
	PostService services.PostService
}

func InitPostRoutes(routes *gin.RouterGroup, postService services.PostService) {
	postHandler := &postHandler{PostService: postService}
	postRoutes := routes.Group("/posts")
	//GET
	postRoutes.GET("/", postHandler.GetListPosts)
	postRoutes.GET("/:id", postHandler.GetPost)

}

func (h *postHandler) GetListPosts(c *gin.Context) {
	posts, err := h.PostService.GetListPosts()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, posts)
}

func (h *postHandler) GetPost(c *gin.Context) {
	postId := c.Param("id")
	post, err := h.PostService.GetPost(postId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}
	log.Info(post)
	c.JSON(http.StatusOK, post)
}
