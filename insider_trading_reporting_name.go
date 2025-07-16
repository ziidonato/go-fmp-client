package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InsiderTradingReportingNameResponse represents the response from the search insider trades by reporting name API
type InsiderTradingReportingNameResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Name string `json:"name"`
	// Add other fields as needed based on actual API response
}

// InsiderTradingReportingName searches for insider trading activity by reporting name
func (c *Client) InsiderTradingReportingName(name string) ([]InsiderTradingReportingNameResponse, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	url := "https://financialmodelingprep.com/stable/insider-trading/reporting-name"

	return doRequest[[]InsiderTradingReportingNameResponse](c, url, map[string]string{
		"name": name,
	})
}
