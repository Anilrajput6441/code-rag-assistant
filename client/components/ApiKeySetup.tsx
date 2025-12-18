"use client";

import { useState } from "react";
import { saveApiKey } from "@/lib/api";

export default function ApiKeySetup() {
  const [apiKey, setApiKey] = useState("");
  const [loading, setLoading] = useState(false);
  const [saved, setSaved] = useState(false);

  const handleSave = async () => {
    if (!apiKey.trim()) return;
    
    setLoading(true);
    try {
      await saveApiKey(apiKey);
      setSaved(true);
      setApiKey("");
    } catch (error) {
      alert("Failed to save API key: " + error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-md mb-6">
      <h3 className="text-lg font-semibold mb-4">🔑 Gemini API Key Setup</h3>
      
      {saved && (
        <div className="bg-green-100 text-green-800 p-3 rounded mb-4">
          ✅ API Key saved successfully!
        </div>
      )}
      
      <div className="flex gap-3">
        <input
          type="password"
          placeholder="Enter your Gemini API key"
          value={apiKey}
          onChange={(e) => setApiKey(e.target.value)}
          className="flex-1 px-3 py-2 border rounded-md"
        />
        <button
          onClick={handleSave}
          disabled={loading || !apiKey.trim()}
          className="px-4 py-2 bg-blue-500 text-white rounded-md disabled:opacity-50"
        >
          {loading ? "Saving..." : "Save"}
        </button>
      </div>
      
      <p className="text-sm text-gray-600 mt-2">
        Get your API key from{" "}
        <a 
          href="https://makersuite.google.com/app/apikey" 
          target="_blank" 
          className="text-blue-500 underline"
        >
          Google AI Studio
        </a>
      </p>
    </div>
  );
}