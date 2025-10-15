package middleware

import (
	"net/http"
	"strings"

	"github.com/alfarizidwiprasetyo/be-umc-learn/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid authorization header format",
			})

			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := jwt.ValidateToken(tokenString, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   true,
				"message": "Invalid or expired token",
			})

			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
