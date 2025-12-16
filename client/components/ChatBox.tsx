"use client";

import { useState } from "react";
import { askQuestion } from "@/lib/api";

export default function ChatBox() {
  const [question, setQuestion] = useState("");
  const [answer, setAnswer] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  async function handleAsk() {
    if (!question) return;

    try {
      setLoading(true);
      const res = await askQuestion(question);
      setAnswer(res.answer);
    } catch (err: any) {
      setAnswer(`❌ ${err.message}`);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="border rounded-lg p-4">
      <h2 className="font-semibold mb-2">Ask a Question</h2>

      <textarea
        className="w-full border rounded p-2 mb-2"
        rows={4}
        placeholder="Ask something about the repo..."
        value={question}
        onChange={(e) => setQuestion(e.target.value)}
      />

      <button
        onClick={handleAsk}
        disabled={loading}
        className="px-4 py-2 bg-black text-white rounded"
      >
        {loading ? "Thinking..." : "Ask"}
      </button>

      {answer && (
        <pre className="mt-4 whitespace-pre-wrap text-sm bg-gray-50 p-3 rounded">
          {answer}
        </pre>
      )}
    </div>
  );
}
