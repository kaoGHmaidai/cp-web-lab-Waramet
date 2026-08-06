import { SampleCard } from "@/components/sample-card";

export default function Home() {
  return (
    <div className="flex min-h-screen items-center justify-center bg-zinc-50 px-6 py-16 font-sans text-zinc-950">
      <main className="w-full max-w-xl">
        <SampleCard
          title="Frontend starter"
          description="A small, documented component with a Jest test to validate the frontend test setup."
          actionLabel="Review sample"
        />
      </main>
    </div>
  );
}
