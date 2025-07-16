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

	resp, err := c.doRequest(url, map[string]string{
		"symbol": symbol,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []StockQuoteShortResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
