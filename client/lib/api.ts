import { auth } from "./firebase";

const BASE_URL = "http://localhost:9000";

export async function ingestRepo(repoUrl: string) {
  const authHeader = await getAuthHeader();

  const res = await fetch(`${BASE_URL}/ingest`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...authHeader,
    },
    body: JSON.stringify({ repo_url: repoUrl }),
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Ingest failed");
  }

  return res.json();
}

export async function askQuestion(question: string, topK = 5) {
  const authHeader = await getAuthHeader();

  const res = await fetch(`${BASE_URL}/query`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      ...authHeader,
    },
    body: JSON.stringify({ question, top_k: topK }),
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Query failed");
  }

  return res.json();
}

async function getAuthHeader() {
  const user = auth.currentUser;
  if (!user) {
    throw new Error("User not authenticated");
  }

  const token = await user.getIdToken();
  return {
    Authorization: `Bearer ${token}`,
  };
}