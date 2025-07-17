package go_fmp

import (
	"time"
	"fmt"
)

// SectorPESnapshotResponse represents the response from the sector PE snapshot API
type SectorPESnapshotResponse struct {
	Date time.Time `json:"date"`
	Sector   string  `json:"sector"`
	Exchange string  `json:"exchange"`
	PE       float64 `json:"pe"`
}

// GetSectorPESnapshot retrieves the price-to-earnings (P/E) ratios for various sectors
func (c *Client) GetSectorPESnapshot(date, exchange, sector string) ([]SectorPESnapshotResponse, error) {
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

	url := c.BaseURL + "/sector-pe-snapshot"

	var result []SectorPESnapshotResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
