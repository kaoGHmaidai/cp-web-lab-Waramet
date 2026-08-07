"use client";

import { useState } from "react";

type GreetResponse = {
  message: string;
};

/**
 * Small form that asks the Go backend (via the /api/greet route handler)
 * for a greeting and displays the response.
 */
export function GreetForm() {
  const [name, setName] = useState("");
  const [message, setMessage] = useState<string | null>(null);
  const [status, setStatus] = useState<"idle" | "loading" | "error">("idle");

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setStatus("loading");
    setMessage(null);

    try {
      const query = name ? `?name=${encodeURIComponent(name)}` : "";
      const response = await fetch(`/api/greet${query}`);

      if (!response.ok) {
        throw new Error("Request failed");
      }

      const data: GreetResponse = await response.json();
      setMessage(data.message);
      setStatus("idle");
    } catch {
      setStatus("error");
    }
  }

  return (
    <form
      onSubmit={handleSubmit}
      className="rounded-3xl border border-zinc-200 bg-white p-8 shadow-sm"
    >
      <p className="text-sm font-medium uppercase tracking-[0.24em] text-zinc-500">
        Greeting demo
      </p>
      <h1 className="mt-4 text-3xl font-semibold tracking-tight text-zinc-950">
        Say hello
      </h1>
      <label htmlFor="greet-name" className="mt-6 block text-sm text-zinc-600">
        Your name
      </label>
      <input
        id="greet-name"
        name="name"
        type="text"
        value={name}
        onChange={(event) => setName(event.target.value)}
        placeholder="e.g. Ann"
        className="mt-2 w-full rounded-xl border border-zinc-300 px-4 py-2 text-zinc-950"
      />
      <button
        type="submit"
        disabled={status === "loading"}
        className="mt-6 inline-flex rounded-full bg-zinc-950 px-5 py-3 text-sm font-medium text-white transition-colors hover:bg-zinc-800 disabled:opacity-50"
      >
        {status === "loading" ? "Asking backend..." : "Get greeting"}
      </button>

      {message !== null && (
        <p role="status" className="mt-6 text-base text-zinc-700">
          {message}
        </p>
      )}
      {status === "error" && (
        <p role="alert" className="mt-6 text-base text-red-600">
          Something went wrong. Try again.
        </p>
      )}
    </form>
  );
}
