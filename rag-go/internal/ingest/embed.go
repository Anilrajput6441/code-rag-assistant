package ingest

import (
	"context"
	"errors"

	"google.golang.org/genai"
)

type Embedder struct {
	client *genai.Client
	ctx    context.Context
}

func NewEmbedder(apiKey string) (*Embedder, error) {
	ctx := context.Background()

	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, err
	}

	return &Embedder{
		client: client,
		ctx:    ctx,
	}, nil
}

func (e *Embedder) EmbedText(text string) ([]float64, error) {
	contents := []*genai.Content{
		genai.NewContentFromText(text, genai.RoleUser)}

	result, err := e.client.Models.EmbedContent(
		e.ctx,
		"gemini-embedding-001",
		contents,
		nil,
	)
	if err != nil {
		return nil, err
	}

	if len(result.Embeddings) == 0 {
		return nil, errors.New("no embeddings returned")
	}

	values := result.Embeddings[0].Values
	float64Values := make([]float64, len(values))
	for i, v := range values {
		float64Values[i] = float64(v)
	}

	return float64Values, nil
}
