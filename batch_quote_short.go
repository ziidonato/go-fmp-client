package go_fmp

import (
	"fmt"
)

// BatchQuoteShortResponse represents the response from the batch quote short API
type BatchQuoteShortResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetBatchQuoteShort retrieves quick snapshots of real-time stock quotes for multiple stocks
func (c *Client) GetBatchQuoteShort(symbols string) ([]BatchQuoteShortResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}

	url := "https://financialmodelingprep.com/stable/batch-quote-short"
	params := map[string]string{"symbols": symbols}
	var result []BatchQuoteShortResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
