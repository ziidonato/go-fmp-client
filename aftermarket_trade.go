package go_fmp

import (
	"fmt"
)

// AftermarketTradeResponse represents the response from the aftermarket trade API
type AftermarketTradeResponse struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	TradeSize int64   `json:"tradeSize"`
	Timestamp int64   `json:"timestamp"`
}

// AftermarketTrade retrieves real-time aftermarket trading activity
func (c *Client) AftermarketTrade(symbol string) ([]AftermarketTradeResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := c.BaseURL + "/aftermarket-trade"
	params := map[string]string{"symbol": symbol}
	var result []AftermarketTradeResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
