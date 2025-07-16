package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ESGBenchmarkParams represents the parameters for the ESG Benchmark Comparison API
type ESGBenchmarkParams struct {
	Year string `json:"year"` // Required: Year for benchmark data (e.g., "2023")
}

// ESGBenchmarkResponse represents the response from the ESG Benchmark Comparison API
type ESGBenchmarkResponse struct {
	FiscalYear         int     `json:"fiscalYear"`
	Sector             string  `json:"sector"`
	EnvironmentalScore float64 `json:"environmentalScore"`
	SocialScore        float64 `json:"socialScore"`
	GovernanceScore    float64 `json:"governanceScore"`
	ESGScore           float64 `json:"ESGScore"`
}

// ESGBenchmark retrieves ESG benchmark comparison data for sectors
func (c *Client) ESGBenchmark(params ESGBenchmarkParams) ([]ESGBenchmarkResponse, error) {
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	urlParams := map[string]string{
		"year": params.Year,
	}

	return doRequest[[]ESGBenchmarkResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
