package go_fmp

import (
	"fmt"
	"time"
)

// AnalystEstimatesParams represents the parameters for the Analyst Estimates API
type AnalystEstimatesParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// AnalystEstimatesResponse represents the response from the Analyst Estimates API
type AnalystEstimatesResponse struct {
	Symbol            string    `json:"symbol"`
	Date              time.Time `json:"date"`
	EstimatedRevenueLow   float64 `json:"estimatedRevenueLow"`
	EstimatedRevenueHigh  float64 `json:"estimatedRevenueHigh"`
	EstimatedRevenueAvg   float64 `json:"estimatedRevenueAvg"`
	EstimatedEbitdaLow    float64 `json:"estimatedEbitdaLow"`
	EstimatedEbitdaHigh   float64 `json:"estimatedEbitdaHigh"`
	EstimatedEbitdaAvg    float64 `json:"estimatedEbitdaAvg"`
	EstimatedEbitLow      float64 `json:"estimatedEbitLow"`
	EstimatedEbitHigh     float64 `json:"estimatedEbitHigh"`
	EstimatedEbitAvg      float64 `json:"estimatedEbitAvg"`
	EstimatedNetIncomeLow float64 `json:"estimatedNetIncomeLow"`
	EstimatedNetIncomeHigh float64 `json:"estimatedNetIncomeHigh"`
	EstimatedNetIncomeAvg  float64 `json:"estimatedNetIncomeAvg"`
	EstimatedSgaExpenseLow float64 `json:"estimatedSgaExpenseLow"`
	EstimatedSgaExpenseHigh float64 `json:"estimatedSgaExpenseHigh"`
	EstimatedSgaExpenseAvg  float64 `json:"estimatedSgaExpenseAvg"`
	EstimatedEpsAvg      float64 `json:"estimatedEpsAvg"`
	EstimatedEpsLow      float64 `json:"estimatedEpsLow"`
	EstimatedEpsHigh     float64 `json:"estimatedEpsHigh"`
	NumberAnalystEstimatedRevenue float64 `json:"numberAnalystEstimatedRevenue"`
	NumberAnalystsEstimatedEps float64 `json:"numberAnalystsEstimatedEps"`
}

// AnalystEstimates retrieves analyst financial estimates for stock symbols
func (c *Client) AnalystEstimates(params AnalystEstimatesParams) ([]AnalystEstimatesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Period == nil || (*params.Period != "annual" && *params.Period != "quarter") {
		return nil, fmt.Errorf("period parameter is required and must be either 'annual' or 'quarter'")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
		"period": *params.Period,
	}

	if params.Limit != nil {
		if *params.Limit > 1000 {
			return nil, fmt.Errorf("limit cannot exceed 1000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []AnalystEstimatesResponse
	err := c.doRequest(c.BaseURL+"/analyst-estimates", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
