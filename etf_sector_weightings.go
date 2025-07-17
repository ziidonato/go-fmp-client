package go_fmp

import (
	"fmt"
)

// ETFSectorWeightingsParams represents the parameters for the ETF Sector Weighting API
type ETFSectorWeightingsParams struct {
	Symbol string `json:"symbol"` // Required: ETF/Fund symbol (e.g., "SPY")
}

// ETFSectorWeightingsResponse represents the response from the ETF Sector Weighting API
type ETFSectorWeightingsResponse struct {
	Symbol           string  `json:"symbol"`
	Sector Sector `json:"sector"`
	WeightPercentage float64 `json:"weightPercentage"`
}

// ETFSectorWeightings retrieves sector weighting breakdown for ETFs
func (c *Client) ETFSectorWeightings(params ETFSectorWeightingsParams) ([]ETFSectorWeightingsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []ETFSectorWeightingsResponse
	err := c.doRequest(c.BaseURL+"/etf/sector-weightings", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
