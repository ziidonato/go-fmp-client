package go_fmp

import (
	"fmt"
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

	url := c.BaseURL + "/insider-trading/statistics"

	var result []InsiderTradingStatisticsResponse
	err := c.doRequest(url, map[string]string{
		"symbol": symbol,
	}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
