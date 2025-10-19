package main

import (
	"fmt"
	"log"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
	authHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/authentications"
	postHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/posts"
	userHandler "github.com/alfarizidwiprasetyo/be-umc-learn/internal/handlers/users"
	authRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/authentications"
	postRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/posts"
	userRepository "github.com/alfarizidwiprasetyo/be-umc-learn/internal/repository/users"
	authService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/authentications"
	postService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/posts"
	userService "github.com/alfarizidwiprasetyo/be-umc-learn/internal/service/users"
	"github.com/alfarizidwiprasetyo/be-umc-learn/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Config
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("failed to parse YAML config: %v", err)
	}

	// Database
	db := database.Connect(cfg)
	database.Migrate(db)

	// User
	userRepo := userRepository.NewRepository(db)
	userSvc := userService.NewService(cfg, userRepo)
	userHandler := userHandler.NewHandler(r, cfg, userSvc)

	userHandler.RegisterRoute()

	// Authentication
	authRepo := authRepository.NewRepository(db)
	authSvc := authService.NewService(cfg, userRepo, authRepo)
	authHandler := authHandler.NewHandler(r, authSvc)

	authHandler.RegisterRoute()

	// Post
	postRepo := postRepository.NewRepository(db)
	postSvc := postService.NewService(cfg, postRepo)
	postHandler := postHandler.NewHandler(r, cfg, postSvc)

	postHandler.RegisterRoute()

	// Run server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	log.Printf("Server running on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
