package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BatchAftermarketTradeResponse represents the response from the batch aftermarket trade API
type BatchAftermarketTradeResponse struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	TradeSize int64   `json:"tradeSize"`
	Timestamp int64   `json:"timestamp"`

// GetBatchAftermarketTrade retrieves real-time aftermarket trading data for multiple stocks
func (c *Client) GetBatchAftermarketTrade(symbols string) ([]BatchAftermarketTradeResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}

	url := "https://financialmodelingprep.com/stable/batch-aftermarket-trade"

	return doRequest[[]BatchAftermarketTradeResponse](c, url, map[string]string{
		"symbols": symbols,
	}
