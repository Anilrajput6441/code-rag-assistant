package user

import (
	"context"
	"fmt"
	"net/http"

	"rag-go/internal/auth"
	"rag-go/internal/security"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

type APIKeyReq struct {
	APIKey string `json:"api_key"`
}

func SaveAPIKey(secret string) gin.HandlerFunc {
	fmt.Print("save api hit")
	return func(c *gin.Context) {
		uid := c.GetString("userId")

		var req APIKeyReq
		if err := c.ShouldBindJSON(&req); err != nil || req.APIKey == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "api_key required"})
			return
		}

		enc, err := security.Encrypt(req.APIKey, secret)
		if err != nil {
			c.JSON(500, gin.H{"error": "encryption failed"})
			return
		}
		fmt.Println("encrypted key:", enc)

		// Use proper context for Firestore
		ctx := context.Background()
		_, err = auth.Firestore.Collection("users").
			Doc(uid).
			Set(ctx, map[string]interface{}{
				"gemini_api_key": enc,
			}, firestore.MergeAll)

		if err != nil {
			fmt.Printf("Firestore save error: %v\n", err)
			c.JSON(500, gin.H{"error": "save failed: " + err.Error()})
			return
		}

		c.JSON(200, gin.H{"status": "api key saved"})
	}
}
