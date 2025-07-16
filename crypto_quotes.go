package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CryptoQuotesResponse represents the response from the full cryptocurrency quotes API
type CryptoQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// CryptoQuotes retrieves real-time cryptocurrency quotes
func (c *Client) CryptoQuotes(short bool) ([]CryptoQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-crypto-quotes"

	resp, err := c.doRequest(url, map[string]string{
		"short": strconv.FormatBool(short),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []CryptoQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
