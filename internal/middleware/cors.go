package middleware

import (
	"lpmaarifnu-site-api/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS(cfg *config.CORSConfig) gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins:     cfg.AllowedOrigins,
		AllowMethods:     cfg.AllowedMethods,
		AllowHeaders:     cfg.AllowedHeaders,
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	return cors.New(corsConfig)
}
