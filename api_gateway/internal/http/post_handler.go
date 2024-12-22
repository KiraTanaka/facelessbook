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

func NewPostHandler(routes *gin.RouterGroup, postService services.PostService) {
	handler := &postHandler{}
	handler.postService = postService
	postRoutes := routes.Group("/posts")
	{
		postRoutes.POST("/create", handler.Create)

		postRoutes.GET("/", handler.ListPosts)
		postRoutes.GET("/:postId", handler.PostById)

		postRoutes.PUT("/:postId", handler.Update)

		postRoutes.DELETE("/:postId", handler.Delete)

	}

}

func (h *postHandler) Create(c *gin.Context) {
	model := &models.Post{}
	err := c.BindJSON(&model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "неудалось обработать входные данные для создания поста")
		return
	}

	id, err := h.postService.Create(model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось создать пост")
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *postHandler) PostById(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "указание ид поста обязательно")
	}

	post, err := h.postService.PostById(postId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось найти пост")
	}

	c.JSON(http.StatusOK, post)
}

func (h *postHandler) ListPosts(c *gin.Context) {
	posts, err := h.postService.ListPosts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось найти посты")
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *postHandler) Update(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "указание ид поста обязательно")
	}

	jsonBody, err := ReadBody(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	if err = h.postService.Update(postId, jsonBody["text"]); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось отредактировать пост")
		return

	}

	c.JSON(http.StatusOK, "ok")
}

func (h *postHandler) Delete(c *gin.Context) {
	postId := c.Param("postId")
	if postId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "указание ид поста обязательно")
	}

	if err := h.postService.Delete(postId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "неудалось удалить пост")
		return
	}

	c.JSON(http.StatusOK, "ok")
}
