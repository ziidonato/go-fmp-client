package go_fmp

import (
	"fmt"
)

// IndexQuoteShortResponse represents the response from the index quote short API
type IndexQuoteShortResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetIndexQuoteShort retrieves quick snapshots of real-time index quotes
func (c *Client) GetIndexQuoteShort(symbol string) ([]IndexQuoteShortResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := c.BaseURL + "/index-quote-short"
	params := map[string]string{"symbol": symbol}
	var result []IndexQuoteShortResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
