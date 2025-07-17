package go_fmp

import (
	"fmt"
)

// IndustryPerformanceSnapshotResponse represents the response from the Industry Performance Snapshot API
type IndustryPerformanceSnapshotResponse struct {
	Date          string  `json:"date"`
	Industry      string  `json:"industry"`
	AverageChange float64 `json:"averageChange"`
}

// IndustryPerformanceSnapshot retrieves industry performance snapshot for a specific date
func (c *Client) IndustryPerformanceSnapshot(date, exchange, industry string) ([]IndustryPerformanceSnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date parameter is required")
	}

	url := fmt.Sprintf("%s/industries-performance/%s", c.BaseURL, date)
	params := map[string]string{}

	if exchange != "" {
		params["exchange"] = exchange
	}

	if industry != "" {
		params["industry"] = industry
	}

	var result []IndustryPerformanceSnapshotResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}

	return result, nil
}
