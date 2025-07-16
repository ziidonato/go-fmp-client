package go_fmp

import (
	"encoding/json"
	"fmt"
)

// HolidaysByExchangeResponse represents the response from the holidays by exchange API
type HolidaysByExchangeResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Exchange string `json:"exchange"`
	// Add other fields as needed based on actual API response

// HolidaysByExchange retrieves holidays for specific stock exchanges
func (c *Client) HolidaysByExchange(exchange string) ([]HolidaysByExchangeResponse, error) {
	if exchange == "" {
		return nil, fmt.Errorf("exchange is required")
	}

	url := "https://financialmodelingprep.com/stable/holidays-by-exchange"

	return doRequest[[]HolidaysByExchangeResponse](c, url, map[string]string{
		"exchange": exchange,
	}
