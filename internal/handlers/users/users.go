package users

import (
	"net/http"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.GetInt64("userID")

	var request users.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err := h.UserSvc.UpdateUser(ctx, userID, request)

	if err != nil {
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
