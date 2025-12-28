package xray

import (
	"sync"
	"time"
)

type Tracer struct {
	mu    sync.RWMutex
	trace Trace
}

func NewTracer(runID string) *Tracer {
	return &Tracer{
		trace: Trace{
			RunID:     runID,
			StartedAt: time.Now(),
			Steps:     []Step{},
		},
	}
}

func (t *Tracer) AddStep(step Step) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.trace.Steps = append(t.trace.Steps, step)
}

func (t *Tracer) Trace() Trace {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return cloneTrace(t.trace)
}

// RecordStep wraps step creation with timing metadata and reduces boilerplate.
// The configure function can set inputs, outputs, filters, evaluations, and reasoning.
// StepID defaults to the step name but can be overridden inside configure.
func (t *Tracer) RecordStep(name string, configure func(*Step)) {
	started := time.Now()
	step := Step{
		Name:      name,
		StepID:    name,
		StartedAt: started,
	}
	configure(&step)
	step.DurationMs = time.Since(started).Milliseconds()
	t.AddStep(step)
}

func cloneTrace(src Trace) Trace {
	copy := Trace{
		RunID:     src.RunID,
		StartedAt: src.StartedAt,
		Steps:     make([]Step, 0, len(src.Steps)),
	}
	for _, s := range src.Steps {
		copy.Steps = append(copy.Steps, cloneStep(s))
	}
	return copy
}

func cloneStep(s Step) Step {
	return Step{
		Name:       s.Name,
		StepID:     s.StepID,
		StartedAt:  s.StartedAt,
		DurationMs: s.DurationMs,
		Input:      cloneMap(s.Input),
		Filters:    cloneFilters(s.Filters),
		Evals:      cloneEvals(s.Evals),
		Output:     cloneMap(s.Output),
		Reasoning:  s.Reasoning,
	}
}

func cloneMap(src map[string]interface{}) map[string]interface{} {
	if src == nil {
		return nil
	}
	cp := make(map[string]interface{}, len(src))
	for k, v := range src {
		cp[k] = v
	}
	return cp
}

func cloneFilters(src map[string]FilterRule) map[string]FilterRule {
	if src == nil {
		return nil
	}
	cp := make(map[string]FilterRule, len(src))
	for k, v := range src {
		cp[k] = v
	}
	return cp
}

func cloneEvals(src []CandidateEvaluation) []CandidateEvaluation {
	if src == nil {
		return nil
	}
	cp := make([]CandidateEvaluation, len(src))
	for i, ev := range src {
		cp[i] = CandidateEvaluation{
			ID:            ev.ID,
			Title:         ev.Title,
			Metrics:       cloneMap(ev.Metrics),
			FilterResults: cloneFilterChecks(ev.FilterResults),
			Qualified:     ev.Qualified,
		}
	}
	return cp
}

func cloneFilterChecks(src map[string]FilterCheck) map[string]FilterCheck {
	if src == nil {
		return nil
	}
	cp := make(map[string]FilterCheck, len(src))
	for k, v := range src {
		cp[k] = v
	}
	return cp
}
