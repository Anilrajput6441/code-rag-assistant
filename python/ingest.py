import os
import json

OUTPUT_FILE = "raw_chunks.jsonl"

def chunk_text(text, chunk_size=500):
    chunks = []
    for i in range(0, len(text), chunk_size):
        chunk = text[i:i + chunk_size]
        chunks.append(chunk)
    return chunks

def read_file(path):
    try:
        with open(path, "r", encoding="utf-8", errors="ignore") as f:
            return f.read()
    except Exception as e:
        return ""

def index_repo(repo_path):
    repo_name = os.path.basename(os.path.abspath(repo_path))
    print("Indexing repo:", repo_name)

    with open(OUTPUT_FILE, "w") as out:
        for root, _, files in os.walk(repo_path):
            for file in files:
                if not file.endswith((".go", ".js", ".ts")):
                    continue

                full_path = os.path.join(root, file)
                rel_path = os.path.relpath(full_path, repo_path)
                content = read_file(full_path)

                if not content.strip():
                    continue

                chunks = chunk_text(content)

                for idx, chunk in enumerate(chunks):
                    record = {
                        "repo": repo_name,
                        "file": rel_path,
                        "chunk_index": idx,
                        "content": chunk
                    }
                    out.write(json.dumps(record) + "\n")

    print("Indexing complete → raw_chunks.jsonl created")

if __name__ == "__main__":
    repo_path = input("Enter repo path: ")
    index_repo(repo_path)
