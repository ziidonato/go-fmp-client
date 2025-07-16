package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SectorPerformanceSnapshotResponse represents the response from the market sector performance snapshot API
type SectorPerformanceSnapshotResponse struct {
	Date          string  `json:"date"`
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

	url := "https://financialmodelingprep.com/stable/sector-performance-snapshot"

	resp, err := c.doRequest(url, params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []SectorPerformanceSnapshotResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
