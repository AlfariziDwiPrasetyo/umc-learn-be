package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/posts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) CreatePost(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt64("userID")

	var req posts.PostRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err := h.PostSvc.CreatePost(ctx, userID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Post created",
	})
}

func (h *Handler) UpdatePost(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid id param",
		})

		return
	}

	var req posts.PostUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}

	err = h.PostSvc.UpdatePost(ctx, userID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Post updated",
	})
}

func (h *Handler) DeletePost(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")
	postId, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid id param",
		})

		return
	}

	err = h.PostSvc.DeletePost(ctx, postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Post deleted",
	})
}

func (h *Handler) GetAllPost(c *gin.Context) {
	ctx := c.Request.Context()
	limitQuery := c.DefaultQuery("limit", "15")
	limit, err := strconv.Atoi(limitQuery)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid limit value",
		})
		return
	}

	var posts []posts.Post

	posts, err = h.PostSvc.GetPosts(ctx, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "posts retrieved",
		"data":    posts,
	})
}

func (h *Handler) GetPostById(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")
	postParam, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid id value",
		})
		return
	}

	post, err := h.PostSvc.GetPostById(ctx, postParam)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   true,
				"message": "User not found",
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "posts retrieved",
		"data":    post,
	})
}
