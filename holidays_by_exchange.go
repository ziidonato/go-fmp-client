package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HolidaysByExchangeResponse represents the response from the holidays by exchange API
type HolidaysByExchangeResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Exchange string `json:"exchange"`
	// Add other fields as needed based on actual API response
}

// GetHolidaysByExchange retrieves holidays for specific stock exchanges
func (c *Client) GetHolidaysByExchange(exchange string) ([]HolidaysByExchangeResponse, error) {
	if exchange == "" {
		return nil, fmt.Errorf("exchange is required")
	}

	url := "https://financialmodelingprep.com/stable/holidays-by-exchange"

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

	var result []HolidaysByExchangeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
