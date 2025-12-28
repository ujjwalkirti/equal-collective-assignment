import { TraceCard, type Trace } from "../components/trace-card";
import { Badge } from "../components/ui/badge";
import { Card, CardContent } from "../components/ui/card";
import { Separator } from "../components/ui/separator";

const endpoint = process.env.NEXT_PUBLIC_TRACE_ENDPOINT ?? "http://localhost:8080/trace";

async function fetchTrace(): Promise<Trace> {
  const res = await fetch(endpoint, { cache: "no-store" });
  if (!res.ok) {
    throw new Error(`Failed to fetch trace: ${res.status} ${res.statusText}`);
  }
  return res.json();
}

export default async function Page() {
  let trace: Trace | null = null;
  let error: string | null = null;

  try {
    trace = await fetchTrace();
  } catch (err) {
    error = err instanceof Error ? err.message : "Unable to load trace";
  }

  return (
    <main className="mx-auto flex max-w-6xl flex-col gap-8 px-6 py-10 md:py-14">
      <header className="space-y-6 rounded-3xl border border-white/10 bg-canvas-subtle/70 p-8 shadow-glow">
        <div className="inline-flex items-center gap-2 rounded-full border border-white/5 bg-white/5 px-4 py-2 text-xs font-semibold uppercase tracking-wide text-accent-mint">
          <span>Equal Collective</span>
          <Separator className="h-4 w-px" />
          <span>X-Ray Dashboard</span>
        </div>
        <div className="grid gap-4 lg:grid-cols-[2fr,1fr] lg:items-center">
          <div className="space-y-3">
            <h1 className="text-4xl font-semibold leading-tight text-white md:text-5xl">
              See every decision your pipeline makes.
            </h1>
            <p className="text-base text-slate-300 md:text-lg">
              Inspect inputs, filters, evaluations, and reasoning across every step. Built for non-deterministic workflows
              like keyword generation, candidate search, and filter stages.
            </p>
            <div className="flex flex-wrap gap-3">
              <Badge variant="accent">Next.js + Tailwind + shadcn</Badge>
              <Badge variant="muted">Live pull: {endpoint}</Badge>
            </div>
          </div>
          <Card className="border-white/5 bg-white/5">
            <CardContent className="space-y-3 text-sm text-slate-200">
              <div className="flex items-center justify-between">
                <span>Trace Source</span>
                <span className="font-semibold text-white">Go demo service</span>
              </div>
              <div className="flex items-center justify-between">
                <span>Refresh</span>
                <span>Server fetch on load (no cache)</span>
              </div>
              <div className="flex items-center justify-between">
                <span>Environment</span>
                <span className="font-semibold text-accent-mint">{process.env.NODE_ENV ?? "development"}</span>
              </div>
            </CardContent>
          </Card>
        </div>
      </header>

      {error && (
        <Card className="border-coral-500/40 bg-coral-500/10">
          <CardContent className="text-sm text-coral-100">{error}</CardContent>
        </Card>
      )}

      {trace && <TraceCard trace={trace} />}
    </main>
  );
}
