package go_fmp

import (
	"time"
	"fmt"
)

// IndustryPerformanceSnapshotResponse represents the response from the industry performance snapshot API
type IndustryPerformanceSnapshotResponse struct {
	Date time.Time `json:"date"`
	Industry      string  `json:"industry"`
	Exchange Exchange `json:"exchange"`
	AverageChange float64 `json:"averageChange"`
}

// IndustryPerformanceSnapshot retrieves detailed performance data by industry
func (c *Client) IndustryPerformanceSnapshot(date string, exchange Exchange, industry string) ([]IndustryPerformanceSnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date is required")
	}
	if exchange == "" {
		return nil, fmt.Errorf("exchange is required")
	}
	if industry == "" {
		return nil, fmt.Errorf("industry is required")
	}

	url := c.BaseURL + "/industry-performance-snapshot"
	params := map[string]string{
		"date":     date,
		"exchange": string(exchange),
		"industry": industry,
	}
	var result []IndustryPerformanceSnapshotResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
