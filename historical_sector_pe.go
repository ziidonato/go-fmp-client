package go_fmp

import (
	"encoding/json"
	"fmt"
)

// HistoricalSectorPEResponse represents the response from the historical sector PE API
type HistoricalSectorPEResponse struct {
	Date     string  `json:"date"`
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

	if from != " {
		params["from"] = from
	}
	if to != " {
		params["to"] = to
	}
	if exchange != " {
		params["exchange"] = exchange
	}

	url := "https://financialmodelingprep.com/stable/historical-sector-pe"

	return doRequest[[]HistoricalSectorPEResponse](c, url, params)
}
