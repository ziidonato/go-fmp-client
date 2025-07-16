package go_fmp

import (
	"fmt"
)

// AftermarketQuoteResponse represents the response from the aftermarket quote API
type AftermarketQuoteResponse struct {
	Symbol    string  `json:"symbol"`
	BidSize   int64   `json:"bidSize"`
	BidPrice  float64 `json:"bidPrice"`
	AskSize   int64   `json:"askSize"`
	AskPrice  float64 `json:"askPrice"`
	Volume    int64   `json:"volume"`
	Timestamp int64   `json:"timestamp"`
}

// GetAftermarketQuote retrieves real-time aftermarket quotes for stocks
func (c *Client) GetAftermarketQuote(symbol string) ([]AftermarketQuoteResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := "https://financialmodelingprep.com/stable/aftermarket-quote"
	params := map[string]string{"symbol": symbol}
	var result []AftermarketQuoteResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
