type SampleCardProps = {
  title: string;
  description: string;
  actionLabel?: string;
};

/**
 * Renders a simple callout card for the frontend landing page.
 */
export function SampleCard({
  title,
  description,
  actionLabel = "Learn more",
}: SampleCardProps) {
  return (
    <article className="rounded-3xl border border-zinc-200 bg-white p-8 shadow-sm">
      <p className="text-sm font-medium uppercase tracking-[0.24em] text-zinc-500">
        Sample component
      </p>
      <h1 className="mt-4 text-3xl font-semibold tracking-tight text-zinc-950">
        {title}
      </h1>
      <p className="mt-4 max-w-prose text-base leading-7 text-zinc-600">
        {description}
      </p>
      <button
        type="button"
        className="mt-8 inline-flex rounded-full bg-zinc-950 px-5 py-3 text-sm font-medium text-white transition-colors hover:bg-zinc-800"
      >
        {actionLabel}
      </button>
    </article>
  );
}
