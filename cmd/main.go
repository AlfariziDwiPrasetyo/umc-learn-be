package main

import (
	"fmt"
	"log"

	"github.com/alfarizidwiprasetyo/be-umc-learn/internal/configs"
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
	database.Connect(cfg)
	database.Migrate()

	// Run server
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	log.Printf("Server running on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
