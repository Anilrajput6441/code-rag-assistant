const BASE_URL = "http://localhost:9000";

export async function ingestRepo(repoUrl: string) {
  const res = await fetch(`${BASE_URL}/ingest`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ repo_url: repoUrl }),
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Ingest failed");
  }

  return res.json();
}

export async function askQuestion(question: string, topK = 5) {
  const res = await fetch(`${BASE_URL}/query`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ question, top_k: topK }),
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Query failed");
  }

  return res.json();
}
