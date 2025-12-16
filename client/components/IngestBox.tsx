"use client";

import { useState } from "react";
import { ingestRepo } from "@/lib/api";

export default function IngestBox() {
  const [repo, setRepo] = useState("");
  const [status, setStatus] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  async function handleIngest() {
    if (!repo) return;

    try {
      setLoading(true);
      setStatus("Ingesting repository...");
      await ingestRepo(repo);
      setStatus("✅ Repository ingested successfully");
    } catch (err: any) {
      setStatus(`❌ ${err.message}`);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="border rounded-lg p-4 mb-6">
      <h2 className="font-semibold mb-2">Ingest Repository</h2>

      <input
        className="w-full border rounded p-2 mb-2"
        placeholder="https://github.com/user/repo"
        value={repo}
        onChange={(e) => setRepo(e.target.value)}
      />

      <button
        onClick={handleIngest}
        disabled={loading}
        className="px-4 py-2 bg-black text-white rounded"
      >
        {loading ? "Ingesting..." : "Ingest"}
      </button>

      {status && <p className="mt-2 text-sm">{status}</p>}
    </div>
  );
}
