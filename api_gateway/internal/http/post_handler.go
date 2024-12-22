package http

import (
	"api_gateway/internal/models"
	"api_gateway/internal/services"
	"encoding/json"
	"io"
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

	bodyAsByteArray, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "недопустимое тело запроса для редактирования поста")
		return
	}

	jsonBody := make(map[string]string)
	if err = json.Unmarshal(bodyAsByteArray, &jsonBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "неудалось обработать входные данные для редактирования поста")
		return
	}

	if err = h.postService.Update(postId, jsonBody["new_text"]); err != nil {
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
