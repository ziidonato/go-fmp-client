package go_fmp

import (
	"encoding/json"
	"fmt"
)

// IndustryPerformanceSnapshotResponse represents the response from the industry performance snapshot API
type IndustryPerformanceSnapshotResponse struct {
	Date          string  `json:"date"`
	Industry      string  `json:"industry"`
	Exchange      string  `json:"exchange"`
	AverageChange float64 `json:"averageChange"`
}

// IndustryPerformanceSnapshot retrieves detailed performance data by industry
func (c *Client) IndustryPerformanceSnapshot(date, exchange, industry string) ([]IndustryPerformanceSnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date is required")
	}

	params := map[string]string{
		"date": date,
	}

	if exchange != "" {
		params["exchange"] = exchange
	}
	if industry != "" {
		params["industry"] = industry
	}

	url := "https://financialmodelingprep.com/stable/industry-performance-snapshot"

	resp, err := c.doRequest(url, params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []IndustryPerformanceSnapshotResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
