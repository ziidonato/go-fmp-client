package go_fmp

import (
	"encoding/json"
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
	url := "https://financialmodelingprep.com/stable/batch-mutualfund-quotes"

	resp, err := c.doRequest(url, map[string]string{
		"short": strconv.FormatBool(short),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []MutualFundQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
