package main

import (
	"rag-go/internal/config"
	"rag-go/internal/ingest"
	"rag-go/internal/query"
	"rag-go/internal/user"
	"time"

	"rag-go/internal/auth"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()
	
	if err := auth.InitFirebase(); err != nil {
		panic("Failed to initialize Firebase: " + err.Error())
	}
	
	if err := auth.InitFirestore(cfg.FirebaseProjectID); err != nil {
		panic("Failed to initialize Firestore: " + err.Error())
	}
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
	r.POST("/user/api-key", auth.RequireAuth(), user.SaveAPIKey(cfg.EncryptionSecret))

	r.Run(":" + cfg.Port)
}
