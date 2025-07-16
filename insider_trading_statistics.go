package go_fmp

import (
	"encoding/json"
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

	url := "https://financialmodelingprep.com/stable/insider-trading/statistics"

	return doRequest[[]InsiderTradingStatisticsResponse](c, url, map[string]string{
		"symbol": symbol,
	})
}
