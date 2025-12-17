"use client";

import { useState } from "react";
import { auth } from "@/lib/firebase";
import {
  createUserWithEmailAndPassword,
  signInWithEmailAndPassword,
} from "firebase/auth";

export default function AuthBox() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [msg, setMsg] = useState("");

  async function signup() {
    try {
      await createUserWithEmailAndPassword(auth, email, password);
      setMsg("Signup successful");
    } catch (err: any) {
      setMsg(err.message);
    }
  }

  async function login() {
    try {
      await signInWithEmailAndPassword(auth, email, password);
      setMsg("Login successful");
    } catch (err: any) {
      setMsg(err.message);
    }
  }

  return (
    <div className="border p-4 rounded mb-6">
      <h2 className="font-semibold mb-2">Login / Signup</h2>

      <input
        className="w-full border p-2 mb-2"
        placeholder="Email"
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        className="w-full border p-2 mb-2"
        placeholder="Password"
        type="password"
        onChange={(e) => setPassword(e.target.value)}
      />

      <div className="flex gap-2">
        <button onClick={login} className="bg-black text-white px-4 py-2 rounded">
          Login
        </button>
        <button onClick={signup} className="border px-4 py-2 rounded">
          Signup
        </button>
      </div>

      {msg && <p className="text-sm mt-2">{msg}</p>}
    </div>
  );
}
