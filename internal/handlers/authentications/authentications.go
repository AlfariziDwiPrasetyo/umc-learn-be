package authentications

import (
	"net/http"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/model/authentications"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	ctx := c.Request.Context()

	var request authentications.RegisterUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	err := h.AuthSvc.SignUp(ctx, request)

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

func (h *Handler) SignIn(c *gin.Context) {
	ctx := c.Request.Context()

	var request authentications.LoginUser

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	token, err := h.AuthSvc.SignIn(ctx, request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Login Successfully",
		"data":    token,
	})

}

func (h *Handler) Refresh(c *gin.Context) {
	ctx := c.Request.Context()

	var request authentications.RefreshTokenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	token, err := h.AuthSvc.Refresh(ctx, request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Token refreshed successfully",
		"data":    token,
	})

}

func (h *Handler) LogOut(c *gin.Context) {
	ctx := c.Request.Context()

	token := c.GetHeader("Authorization")
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	} else {
		token = ""
	}

	_ = h.AuthSvc.LogOut(ctx, token)

	c.JSON(200, gin.H{
		"error":   false,
		"message": "Logout successful",
	})
}
