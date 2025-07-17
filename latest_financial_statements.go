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
	urlParams := map[string]string{}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
	}

	if params.Limit != nil {
		if *params.Limit > 250 {
			return nil, fmt.Errorf("limit cannot exceed 250")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []LatestFinancialStatementsResponse
	err := c.doRequest(c.BaseURL+"/latest-financial-statements", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
