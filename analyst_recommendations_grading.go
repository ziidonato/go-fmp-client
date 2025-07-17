package go_fmp

import (
	"fmt"
	"time"
)

// AnalystRecommendationsGradingParams represents the parameters for the Analyst Recommendations Grading API
type AnalystRecommendationsGradingParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// AnalystRecommendationsGradingResponse represents the response from the Analyst Recommendations Grading API
type AnalystRecommendationsGradingResponse struct {
	Symbol                     string    `json:"symbol"`
	PublishingDate             time.Time `json:"publishingDate"`
	GradingCompany             string    `json:"gradingCompany"`
	Recommendation             Rating    `json:"recommendation"`
	PreviousRecommendation     Rating    `json:"previousRecommendation"`
	AnalystName                string    `json:"analystName"`
	PriceTarget                float64   `json:"priceTarget"`
}

// AnalystRecommendationsGrading retrieves analyst recommendation grades with upgrades and downgrades
func (c *Client) AnalystRecommendationsGrading(params AnalystRecommendationsGradingParams) ([]AnalystRecommendationsGradingResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []AnalystRecommendationsGradingResponse
	err := c.doRequest(c.BaseURL+"/grade", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}