package likes

import (
	"net/http"
	"strconv"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateLike(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.GetInt64("userID")
	postIDParam := c.Param("id")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "param not valid",
		})
		return
	}

	err = h.LikesSvc.CreateLike(ctx, postID, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "like created",
	})
}

func (h *Handler) DeleteLike(c *gin.Context) {
	ctx := c.Request.Context()

	userID := c.GetInt64("userID")
	postIDParam := c.Param("id")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "param not valid",
		})
		return
	}

	err = h.LikesSvc.DeleteLike(ctx, postID, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error":   false,
		"message": "like deleted",
	})
}

func (h *Handler) GetLikesByPostID(c *gin.Context) {
	ctx := c.Request.Context()

	postIDParam := c.Param("id")

	postID, err := strconv.ParseInt(postIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "param not valid",
		})
		return
	}

	likes, err := h.LikesSvc.GetLikesByPostID(ctx, postID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	data := dto.ToLikeResponses(likes)

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "likes retrieved",
		"data":    data,
	})
}
