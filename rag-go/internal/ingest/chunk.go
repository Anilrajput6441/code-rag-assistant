package ingest

// ================== TEXT PROCESSING ==================
func ChunkText(text string, size int) []string {
	var chunks []string

	// ========== SPLIT TEXT INTO FIXED-SIZE CHUNKS ==========
	for i := 0; i < len(text); i += size {
		// Calculate chunk end position
		end := i + size
		if end > len(text) {
			end = len(text)
		}

		// Extract chunk and add to collection
		chunks = append(chunks, text[i:end])
	}

	return chunks
}
