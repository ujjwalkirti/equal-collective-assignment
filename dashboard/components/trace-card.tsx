import { ArrowRight, Filter, ListChecks, Search } from "lucide-react";
import { Badge } from "./ui/badge";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import { Separator } from "./ui/separator";

export type Trace = {
  run_id: string;
  started_at: string;
  steps: Step[];
};

export type Step = {
  step: string;
  input?: Record<string, unknown>;
  filters_applied?: Record<string, FilterRule>;
  evaluations?: CandidateEvaluation[];
  output?: Record<string, unknown>;
  reasoning?: string;
};

export type FilterRule = {
  rule?: string;
  min?: unknown;
  max?: unknown;
};

export type CandidateEvaluation = {
  id: string;
  title: string;
  metrics?: Record<string, unknown>;
  filter_results?: Record<string, FilterCheck>;
  qualified?: boolean;
};

export type FilterCheck = {
  passed: boolean;
  detail?: string;
};

const icons: Record<string, JSX.Element> = {
  keyword_generation: <Search className="h-4 w-4" />,
  candidate_search: <ListChecks className="h-4 w-4" />,
  apply_filters: <Filter className="h-4 w-4" />,
};

export function TraceCard({ trace }: { trace: Trace }) {
  return (
    <Card className="border-white/10 bg-canvas-subtle/80 shadow-none">
      <CardHeader>
        <div className="flex flex-wrap items-center gap-3">
          <Badge variant="accent">Run {trace.run_id}</Badge>
          <span className="text-sm text-slate-300">Started {formatDate(trace.started_at)}</span>
        </div>
        <CardTitle className="text-2xl">Decision Trail</CardTitle>
      </CardHeader>
      <CardContent className="space-y-6">
        <div className="relative space-y-6">
          {trace.steps.map((step, idx) => (
            <StepRow key={`${step.step}-${idx}`} step={step} isLast={idx === trace.steps.length - 1} />
          ))}
        </div>
      </CardContent>
    </Card>
  );
}

