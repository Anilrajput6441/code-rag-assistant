package main

import (
	"rag-go/internal/config"
	"rag-go/internal/ingest"
	"rag-go/internal/query"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	r := gin.Default()

	r.POST("/ingest", ingest.IngestHandler(cfg))
	r.POST("/query", query.QueryHandler())

	r.Run(":" + cfg.Port)
}
