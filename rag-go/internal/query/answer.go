package query

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"sort"
	"strings"

	"rag-go/internal/storage"

	"google.golang.org/genai"
)

type scoredChunk struct {
	Text  string
	Score float64
}

func AnswerQuestion(question string, topK int) (string, error) {
	// 1️⃣ Load vectors
	data, err := os.ReadFile("data/vectors.json")
	if err != nil {
		return "", err
	}

	var records []storage.VectorRecord
	if err := json.Unmarshal(data, &records); err != nil {
		return "", err
	}

	// 2️⃣ Create Gemini client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return "", err
	}

	// 3️⃣ Embed the question
	qEmbedRes, err := client.Models.EmbedContent(
		ctx,
		"gemini-embedding-001",
		[]*genai.Content{
			genai.NewContentFromText(question, genai.RoleUser),
		},
		nil,
	)
	if err != nil {
		return "", err
	}

	if len(qEmbedRes.Embeddings) == 0 {
		return "", errors.New("question embedding failed")
	}

	queryVec := qEmbedRes.Embeddings[0].Values

	// 4️⃣ Compute similarity
	var scored []scoredChunk
	for _, r := range records {
		score := CosineSimilarity(queryVec, r.Embedding)
		scored = append(scored, scoredChunk{
			Text:  r.Content,
			Score: score,
		})
	}

	// 5️⃣ Sort by similarity
	sort.Slice(scored, func(i, j int) bool {
		return scored[i].Score > scored[j].Score
	})

	if topK > len(scored) {
		topK = len(scored)
	}

	// 6️⃣ Build context
	var contextText strings.Builder
	for i := 0; i < topK; i++ {
		contextText.WriteString(scored[i].Text)
		contextText.WriteString("\n---\n")
	}

	// 7️⃣ Ask Gemini with context
	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		[]*genai.Content{
			genai.NewContentFromText(
				"Answer the question using ONLY the context below:\n\n"+
					contextText.String()+
					"\nQuestion: "+question,
				genai.RoleUser,
			),
		},
		nil,
	)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
