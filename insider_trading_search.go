package go_fmp

import (
	"fmt"
)

// InsiderTradingSearchResponse represents the response from the insider trading search API
type InsiderTradingSearchResponse struct {
	Symbol         string  `json:"symbol"`
	FilingDate     string  `json:"filingDate"`
	TransactionDate string `json:"transactionDate"`
	ReportingCik   string  `json:"reportingCik"`
	TransactionType string `json:"transactionType"`
	SecuritiesOwned float64 `json:"securitiesOwned"`
	CompanyName    string  `json:"companyName"`
	ReportingName  string  `json:"reportingName"`
}

// GetInsiderTradingSearch searches for insider trading transactions based on symbol
func (c *Client) GetInsiderTradingSearch(symbol string) ([]InsiderTradingSearchResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := "https://financialmodelingprep.com/stable/insider-trading-search"
	params := map[string]string{"symbol": symbol}
	var result []InsiderTradingSearchResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
