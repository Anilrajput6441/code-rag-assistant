package storage

type VectorRecord struct {
	Repo      string    `json:"repo"`
	File      string    `json:"file"`
	Content   string    `json:"content"`
	Embedding []float32 `json:"embedding"`
}
