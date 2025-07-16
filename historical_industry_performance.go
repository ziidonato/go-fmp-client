package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HistoricalIndustryPerformanceResponse represents the response from the historical industry performance API
type HistoricalIndustryPerformanceResponse struct {
	Date          string  `json:"date"`
	Industry      string  `json:"industry"`
	Exchange      string  `json:"exchange"`
	AverageChange float64 `json:"averageChange"`
}

// GetHistoricalIndustryPerformance retrieves historical performance data for industries
func (c *Client) GetHistoricalIndustryPerformance(industry, from, to, exchange string) ([]HistoricalIndustryPerformanceResponse, error) {
	if industry == "" {
		return nil, fmt.Errorf("industry is required")
	}

	params := map[string]string{
		"industry": industry,
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

	url := "https://financialmodelingprep.com/stable/historical-industry-performance"

	resp, err := c.get(url, params)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []HistoricalIndustryPerformanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
