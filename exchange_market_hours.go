package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	resp, err := c.get(url, map[string]string{
		"exchange": exchange,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []ExchangeMarketHoursResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
