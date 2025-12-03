import json
import os
from typing import List

import google.generativeai as genai


# -------- CONFIG --------
INPUT_FILE = "raw_chunks.jsonl"
OUTPUT_FILE = "embeddings.jsonl"
EMBED_MODEL = "models/gemini-embedding-001"
# ------------------------


def load_chunks(path: str):
    """Read raw chunk records from JSONL file."""
    with open(path, "r", encoding="utf-8") as f:
        for line in f:
            yield json.loads(line)


def write_embedding(record: dict, embedding: List[float], out):
    """Write embedded record to output file."""
    embedded_record = {
        "repo": record["repo"],
        "file": record["file"],
        "chunk_index": record["chunk_index"],
        "content": record["content"],
        "embedding": embedding,
    }
    out.write(json.dumps(embedded_record) + "\n")


def main():
    api_key = os.getenv("GEMINI_API_KEY")
    if not api_key:
        raise RuntimeError("GEMINI_API_KEY not set")

    genai.configure(api_key=api_key)

    model = genai.embed_content

    print("Loading chunks from:", INPUT_FILE)
    print("Writing embeddings to:", OUTPUT_FILE)

    with open(OUTPUT_FILE, "w", encoding="utf-8") as out:
        for idx, record in enumerate(load_chunks(INPUT_FILE)):
            text = record["content"]

            try:
                result = model(
                    model=EMBED_MODEL,
                    content=text,
                    task_type="retrieval_document",
                )

                embedding = result["embedding"]
                write_embedding(record, embedding, out)

                if idx % 10 == 0:
                    print(f"Embedded {idx} chunks")

            except Exception as e:
                print("Embedding failed:", e)

    print("Embedding completed ✅")


if __name__ == "__main__":
    main()
