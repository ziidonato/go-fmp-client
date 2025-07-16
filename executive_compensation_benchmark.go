package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ExecutiveCompensationBenchmarkParams represents the parameters for the Executive Compensation Benchmark API
type ExecutiveCompensationBenchmarkParams struct {
	Year *string `json:"year"` // Required: Year for benchmark data (e.g., "2024")

// ExecutiveCompensationBenchmarkResponse represents the response from the Executive Compensation Benchmark API
type ExecutiveCompensationBenchmarkResponse struct {
	IndustryTitle       string  `json:"industryTitle"`
	Year                int     `json:"year"`
	AverageCompensation float64 `json:"averageCompensation"`

// ExecutiveCompensationBenchmark retrieves average executive compensation data across various industries
func (c *Client) ExecutiveCompensationBenchmark(params ExecutiveCompensationBenchmarkParams) ([]ExecutiveCompensationBenchmarkResponse, error) {
	if params.Year == nil {
		return nil, fmt.Errorf("year parameter is required")
	}

	urlParams := map[string]string{
		"year": *params.Year,
	}

	return doRequest[[]ExecutiveCompensationBenchmarkResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
