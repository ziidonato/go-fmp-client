package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ETFQuotesResponse represents the response from the ETF price quotes API
type ETFQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetETFQuotes retrieves real-time price quotes for exchange-traded funds (ETFs)
func (c *Client) GetETFQuotes(short bool) ([]ETFQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-etf-quotes"

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

	var result []ETFQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
