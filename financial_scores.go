package go_fmp

import (
	"fmt"
)

// FinancialScoresParams represents the parameters for the Financial Scores API
type FinancialScoresParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// FinancialScoresResponse represents the response from the Financial Scores API
type FinancialScoresResponse struct {
	Symbol           string  `json:"symbol"`
	ReportedCurrency string  `json:"reportedCurrency"`
	AltmanZScore     float64 `json:"altmanZScore"`
	PiotroskiScore   int     `json:"piotroskiScore"`
	WorkingCapital   int64   `json:"workingCapital"`
	TotalAssets      int64   `json:"totalAssets"`
	RetainedEarnings int64   `json:"retainedEarnings"`
	EBIT             int64   `json:"ebit"`
	MarketCap        int64   `json:"marketCap"`
	TotalLiabilities int64   `json:"totalLiabilities"`
	Revenue          int64   `json:"revenue"`
}

// FinancialScores retrieves financial health scores for a specific stock symbol
func (c *Client) FinancialScores(params FinancialScoresParams) ([]FinancialScoresResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []FinancialScoresResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/financial-scores", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
