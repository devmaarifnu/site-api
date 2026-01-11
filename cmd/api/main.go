package main

import (
	"fmt"
	"log"

	"lpmaarifnu-site-api/internal/config"
	"lpmaarifnu-site-api/internal/database"
	"lpmaarifnu-site-api/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	if err := database.Initialize(&cfg.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Set Gin mode based on environment
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.New()

	// Setup routes
	routes.SetupRoutes(router, cfg)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	log.Printf("Starting %s v%s on %s (env: %s)",
		cfg.App.Name,
		cfg.App.Version,
		addr,
		cfg.App.Env,
	)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
