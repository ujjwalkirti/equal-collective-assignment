package xray

import "time"

type Trace struct {
	RunID     string    `json:"run_id"`
	StartedAt time.Time `json:"started_at"`
	Steps     []Step    `json:"steps"`
}

type Step struct {
	Name       string                 `json:"step"`
	StepID     string                 `json:"step_id,omitempty"`
	StartedAt  time.Time              `json:"started_at,omitempty"`
	DurationMs int64                  `json:"duration_ms,omitempty"`
	Input      map[string]interface{} `json:"input,omitempty"`
	Filters    map[string]FilterRule  `json:"filters_applied,omitempty"`
	Evals      []CandidateEvaluation  `json:"evaluations,omitempty"`
	Output     map[string]interface{} `json:"output,omitempty"`
	Reasoning  string                 `json:"reasoning"`
}

type FilterRule struct {
	Rule string      `json:"rule"`
	Min  interface{} `json:"min,omitempty"`
	Max  interface{} `json:"max,omitempty"`
}

type CandidateEvaluation struct {
	ID            string                 `json:"id"`
	Title         string                 `json:"title"`
	Metrics       map[string]interface{} `json:"metrics"`
	FilterResults map[string]FilterCheck `json:"filter_results"`
	Qualified     bool                   `json:"qualified"`
}

type FilterCheck struct {
	Passed bool   `json:"passed"`
	Detail string `json:"detail"`
}
