package query

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type QueryRequest struct {
	Question string `json:"question"`
	TopK     int    `json:"top_k"`
}

func QueryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req QueryRequest
		if err := c.ShouldBindJSON(&req); err != nil || req.Question == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "question required"})
			return
		}

		if req.TopK == 0 {
			req.TopK = 5
		}

		answer, err := AnswerQuestion(req.Question, req.TopK)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"answer": answer})
	}
}
