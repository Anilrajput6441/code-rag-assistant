package main

import (
	"rag-go/internal/config"
	"rag-go/internal/ingest"
	"rag-go/internal/query"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	// ✅ CORS CONFIG
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/ingest", ingest.IngestHandler(cfg))
	r.POST("/query", query.QueryHandler())

	r.Run(":" + cfg.Port)
}
