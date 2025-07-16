package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	resp, err := c.get(url, map[string]string{
		"symbol": symbol,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []AftermarketTradeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
