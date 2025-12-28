package xray

import (
	"sync"
	"time"
)

type Tracer struct {
	mu    sync.Mutex
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
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.trace
}
