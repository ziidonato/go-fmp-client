package go_fmp

import (
	"encoding/json"
	"fmt"
)

// AllExchangeMarketHoursResponse represents the response from the all exchange market hours API
type AllExchangeMarketHoursResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Exchange string `json:"exchange"`
	// Add other fields as needed based on actual API response
}

// AllExchangeMarketHours retrieves market hours for all exchanges
func (c *Client) AllExchangeMarketHours() ([]AllExchangeMarketHoursResponse, error) {
	url := "https://financialmodelingprep.com/stable/all-exchange-market-hours"

	return doRequest[[]AllExchangeMarketHoursResponse](c, url, map[string]string{})
}
