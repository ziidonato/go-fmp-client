package go_fmp

import (
	"fmt"
	"strconv"
)

// ETFQuotesResponse represents the response from the ETF quotes API
type ETFQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetETFQuotes retrieves real-time quotes for exchange-traded funds
func (c *Client) GetETFQuotes(short bool) ([]ETFQuotesResponse, error) {
	url := c.BaseURL + "/batch-etf-quotes"
	params := map[string]string{"short": strconv.FormatBool(short)}
	var result []ETFQuotesResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	return result, nil
}
