package go_fmp

import (
	"time"
	"fmt"
)

// SectorPerformanceSnapshotResponse represents the response from the market sector performance snapshot API
type SectorPerformanceSnapshotResponse struct {
	Date time.Time `json:"date"`
	Sector        string  `json:"sector"`
	Exchange      string  `json:"exchange"`
	AverageChange float64 `json:"averageChange"`
}

// GetSectorPerformanceSnapshot retrieves a snapshot of sector performance
func (c *Client) GetSectorPerformanceSnapshot(date, exchange, sector string) ([]SectorPerformanceSnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date is required")
	}

	params := map[string]string{
		"date": date,
	}

	if exchange != "" {
		params["exchange"] = exchange
	}
	if sector != "" {
		params["sector"] = sector
	}

	url := c.BaseURL + "/sector-performance-snapshot"

	var result []SectorPerformanceSnapshotResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
