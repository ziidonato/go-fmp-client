package go_fmp

import (
	"fmt"
)

// SectorPESnapshotResponse represents the response from the Sector PE Snapshot API
type SectorPESnapshotResponse struct {
	Date     string  `json:"date"`
	Sector   string  `json:"sector"`
	Exchange string  `json:"exchange"`
	PE       float64 `json:"pe"`
}

// GetSectorPESnapshot retrieves a snapshot of sector price-to-earnings ratios
func (c *Client) GetSectorPESnapshot(date, exchange, sector string) ([]SectorPESnapshotResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date parameter is required")
	}

	url := fmt.Sprintf("%s/sectors-pe/%s", c.BaseURL, date)
	params := map[string]string{}

	if exchange != "" {
		params["exchange"] = exchange
	}

	if sector != "" {
		params["sector"] = sector
	}

	var result []SectorPESnapshotResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