function StepRow({ step, isLast }: { step: Step; isLast: boolean }) {
  const icon = icons[step.step] ?? <ArrowRight className="h-4 w-4" />;
  return (
    <div className="grid grid-cols-[auto,1fr] gap-4">
      <div className="relative flex flex-col items-center">
        <div className="flex h-9 w-9 items-center justify-center rounded-full border border-white/10 bg-white/5 text-accent-mint shadow-glow">
          {icon}
        </div>
        {!isLast && <div className="mt-1 h-full w-px flex-1 bg-white/10" />}
      </div>
      <div className="rounded-2xl border border-white/5 bg-canvas-subtle/60 p-4 shadow-none">
        <div className="flex flex-wrap items-center gap-2">
          <Badge variant="muted" className="capitalize">{step.step.replaceAll("_", " ")}</Badge>
          {step.reasoning && <span className="text-sm text-slate-300">{step.reasoning}</span>}
        </div>
        <div className="mt-4 grid gap-3 ">
          {step.input && <DataBlock label="Input" data={step.input} />}
          {step.output && <DataBlock label="Output" data={step.output} />}
        </div>
        {step.filters_applied && (
          <div className="mt-4 space-y-2">
            <Label>Filters Applied</Label>
            <div className="grid gap-2 md:grid-cols-2">
              {Object.entries(step.filters_applied).map(([name, rule]) => (
                <div key={name} className="rounded-xl border border-white/5 bg-white/5 p-3 text-sm text-slate-200">
                  <div className="flex items-center justify-between">
                    <span className="font-semibold capitalize">{name.replaceAll("_", " ")}</span>
                    <Badge variant="muted">{rule.rule ?? "rule"}</Badge>
                  </div>
                  <div className="mt-2 text-xs text-slate-400">Min: {formatValue(rule.min)} · Max: {formatValue(rule.max)}</div>
                </div>
              ))}
            </div>
          </div>
        )}
        {step.evaluations && step.evaluations.length > 0 && (
          <div className="mt-4 space-y-2">
            <Label>Evaluations</Label>
            <div className="grid gap-2">
              {step.evaluations.map((ev) => (
                <div key={ev.id} className="rounded-xl border border-white/5 bg-white/5 p-3">
                  <div className="flex flex-wrap items-center gap-2">
                    <Badge variant={ev.qualified ? "success" : "error"}>{ev.qualified ? "Qualified" : "Rejected"}</Badge>
                    <span className="font-semibold text-slate-100">{ev.title}</span>
                    <span className="text-xs text-slate-400">({ev.id})</span>
                  </div>
                  {ev.metrics && <DataGrid data={ev.metrics} />}
                  {ev.filter_results && (
                    <div className="mt-2 grid gap-2 md:grid-cols-2">
                      {Object.entries(ev.filter_results).map(([name, result]) => (
                        <div key={name} className="rounded-lg border border-white/5 bg-canvas-subtle/80 p-2 text-xs text-slate-200">
                          <div className="flex items-center justify-between">
                            <span className="capitalize text-slate-300">{name.replaceAll("_", " ")}</span>
                            <Badge variant={result.passed ? "success" : "error"}>{result.passed ? "pass" : "fail"}</Badge>
                          </div>
                          {result.detail && <div className="mt-1 text-slate-400">{result.detail}</div>}
                        </div>
                      ))}
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

function DataBlock({ label, data }: { label: string; data: Record<string, unknown> }) {
  return (
    <div className="rounded-xl border border-white/5 bg-white/5 p-3">
      <Label>{label}</Label>
      <DataGrid data={data} />
    </div>
  );
}

function DataGrid({ data }: { data: Record<string, unknown> }) {
  return (
    <div className="mt-2 grid gap-2 text-xs text-slate-200">
      {Object.entries(data).map(([key, value]) => (
        <div key={key} className="flex flex-col gap-2">
          <div className="flex font-semibold capitalize text-slate-100">{key.replaceAll("_", " ")}</div>
          <div className="text-slate-300">{formatValue(value)}</div>
        </div>
      ))}
    </div>
  );
}

function Label({ children }: { children: React.ReactNode }) {
  return <div className="text-xs font-semibold uppercase tracking-[0.08em] text-slate-400">{children}</div>;
}

function formatValue(value: unknown): React.ReactNode {
  if (value === null || value === undefined) return "—";
  if (typeof value === "number") return Number.isInteger(value) ? value.toString() : value.toFixed(2);
  if (typeof value === "string") return value;
  if (typeof value === "boolean") return value ? "true" : "false";
  if (Array.isArray(value)) {
    const items = value.map((v, idx) => (
      <span key={idx} className="text-slate-200">
        {formatValue(v)}
        {idx < value.length-1 ? ", " : ""}
      </span>
    ));
    return <div className="grid md:grid-cols-2 gap-3">{items}</div>;
  }
  if (typeof value === "object") return renderObjectCard(value as Record<string, unknown>);
  return String(value);
}

function renderObjectCard(obj: Record<string, unknown>) {
  return (
    <Card className="border-white/5 bg-white/5">
      <CardContent className="space-y-2 text-[11px] text-slate-200">
        {Object.entries(obj).map(([k, v]) => (
          <div key={k} className="flex flex-col gap-1 rounded-md bg-white/5 px-2 py-2">
            <p className="font-semibold capitalize text-slate-100">{k.replaceAll("_", " ")}</p>
            <div className="text-slate-300">{formatValue(v)}</div>
          </div>
        ))}
      </CardContent>
    </Card>
  );
}

function formatDate(iso: string) {
  const date = new Date(iso);
  return date.toLocaleString(undefined, {
    hour: "2-digit",
    minute: "2-digit",
    day: "numeric",
    month: "short",
    year: "numeric",
  });
}
