package main

import (
	"rag-go/internal/config"
	"rag-go/internal/ingest"
	"rag-go/internal/query"
	"time"

	"rag-go/internal/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()
	auth.InitFirebase()

	// ✅ CORS CONFIG
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/ingest", auth.RequireAuth(), ingest.IngestHandler(cfg))
	r.POST("/query", auth.RequireAuth(), query.QueryHandler())

	r.Run(":" + cfg.Port)
}
