package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SectorPESnapshotResponse represents the response from the sector PE snapshot API
type SectorPESnapshotResponse struct {
	Date     string  `json:"date"`
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

	url := "https://financialmodelingprep.com/stable/sector-pe-snapshot"

	resp, err := c.doRequest(url, params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []SectorPESnapshotResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
