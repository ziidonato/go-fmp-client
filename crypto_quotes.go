package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// CryptoQuotesResponse represents the response from the full cryptocurrency quotes API
type CryptoQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetCryptoQuotes retrieves real-time cryptocurrency quotes
func (c *Client) GetCryptoQuotes(short bool) ([]CryptoQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-crypto-quotes"

	resp, err := c.get(url, map[string]string{
		"short": strconv.FormatBool(short),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []CryptoQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
