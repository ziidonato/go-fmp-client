package go_fmp

import (
	"fmt"
	"strconv"
)

// MutualFundQuotesResponse represents the response from the mutual fund price quotes API
type MutualFundQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetMutualFundQuotes retrieves real-time quotes for mutual funds
func (c *Client) GetMutualFundQuotes(short bool) ([]MutualFundQuotesResponse, error) {
	url := c.BaseURL + "/batch-mutualfund-quotes"

	urlParams := map[string]string{
		"short": strconv.FormatBool(short),
	}

	var result []MutualFundQuotesResponse
	err := c.doRequest(url, urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
