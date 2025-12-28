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
	t.AddStep(xray.Step{
		Name:      "keyword_generation",
		Input:     keywordInput,
		Output:    keywordOutput,
		Reasoning: "Extracted material, insulation, and size attributes from title and category",
	})
}

func step2CandidateSearch(t *xray.Tracer) {
	t.AddStep(xray.Step{
		Name:      "candidate_search",
		Input:     candidateSearchInput,
		Output:    candidateSearchOutput,
		Reasoning: "Fetched top candidates ranked by search relevance",
	})
}

func step3ApplyFilters(t *xray.Tracer) {
	t.AddStep(xray.Step{
		Name:      "apply_filters",
		Input:     map[string]interface{}{"reference_product": referenceProduct},
		Filters:   filterRules,
		Evals:     filterEvals,
		Output:    filterOutput,
		Reasoning: "Applied price and rating filters to eliminate non-competitive products",
	})
}

func step4LLMRelevanceEvaluation(t *xray.Tracer) {
	t.AddStep(xray.Step{
		Name: "llm_relevance_evaluation",
		Input: map[string]interface{}{
			"candidates_count":  2,
			"reference_product": referenceProductWithCategory,
			"model":             "gpt-4o-mini",
		},
		Evals:     relevanceEvals,
		Output:    relevanceOutput,
		Reasoning: "LLM labeled accessories as false positives and kept only true competitors",
	})
}

func step5RankAndSelect(t *xray.Tracer) {
	t.AddStep(xray.Step{
		Name: "rank_and_select",
		Input: map[string]interface{}{
			"candidates_count":  2,
			"reference_product": referenceProduct,
		},
		Output: map[string]interface{}{
			"ranking_criteria":  rankingCriteria,
			"ranked_candidates": rankedCandidates,
			"selection":         rankingSelection,
		},
		Reasoning: "Ranked by reviews, rating, and price proximity; selected HydroFlask as best competitor",
	})
}
