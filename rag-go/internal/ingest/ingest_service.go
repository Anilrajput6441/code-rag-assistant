package ingest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rag-go/internal/storage"
)

func IngestRepo(apiKey string, repoURL string) error {
	fmt.Println("Ingesting:", repoURL)

	repoDir := filepath.Join(os.TempDir(), "rag_repo")

	// ================== SETUP & CLEANUP ==================
	_ = os.RemoveAll(repoDir)

	// ================== STEP 1: CLONE REPO ==================
	if err := CloneRepo(repoURL, repoDir); err != nil {
		return fmt.Errorf("clone failed: %w", err)
	}

	var records []storage.VectorRecord

	// ================== STEP 2: READ & CHUNK FILES ==================
	err := filepath.Walk(repoDir, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		// Skip unwanted directories
		if strings.Contains(path, "vendor") ||
			strings.Contains(path, "node_modules") ||
			strings.Contains(path, ".git") ||
			strings.Contains(path, "test") ||
			strings.Contains(path, "docs") {
			return nil
		}

		// Allow only core code files
		if !strings.HasSuffix(path, ".go") &&
			!strings.HasSuffix(path, ".ts") {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		chunks := ChunkText(string(data), 1200)

		for _, chunk := range chunks {
			records = append(records, storage.VectorRecord{
				Repo:    repoURL,
				File:    path,
				Content: chunk,
			})
		}

		return nil
	})

	if err != nil {
		return err
	}

	// ================== SAFETY LIMIT ==================
	MAX_CHUNKS := 600
	if len(records) > MAX_CHUNKS {
		records = records[:MAX_CHUNKS]
		fmt.Println("Applied safety limit - reduced to:", MAX_CHUNKS, "chunks")
	}

	fmt.Println("Total chunks to process:", len(records))

	// ==================  CREATE EMBEDDER ONCE ==================
	fmt.Println("Creating embedder...")
	embedder, err := NewEmbedder(apiKey)
	if err != nil {
		fmt.Printf("❌ Failed to create embedder: %v\n", err)
		return err
	}
	fmt.Println("✅ Embedder created successfully")

	// ================== STEP 3: GENERATE EMBEDDINGS ==================
	fmt.Println("Generating embeddings...")

	for i := range records {
		// Generate embedding
		vec, err := embedder.EmbedText(records[i].Content)
		if err != nil {
			return err
		}

		// Convert []float64 to []float32
		float32Vec := make([]float32, len(vec))
		for j, v := range vec {
			float32Vec[j] = float32(v)
		}
		records[i].Embedding = float32Vec

		if i%20 == 0 {
			fmt.Printf("Progress: %d/%d chunks embedded\n", i, len(records))
		}
	}

	// ================== STEP 4: SAVE TO DISK ==================
	os.MkdirAll("data", 0755)

	fmt.Println("Serializing", len(records), "records to JSON...")
	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		fmt.Printf("❌ Failed to marshal JSON: %v\n", err)
		return err
	}

	fmt.Printf("Writing %d bytes to vectors.json...\n", len(data))
	if err := os.WriteFile("data/vectors.json", data, 0644); err != nil {
		fmt.Printf("❌ Failed to write file: %v\n", err)
		return err
	}

	fmt.Println("✅ Successfully saved", len(records), "vectors to data/vectors.json")
	return nil
}
