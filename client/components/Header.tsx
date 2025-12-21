"use client";

import { auth } from "@/lib/firebase";
import { useEffect, useState } from "react";
import { onAuthStateChanged, signOut, User } from "firebase/auth";

export default function Header() {
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const unsub = onAuthStateChanged(auth, setUser);
    return () => unsub();
  }, []);

  return (
    <header className="border-b mb-8 text-white bg-blue-400">
      <div className="max-w-5xl mx-auto px-6 py-4 flex justify-between items-center">
        {/* Left */}
        <div>
          <h1 className="text-xl font-bold">Code RAG Assistant</h1>
          <p className="text-sm text-black">
            Ask questions about any GitHub repository using AI
          </p>
        </div>

        {/* Right */}
        <div className="text-sm">
          {user ? (
            <div className="flex items-center gap-3">
              <span className="text-black">{user.email}</span>
              <button
                onClick={() => signOut(auth)}
                className="border px-3 py-1 rounded hover:bg-gray-100"
              >
                Logout
              </button>
            </div>
          ) : (
            <span className="text-black">Not logged in</span>
          )}
        </div>
      </div>
    </header>
  );
}
