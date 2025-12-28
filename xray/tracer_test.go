package xray

import "testing"

func TestNewTracerInitializesTrace(t *testing.T) {
	tracer := NewTracer("test-run-001")

	trace := tracer.Trace()

	if trace.RunID != "test-run-001" {
		t.Fatalf("expected run_id 'test-run-001', got %s", trace.RunID)
	}

	if trace.StartedAt.IsZero() {
		t.Fatalf("expected StartedAt to be set")
	}

	if len(trace.Steps) != 0 {
		t.Fatalf("expected no steps initially, got %d", len(trace.Steps))
	}
}

func TestAddStepAppendsStep(t *testing.T) {
	tracer := NewTracer("test-run")

	step := Step{
		Name: "test_step",
		Input: map[string]interface{}{
			"foo": "bar",
		},
		Reasoning: "testing step addition",
	}

	tracer.AddStep(step)

	trace := tracer.Trace()

	if len(trace.Steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(trace.Steps))
	}

	if trace.Steps[0].Name != "test_step" {
		t.Fatalf("unexpected step name: %s", trace.Steps[0].Name)
	}
}

func TestStepsPreserveOrder(t *testing.T) {
	tracer := NewTracer("order-test")

	tracer.AddStep(Step{Name: "step-1", Reasoning: "first"})
	tracer.AddStep(Step{Name: "step-2", Reasoning: "second"})
	tracer.AddStep(Step{Name: "step-3", Reasoning: "third"})

	trace := tracer.Trace()

	if len(trace.Steps) != 3 {
		t.Fatalf("expected 3 steps, got %d", len(trace.Steps))
	}

	if trace.Steps[2].Name != "step-3" {
		t.Fatalf("expected last step to be 'step-3', got %s", trace.Steps[2].Name)
	}
}

func TestStepWithEvaluationsAndFilters(t *testing.T) {
	tracer := NewTracer("complex-step-test")

	step := Step{
		Name: "apply_filters",
		Filters: map[string]FilterRule{
			"min_rating": {
				Rule: ">= 3.8",
				Min:  3.8,
			},
		},
		Evals: []CandidateEvaluation{
			{
				ID:    "A1",
				Title: "Good Product",
				Metrics: map[string]interface{}{
					"rating": 4.5,
				},
				FilterResults: map[string]FilterCheck{
					"min_rating": {
						Passed: true,
						Detail: "4.5 >= 3.8",
					},
				},
				Qualified: true,
			},
			{
				ID:    "B1",
				Title: "Bad Product",
				Metrics: map[string]interface{}{
					"rating": 3.2,
				},
				FilterResults: map[string]FilterCheck{
					"min_rating": {
						Passed: false,
						Detail: "3.2 < 3.8",
					},
				},
				Qualified: false,
			},
		},
		Output: map[string]interface{}{
			"passed": 1,
			"failed": 1,
		},
		Reasoning: "Filtered candidates by minimum rating",
	}

	tracer.AddStep(step)

	trace := tracer.Trace()

	if len(trace.Steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(trace.Steps))
	}

	if len(trace.Steps[0].Evals) != 2 {
		t.Fatalf("expected 2 candidate evaluations, got %d", len(trace.Steps[0].Evals))
	}
}
