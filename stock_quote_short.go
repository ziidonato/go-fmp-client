package go_fmp

import (
	"encoding/json"
	"fmt"
)

// StockQuoteShortResponse represents the response from the stock quote short API
type StockQuoteShortResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Change float64 `json:"change"`
	Volume int64   `json:"volume"`
}

// GetStockQuoteShort retrieves quick snapshots of real-time stock quotes
func (c *Client) GetStockQuoteShort(symbol string) ([]StockQuoteShortResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := "https://financialmodelingprep.com/stable/quote-short"

	return doRequest[[]StockQuoteShortResponse](c, url, map[string]string{
		"symbol": symbol,
	})
}
