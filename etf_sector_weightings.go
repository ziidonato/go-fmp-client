package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ETFSectorWeightingsParams represents the parameters for the ETF Sector Weighting API
type ETFSectorWeightingsParams struct {
	Symbol string `json:"symbol"` // Required: ETF/Fund symbol (e.g., "SPY")
}

// ETFSectorWeightingsResponse represents the response from the ETF Sector Weighting API
type ETFSectorWeightingsResponse struct {
	Symbol           string  `json:"symbol"`
	Sector           string  `json:"sector"`
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

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/etf/sector-weightings", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ETFSectorWeightingsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
