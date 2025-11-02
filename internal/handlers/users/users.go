package users

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/dto"
	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")

	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	var request users.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err = h.UserSvc.UpdateUser(ctx, userID, request)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   true,
				"message": "User not found",
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "User updated",
	})

}

func (h *Handler) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")

	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err = h.UserSvc.DeleteUser(ctx, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "User deleted",
	})

}

func (h *Handler) GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")

	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	user, err := h.UserSvc.GetUser(ctx, userID)

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

	data := dto.ToUserResponse(*user)

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "User retrieved",
		"data":    data,
	})

}
