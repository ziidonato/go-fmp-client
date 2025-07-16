package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// InsiderTradingReportingNameResponse represents the response from the search insider trades by reporting name API
type InsiderTradingReportingNameResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Name string `json:"name"`
	// Add other fields as needed based on actual API response
}

// GetInsiderTradingReportingName searches for insider trading activity by reporting name
func (c *Client) GetInsiderTradingReportingName(name string) ([]InsiderTradingReportingNameResponse, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	url := "https://financialmodelingprep.com/stable/insider-trading/reporting-name"

	resp, err := c.get(url, map[string]string{
		"name": name,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []InsiderTradingReportingNameResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
