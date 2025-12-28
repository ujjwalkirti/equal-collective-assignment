package workflow

import "github.com/equal-collective-assignment/xray"

// RunDemo executes the mocked multi-step pipeline to populate the tracer.
func RunDemo(t *xray.Tracer) {
	step1KeywordGeneration(t)
	step2CandidateSearch(t)
	step3ApplyFilters(t)
	step4LLMRelevanceEvaluation(t)
	step5RankAndSelect(t)
}

func step1KeywordGeneration(t *xray.Tracer) {
	t.RecordStep("keyword_generation", func(s *xray.Step) {
		s.StepID = "step-1"
		s.Input = keywordInput
		s.Output = keywordOutput
		s.Reasoning = "Extracted material, insulation, and size attributes from title and category"
	})
}

func step2CandidateSearch(t *xray.Tracer) {
	t.RecordStep("candidate_search", func(s *xray.Step) {
		s.StepID = "step-2"
		s.Input = candidateSearchInput
		s.Output = candidateSearchOutput
		s.Reasoning = "Fetched top candidates ranked by search relevance"
	})
}

func step3ApplyFilters(t *xray.Tracer) {
	t.RecordStep("apply_filters", func(s *xray.Step) {
		s.StepID = "step-3"
		s.Input = map[string]interface{}{"candidates_count": 50, "reference_product": referenceProduct}
		s.Filters = filterRules
		s.Evals = filterEvals
		s.Output = filterOutput
		s.Reasoning = "Applied price and rating filters to eliminate non-competitive products"
	})
}

func step4LLMRelevanceEvaluation(t *xray.Tracer) {
	t.RecordStep("llm_relevance_evaluation", func(s *xray.Step) {
		s.StepID = "step-4"
		s.Input = map[string]interface{}{
			"candidates_count":  2,
			"reference_product": referenceProductWithCategory,
			"model":             "gpt-4o-mini",
		}
		s.Evals = relevanceEvals
		s.Output = relevanceOutput
		s.Reasoning = "LLM labeled accessories as false positives and kept only true competitors"
	})
}

func step5RankAndSelect(t *xray.Tracer) {
	t.RecordStep("rank_and_select", func(s *xray.Step) {
		s.StepID = "step-5"
		s.Input = map[string]interface{}{
			"candidates_count":  2,
			"reference_product": referenceProduct,
		}
		s.Output = map[string]interface{}{
			"ranking_criteria":  rankingCriteria,
			"ranked_candidates": rankedCandidates,
			"selection":         rankingSelection,
		}
		s.Reasoning = "Ranked by reviews, rating, and price proximity; selected HydroFlask as best competitor"
	})
}
