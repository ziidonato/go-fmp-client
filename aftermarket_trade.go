package go_fmp

import (
	"encoding/json"
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

	url := "https://financialmodelingprep.com/stable/aftermarket-trade"

	return doRequest[[]AftermarketTradeResponse](c, url, map[string]string{
		"symbol": symbol,
	})
}
