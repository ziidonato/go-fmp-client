package go_fmp

import (
	"encoding/json"
	"fmt"
)

// HistoricalIndustryPerformanceResponse represents the response from the historical industry performance API
type HistoricalIndustryPerformanceResponse struct {
	Date          string  `json:"date"`
	Industry      string  `json:"industry"`
	Exchange      string  `json:"exchange"`
	AverageChange float64 `json:"averageChange"`
}

// HistoricalIndustryPerformance retrieves historical performance data for industries
func (c *Client) HistoricalIndustryPerformance(industry, from, to, exchange string) ([]HistoricalIndustryPerformanceResponse, error) {
	if industry == "" {
		return nil, fmt.Errorf("industry is required")
	}

	params := map[string]string{
		"industry": industry,
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

	url := "https://financialmodelingprep.com/stable/historical-industry-performance"

	return doRequest[[]HistoricalIndustryPerformanceResponse](c, url, params)
}
