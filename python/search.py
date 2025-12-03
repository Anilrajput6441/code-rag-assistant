import json
import os
import numpy as np
import google.generativeai as genai


EMBED_MODEL = "models/gemini-embedding-001"
EMBEDDINGS_FILE = "embeddings.jsonl"


def load_embeddings(path):
    records = []
    with open(path, "r", encoding="utf-8") as f:
        for line in f:
            records.append(json.loads(line))
    return records


def cosine_similarity(a, b):
    a = np.array(a)
    b = np.array(b)
    return np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b))


def embed_query(text):
    result = genai.embed_content(
        model=EMBED_MODEL,
        content=text,
        task_type="retrieval_query",
    )
    return result["embedding"]


def search(query, top_k=3):
    query_embedding = embed_query(query)

    records = load_embeddings(EMBEDDINGS_FILE)
    scored = []

    for record in records:
        score = cosine_similarity(query_embedding, record["embedding"])
        scored.append((score, record))

    scored.sort(key=lambda x: x[0], reverse=True)
    return scored[:top_k]


def main():
    api_key = os.getenv("GEMINI_API_KEY")
    if not api_key:
        raise RuntimeError("GEMINI_API_KEY not set")

    genai.configure(api_key=api_key)

    query = input("Ask a question about the code: ")
    results = search(query)

    print("\nTop matching code chunks:\n")

    for score, r in results:
        print(f"[Score: {score:.4f}]")
        print(f"Repo: {r['repo']}")
        print(f"File: {r['file']}")
        print(f"Chunk: {r['chunk_index']}")
        print("-" * 60)


if __name__ == "__main__":
    main()
