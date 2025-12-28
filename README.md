# X-Ray Demo (Go + Next.js Dashboard)

An end-to-end demo of an X-Ray tracing SDK for multi-step decision pipelines. The Go service emits trace data; a Next.js 15 dashboard visualizes runs.

## Prerequisites
- Go 1.22+ (adjust if you use a different toolchain)
- Node 18+ (or latest LTS)

## Backend (Go)
- Run the demo server: `make run`
- Build a Windows binary: `make build` (outputs `bin/xray-demo.exe`)
- Tests: `make test`
- Trace endpoint: `http://localhost:8080/trace` (JSON)

### What it does
The Go app runs a mocked multi-step flow and records into the X-Ray tracer:
1) keyword_generation
2) candidate_search
3) apply_filters
4) llm_relevance_evaluation
5) rank_and_select

Each step records inputs, outputs, filters, evaluations, and reasoning. The tracer holds one in-memory run (`demo-run-001`).

## Frontend Dashboard (Next.js 15 + Tailwind)
- Location: `dashboard/`
- Install deps: `cd dashboard && npm install`
- Run dev server: `npm run dev`
- Build: `npm run build`
- Env: set `NEXT_PUBLIC_TRACE_ENDPOINT` if your trace URL differs (default `http://localhost:8080/trace`).
- Open: `http://localhost:3000`

### Dashboard features
- Live fetch (no cache) of the trace JSON
- Timeline of steps with inputs/outputs, filters, evaluations, and reasoning
- Nested object rendering inside cards for richer context

## Project Layout
- `main.go` – demo workflow + HTTP handler
- `xray/` – tracer and types
- `dashboard/` – Next.js app (components, app router)
- `Makefile` – run/test/build helpers

## Notes / Limitations
- Traces are in-memory; restarting the Go service resets data.
- Single-run demo (`demo-run-001`); extend tracer to persist multiple runs if needed.
- Data is mocked/deterministic to showcase the UX.

## Quickstart
1) In repo root: `make run`
2) In `dashboard/`: `npm install && npm run dev`
3) Visit `http://localhost:3000` and explore the steps.

## Troubleshooting
- If `replaceAll` type errors in TS: ensure TS lib targets ES2021 (tsconfig already set).
- If dashboard can’t fetch: set `NEXT_PUBLIC_TRACE_ENDPOINT` and ensure Go server is running on that host/port.
