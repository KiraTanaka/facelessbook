package handlers

import (
	"net/http"
	"post_service/internal/db"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type postHandler struct {
	repository *db.Repository
}

func InitPostRoutes(routes *gin.RouterGroup, repository *db.Repository) {
	postHandler := &postHandler{}
	postHandler.repository = repository
	postRoutes := routes.Group("/posts")
	//GET
	postRoutes.GET("/", postHandler.GetListPosts)
	postRoutes.GET("/:id", postHandler.GetPost)

}

func (h *postHandler) GetListPosts(c *gin.Context) {
	posts, err := h.repository.GetListPosts()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
	}
	c.JSON(http.StatusOK, posts)
}

func (h *postHandler) GetPost(c *gin.Context) {
	postId := c.Param("id")
	post, err := h.repository.GetPost(postId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
	}
	log.Info(post)
	c.JSON(http.StatusOK, post)
}
