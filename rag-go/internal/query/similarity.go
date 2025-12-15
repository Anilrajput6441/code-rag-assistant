package query

import "math"

// cosine similarity between two vectors
func CosineSimilarity(a, b []float32) float64 {
	var dot, normA, normB float64

	for i := range a {
		dot += float64(a[i] * b[i])
		normA += float64(a[i] * a[i])
		normB += float64(b[i] * b[i])
	}

	if normA == 0 || normB == 0 {
		return 0
	}

	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}
