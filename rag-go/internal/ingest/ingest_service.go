package ingest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rag-go/internal/config"
	"rag-go/internal/storage"
)

func IngestRepo(cfg config.Config, repoURL string) error {
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
	embedder, err := NewEmbedder() // 🔄 CHANGED
	if err != nil {
		return err
	}

	// ================== STEP 3: GENERATE EMBEDDINGS ==================
	fmt.Println("Generating embeddings...")

	for i := range records {
		//  USE SDK EMBEDDER (NO API KEY PASSED)
		vec, err := embedder.EmbedText(records[i].Content)
		if err != nil {
			return err
		}
		records[i].Embedding = vec

		if i%20 == 0 {
			fmt.Printf("Progress: %d/%d chunks embedded\n", i, len(records))
		}
	}

	// ================== STEP 4: SAVE TO DISK ==================
	os.MkdirAll("data", 0755)

	data, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile("data/vectors.json", data, 0644); err != nil {
		return err
	}

	fmt.Println("✅ Successfully saved", len(records), "vectors to data/vectors.json")
	return nil
}
