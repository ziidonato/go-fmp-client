package go_fmp

import (
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
	params := map[string]string{"symbol": symbol}
	var result []StockQuoteShortResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
