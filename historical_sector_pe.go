package go_fmp

import (
	"fmt"
	"time"
)

// HistoricalSectorPEResponse represents a historical sector P/E ratio entry
type HistoricalSectorPEResponse struct {
	Date     time.Time  `json:"date"`
	Sector   string  `json:"sector"`
	Exchange string  `json:"exchange"`
	PE       float64 `json:"pe"`
}

// GetHistoricalSectorPE retrieves historical price-to-earnings (P/E) ratios for various sectors
func (c *Client) GetHistoricalSectorPE(sector, from, to, exchange string) ([]HistoricalSectorPEResponse, error) {
	if sector == "" {
		return nil, fmt.Errorf("sector is required")
	}

	params := map[string]string{
		"sector": sector,
	}

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}
	if exchange != "" {
		params["exchange"] = exchange
	}

	url := c.BaseURL + "/historical-sector-pe"

	var result []HistoricalSectorPEResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
