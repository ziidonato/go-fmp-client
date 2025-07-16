package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ExchangeMarketHoursResponse represents the response from the global exchange market hours API
type ExchangeMarketHoursResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Exchange string `json:"exchange"`
	// Add other fields as needed based on actual API response
}

// GetExchangeMarketHours retrieves trading hours for specific stock exchanges
func (c *Client) GetExchangeMarketHours(exchange string) ([]ExchangeMarketHoursResponse, error) {
	if exchange == "" {
		return nil, fmt.Errorf("exchange is required")
	}

	url := "https://financialmodelingprep.com/stable/exchange-market-hours"

	return doRequest[[]ExchangeMarketHoursResponse](c, url, map[string]string{
		"exchange": exchange,
	})
}
