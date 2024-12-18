package http

import (
	"api_gateway/internal/models"
	"api_gateway/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	postService services.PostService
}

func InitPostHandler(routes *gin.RouterGroup, postService services.PostService) {
	handler := &postHandler{}
	handler.postService = postService
	postRoutes := routes.Group("/posts")
	{
		// POST
		postRoutes.POST("/create", handler.Create)

		//GET
		postRoutes.GET("/", handler.ListPosts)
		postRoutes.GET("/:postId", handler.PostById)

	}

}

func (h *postHandler) Create(c *gin.Context) {
	model := &models.Post{}
	err := c.BindJSON(&model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
		return
	}

	id, err := h.postService.Create(model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *postHandler) PostById(c *gin.Context) {
	postId := c.Param("postId")

	post, err := h.postService.PostById(postId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandler) ListPosts(c *gin.Context) {
	posts, err := h.postService.ListPosts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "error")
		return
	}

	c.JSON(http.StatusOK, posts)
}
