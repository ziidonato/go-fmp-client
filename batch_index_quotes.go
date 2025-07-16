package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BatchIndexQuoteResponse represents the response from the batch index quotes API
type BatchIndexQuoteResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetBatchIndexQuotes retrieves real-time quotes for a wide range of stock indexes
func (c *Client) GetBatchIndexQuotes(short bool) ([]BatchIndexQuoteResponse, error) {
	url := "https://financialmodelingprep.com/stable/batch-index-quotes"

	resp, err := c.doRequest(url, map[string]string{
		"short": strconv.FormatBool(short),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []BatchIndexQuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
