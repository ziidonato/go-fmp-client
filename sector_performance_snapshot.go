package go_fmp

import (
	"fmt"
)

// SectorPerformanceSnapshotResponse represents the response from the Sector Performance Snapshot API
type SectorPerformanceSnapshotResponse struct {
	Date                         string  `json:"date"`
	Sector                       string  `json:"sector"`
	ChangePercentage             string  `json:"changePercentage"`
}

// GetSectorPerformanceSnapshot retrieves sector performance snapshot for a specific date
func (c *Client) GetSectorPerformanceSnapshot(date, exchange, sector string) ([]SectorPerformanceSnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date parameter is required")
	}

	url := fmt.Sprintf("%s/sectors-performance/%s", c.BaseURL, date)
	params := map[string]string{}

	if exchange != "" {
		params["exchange"] = exchange
	}

	if sector != "" {
		params["sector"] = sector
	}

	var result []SectorPerformanceSnapshotResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
