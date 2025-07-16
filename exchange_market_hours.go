package go_fmp

import (
	"fmt"
)

// ExchangeMarketHoursResponse represents the response from the exchange market hours API
type ExchangeMarketHoursResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Exchange string `json:"exchange"`
	// Add other fields as needed based on actual API response
}

// GetExchangeMarketHours retrieves trading hours for a specific stock exchange
func (c *Client) GetExchangeMarketHours(exchange string) ([]ExchangeMarketHoursResponse, error) {
	if exchange == "" {
		return nil, fmt.Errorf("exchange is required")
	}

	url := c.BaseURL + "/exchange-market-hours"
	params := map[string]string{"exchange": exchange}
	var result []ExchangeMarketHoursResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
