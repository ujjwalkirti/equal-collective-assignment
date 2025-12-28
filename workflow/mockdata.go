package workflow

import "github.com/equal-collective-assignment/xray"

var (
	keywordInput = map[string]interface{}{
		"product_title": "Stainless Steel Water Bottle 32oz Insulated",
		"category":      "Sports & Outdoors",
	}

	keywordOutput = map[string]interface{}{
		"keywords": []string{
			"stainless steel water bottle",
			"insulated bottle 32oz",
		},
	}

	candidateSearchInput = map[string]interface{}{
		"keyword": "stainless steel water bottle",
		"limit":   5,
	}

	candidateSearchOutput = map[string]interface{}{
		"total_results":      2847,
		"candidates_fetched": 5,
	}

	referenceProduct = map[string]interface{}{
		"title":   "ProBrand Steel Bottle 32oz Insulated",
		"price":   29.99,
		"rating":  4.2,
		"reviews": 1247,
	}

	filterRules = map[string]xray.FilterRule{
		"price_range": {
			Rule: "0.5x - 2x of reference price",
			Min:  14.99,
			Max:  59.98,
		},
		"min_rating": {
			Rule: "Must be at least 3.8 stars",
			Min:  3.8,
		},
	}

	filterEvals = []xray.CandidateEvaluation{
		{
			ID:    "B0COMP01",
			Title: "HydroFlask 32oz Wide Mouth",
			Metrics: map[string]interface{}{
				"price":   44.99,
				"rating":  4.5,
				"reviews": 8932,
			},
			FilterResults: map[string]xray.FilterCheck{
				"price_range": {
					Passed: true,
					Detail: "$44.99 is within $14.99 - $59.98",
				},
				"min_rating": {
					Passed: true,
					Detail: "4.5 ≥ 3.8",
				},
			},
			Qualified: true,
		},
		{
			ID:    "B0COMP02",
			Title: "Generic Water Bottle",
			Metrics: map[string]interface{}{
				"price":   8.99,
				"rating":  3.2,
				"reviews": 45,
			},
			FilterResults: map[string]xray.FilterCheck{
				"price_range": {
					Passed: false,
					Detail: "$8.99 is below $14.99 minimum",
				},
				"min_rating": {
					Passed: false,
					Detail: "3.2 < 3.8",
				},
			},
			Qualified: false,
		},
	}

	filterOutput = map[string]interface{}{
		"passed": 1,
		"failed": 1,
	}

	referenceProductWithCategory = map[string]interface{}{
		"title":    "ProBrand Steel Bottle 32oz Insulated",
		"category": "Sports & Outdoors > Water Bottles",
	}

	relevanceEvals = []xray.CandidateEvaluation{
		{
			ID:    "B0COMP01",
			Title: "HydroFlask 32oz Wide Mouth",
			Metrics: map[string]interface{}{
				"confidence": 0.95,
			},
			FilterResults: map[string]xray.FilterCheck{
				"relevance": {
					Passed: true,
					Detail: "True competitor; same product type",
				},
			},
			Qualified: true,
		},
		{
			ID:    "B0COMP02",
			Title: "Generic Water Bottle",
			Metrics: map[string]interface{}{
				"confidence": 0.78,
			},
			FilterResults: map[string]xray.FilterCheck{
				"relevance": {
					Passed: true,
					Detail: "Still a bottle; lower confidence",
				},
			},
			Qualified: true,
		},
		{
			ID:    "B0COMP05",
			Title: "Replacement Lid for HydroFlask",
			Metrics: map[string]interface{}{
				"confidence": 0.97,
			},
			FilterResults: map[string]xray.FilterCheck{
				"relevance": {
					Passed: false,
					Detail: "Accessory, not a competitor",
				},
			},
			Qualified: false,
		},
	}

	relevanceOutput = map[string]interface{}{
		"total_evaluated":         3,
		"confirmed_competitors":   2,
		"false_positives_removed": 1,
	}

	rankingCriteria = map[string]interface{}{
		"primary":   "review_count",
		"secondary": "rating",
		"tertiary":  "price_proximity",
	}

	rankedCandidates = []map[string]interface{}{
		{
			"rank":  1,
			"id":    "B0COMP01",
			"title": "HydroFlask 32oz Wide Mouth",
			"metrics": map[string]interface{}{
				"price":   44.99,
				"rating":  4.5,
				"reviews": 8932,
			},
			"score_breakdown": map[string]interface{}{
				"review_count_score":    1.0,
				"rating_score":          0.9,
				"price_proximity_score": 0.7,
				"total_score":           0.92,
			},
		},
		{
			"rank":  2,
			"id":    "B0COMP02",
			"title": "Generic Water Bottle",
			"metrics": map[string]interface{}{
				"price":   8.99,
				"rating":  3.2,
				"reviews": 45,
			},
			"score_breakdown": map[string]interface{}{
				"review_count_score":    0.01,
				"rating_score":          0.3,
				"price_proximity_score": 0.3,
				"total_score":           0.20,
			},
		},
	}

	rankingSelection = map[string]interface{}{
		"id":     "B0COMP01",
		"title":  "HydroFlask 32oz Wide Mouth",
		"reason": "Top review count with strong rating and acceptable price proximity",
	}
)
