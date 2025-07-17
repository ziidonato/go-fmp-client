package go_fmp

import (
	"fmt"
	"time"
)

// LatestFinancialStatementsParams represents the parameters for the Latest Financial Statements API
type LatestFinancialStatementsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Type   string `json:"type"`   // Required: Financial statement type - either "income", "balance", or "cash"
}

// LatestFinancialStatementsResponse represents the response from the Latest Financial Statements API
type LatestFinancialStatementsResponse struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	Date         time.Time `json:"date"`
	DateAdded    time.Time `json:"dateAdded"`
	InformationS map[string]interface{} `json:"informationS"`
}

// LatestFinancialStatements retrieves the latest financial statements data
func (c *Client) LatestFinancialStatements(params LatestFinancialStatementsParams) ([]LatestFinancialStatementsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Type == "" {
		return nil, fmt.Errorf("type parameter is required")
	}

	if params.Type != "income" && params.Type != "balance" && params.Type != "cash" {
		return nil, fmt.Errorf("type must be either 'income', 'balance', or 'cash'")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
		"type":   params.Type,
	}

	var result []LatestFinancialStatementsResponse
	err := c.doRequest(c.BaseURL+"/latest-financial-statements", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
