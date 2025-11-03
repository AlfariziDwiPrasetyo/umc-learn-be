package middleware

import (
	"time"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(cfg *configs.Config) gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins: cfg.Cors.AllowOrigins,
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"PATCH",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"Accept",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
		},
		AllowCredentials: cfg.Cors.AllowCredentials,
		MaxAge:           12 * time.Hour,
	}

	return cors.New(config)
}
