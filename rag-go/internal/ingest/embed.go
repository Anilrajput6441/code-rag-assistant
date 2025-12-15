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

func NewEmbedder() (*Embedder, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &Embedder{
		client: client,
		ctx:    ctx,
	}, nil
}

func (e *Embedder) EmbedText(text string) ([]float32, error) {
	contents := []*genai.Content{
		genai.NewContentFromText(text, genai.RoleUser),
	}

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

	return result.Embeddings[0].Values, nil
}
