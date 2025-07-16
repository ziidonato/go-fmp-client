package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// ForexQuotesResponse represents the response from the full forex quote API
type ForexQuotesResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// ForexQuotes retrieves real-time quotes for multiple forex currency pairs
func (c *Client) ForexQuotes(short bool) ([]ForexQuotesResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-forex-quotes"

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

	var result []ForexQuotesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
