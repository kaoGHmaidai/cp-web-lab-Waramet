import { GreetForm } from "@/components/greet-form";

export default function GreetPage() {
  return (
    <div className="flex min-h-screen items-center justify-center bg-zinc-50 px-6 py-16 font-sans text-zinc-950">
      <main className="w-full max-w-xl">
        <GreetForm />
      </main>
    </div>
  );
}
