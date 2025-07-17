package go_fmp

import (
	"fmt"
)

// BulkScoresParams represents the parameters for the Bulk Scores API
type BulkScoresParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkScoresResponse represents the response from the Bulk Scores API
type BulkScoresResponse struct {
	Symbol                  string  `json:"symbol"`
	AltmanZScore            float64 `json:"altmanZScore"`
	PiotroskiScore          float64 `json:"piotroskiScore"`
	WorkingCapital          float64 `json:"workingCapital"`
	TotalAssets             float64 `json:"totalAssets"`
	RetainedEarnings        float64 `json:"retainedEarnings"`
	EBIT                    float64 `json:"ebit"`
	MarketCap               float64 `json:"marketCap"`
	TotalLiabilities        float64 `json:"totalLiabilities"`
	Revenue                 float64 `json:"revenue"`
}

// BulkScores retrieves financial health scores for multiple stocks
func (c *Client) BulkScores(params BulkScoresParams) ([]BulkScoresResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkScoresResponse
	err := c.doRequest(c.BaseURL+"/bulk-scores", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk scores: %w", err)
	}

	return result, nil
}
