package users

import (
	"net/http"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/users"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request users.RegisterUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err := h.UserSvc.SignUp(ctx, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"error":   false,
		"message": "User created",
	})
}
