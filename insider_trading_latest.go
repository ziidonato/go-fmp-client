package go_fmp

import (
	"fmt"
)

// InsiderTradingLatestResponse represents the response from the latest insider trading API
type InsiderTradingLatestResponse struct {
	Symbol          string  `json:"symbol"`
	FilingDate      string  `json:"filingDate"`
	TransactionDate string  `json:"transactionDate"`
	ReportingCik    string  `json:"reportingCik"`
	TransactionType string  `json:"transactionType"`
	SecuritiesOwned float64 `json:"securitiesOwned"`
	CompanyName     string  `json:"companyName"`
	ReportingName   string  `json:"reportingName"`
}

// GetInsiderTradingLatest retrieves the most recent insider trading transactions for a specific symbol
func (c *Client) GetInsiderTradingLatest(symbol string) ([]InsiderTradingLatestResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := c.BaseURL + "/insider-trading-latest"
	params := map[string]string{"symbol": symbol}
	var result []InsiderTradingLatestResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
