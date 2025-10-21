package comments

import (
	"net/http"
	"strconv"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/comments"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt64("userID")
	postIDParam := c.Param("postID")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid postid param",
		})

		return
	}

	var req comments.CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err = h.CommentSvc.CreateComment(ctx, userID, postID, req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "comment created",
	})
}

func (h *Handler) GetAllCommentByPostID(c *gin.Context) {
	ctx := c.Request.Context()

	postIDParam := c.Param("postID")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid postid param",
		})

		return
	}

	posts, err := h.CommentSvc.GetAllCommentsByPostID(ctx, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "comments retrieved",
		"data":    posts,
	})
}

func (h *Handler) UpdateComment(c *gin.Context) {
	ctx := c.Request.Context()
	postIDParam := c.Param("postID")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid userID param",
		})

		return
	}

	var req comments.CommentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err = h.CommentSvc.UpdateComment(ctx, postID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "comment updated",
	})
}

func (h *Handler) DeleteComment(c *gin.Context) {
	ctx := c.Request.Context()
	postIDParam := c.Param("postID")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "invalid postid param",
		})
		return
	}

	err = h.CommentSvc.DeleteComment(ctx, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"error":   false,
		"message": "comment deleted",
	})
}
