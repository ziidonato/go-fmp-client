package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InsiderTradingStatisticsResponse represents the response from the insider trade statistics API
type InsiderTradingStatisticsResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string `json:"symbol"`
	// Add other fields as needed based on actual API response
}

// GetInsiderTradingStatistics analyzes insider trading activity for a specific symbol
func (c *Client) GetInsiderTradingStatistics(symbol string) ([]InsiderTradingStatisticsResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := "https://financialmodelingprep.com/stable/insider-trading/statistics"

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

	var result []InsiderTradingStatisticsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
