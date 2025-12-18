package ingest

import (
	"net/http"

	"rag-go/internal/config"
	"rag-go/internal/user"

	"github.com/gin-gonic/gin"
)

// ================== REQUEST MODELS ==================
type IngestRequest struct {
	RepoURL string `json:"repo_url"`
}

// ================== HTTP HANDLERS ==================
func IngestHandler(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req IngestRequest

		// ========== INPUT VALIDATION ==========
		if err := c.ShouldBindJSON(&req); err != nil || req.RepoURL == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "repo_url required"})
			return
		}

		// ========== GET USER API KEY ==========
		userID := c.GetString("userId")
		apiKey, err := user.GetUserAPIKey(userID, cfg.EncryptionSecret)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "API key not found. Please save your API key first."})
			return
		}

		// ========== PROCESS INGESTION ==========
		if err := IngestRepo(apiKey, req.RepoURL); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// ========== SUCCESS RESPONSE ==========
		c.JSON(http.StatusOK, gin.H{"status": "repo ingested successfully"})
	}
}
