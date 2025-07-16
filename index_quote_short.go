package go_fmp

import (
	"encoding/json"
	"fmt"
)

// IndexShortQuoteResponse represents the response from the index short quote API
type IndexShortQuoteResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetIndexShortQuote retrieves concise stock index quotes
func (c *Client) GetIndexShortQuote(symbol string) ([]IndexShortQuoteResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := "https://financialmodelingprep.com/stable/quote-short"

	return doRequest[[]IndexShortQuoteResponse](c, url, map[string]string{
		"symbol": symbol,
	})
}
