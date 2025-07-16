package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BatchQuoteShortResponse represents the response from the stock batch quote short API
type BatchQuoteShortResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetBatchQuoteShort retrieves real-time, short-form quotes for multiple stocks
func (c *Client) GetBatchQuoteShort(symbols string) ([]BatchQuoteShortResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}

	url := "https://financialmodelingprep.com/stable/batch-quote-short"

	resp, err := c.doRequest(url, map[string]string{
		"symbols": symbols,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []BatchQuoteShortResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
