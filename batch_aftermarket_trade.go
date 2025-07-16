package go_fmp

import "fmt"

// BatchAftermarketTradeResponse represents the response from the batch aftermarket trade API
type BatchAftermarketTradeResponse struct {
	Symbol    string  `json:"symbol"`
	Price     float64 `json:"price"`
	TradeSize int64   `json:"tradeSize"`
	Timestamp int64   `json:"timestamp"`
}

// GetBatchAftermarketTrade retrieves real-time aftermarket trading data for multiple stocks
func (c *Client) GetBatchAftermarketTrade(symbols string) ([]BatchAftermarketTradeResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}

	url := c.BaseURL + "/batch-aftermarket-trade"

	var result []BatchAftermarketTradeResponse
	err := c.doRequest(url, map[string]string{
		"symbols": symbols,
	}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
