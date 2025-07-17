package go_fmp

import (
	"time"
	"fmt"
)

// HistoricalIndustryPerformanceResponse represents the response from the historical industry performance API
type HistoricalIndustryPerformanceResponse struct {
	Date time.Time `json:"date"`
	Industry      string  `json:"industry"`
	Exchange Exchange `json:"exchange"`
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

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}
	if exchange != "" {
		params["exchange"] = exchange
	}

	url := c.BaseURL + "/historical-industry-performance"

	var result []HistoricalIndustryPerformanceResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
