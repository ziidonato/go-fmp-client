package go_fmp

import (
	"encoding/json"
	"fmt"
)

// HistoricalSectorPerformanceResponse represents the response from the historical market sector performance API
type HistoricalSectorPerformanceResponse struct {
	Date          string  `json:"date"`
	Sector        string  `json:"sector"`
	Exchange      string  `json:"exchange"`
	AverageChange float64 `json:"averageChange"`
}

// HistoricalSectorPerformance retrieves historical sector performance data
func (c *Client) HistoricalSectorPerformance(sector, from, to, exchange string) ([]HistoricalSectorPerformanceResponse, error) {
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

	url := "https://financialmodelingprep.com/stable/historical-sector-performance"

	return doRequest[[]HistoricalSectorPerformanceResponse](c, url, params)
}
