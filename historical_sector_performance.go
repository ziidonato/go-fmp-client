package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}
	if exchange != "" {
		params["exchange"] = exchange
	}

	url := "https://financialmodelingprep.com/stable/historical-sector-performance"

	resp, err := c.get(url, params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []HistoricalSectorPerformanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
