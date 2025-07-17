package go_fmp

import (
	"time"
)

import "fmt"

// AnalystEstimatesParams represents the parameters for the Financial Estimates API
type AnalystEstimatesParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period Period `json:"period"` // Required: Period type - "annual" or "quarter"
	Page   *int   `json:"page"`   // Optional: Page number (default: 0)
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
}

// AnalystEstimatesResponse represents the response from the Financial Estimates API
type AnalystEstimatesResponse struct {
	Symbol             string  `json:"symbol"`
	Date time.Time `json:"date"`
	RevenueLow         int64   `json:"revenueLow"`
	RevenueHigh        int64   `json:"revenueHigh"`
	RevenueAvg         int64   `json:"revenueAvg"`
	EBITDALow          int64   `json:"ebitdaLow"`
	EBITDAHigh         int64   `json:"ebitdaHigh"`
	EBITDAAvg          int64   `json:"ebitdaAvg"`
	EBITLow            int64   `json:"ebitLow"`
	EBITHigh           int64   `json:"ebitHigh"`
	EBITAvg            int64   `json:"ebitAvg"`
	NetIncomeLow       int64   `json:"netIncomeLow"`
	NetIncomeHigh      int64   `json:"netIncomeHigh"`
	NetIncomeAvg       int64   `json:"netIncomeAvg"`
	SGAExpenseLow      int64   `json:"sgaExpenseLow"`
	SGAExpenseHigh     int64   `json:"sgaExpenseHigh"`
	SGAExpenseAvg      int64   `json:"sgaExpenseAvg"`
	EPSAvg             float64 `json:"epsAvg"`
	EPSHigh            float64 `json:"epsHigh"`
	EPSLow             float64 `json:"epsLow"`
	NumAnalystsRevenue int     `json:"numAnalystsRevenue"`
	NumAnalystsEPS     int     `json:"numAnalystsEps"`
}

// AnalystEstimates retrieves analyst financial estimates for stock symbols
func (c *Client) AnalystEstimates(params AnalystEstimatesParams) ([]AnalystEstimatesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	if params.Period != "annual" && params.Period != "quarter" {
		return nil, fmt.Errorf("period must be either 'annual' or 'quarter'")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
		"period": string(params.Period),
	}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
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
