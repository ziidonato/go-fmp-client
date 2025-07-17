package go_fmp

import (
	"fmt"
)

// PriceTargetConsensusParams represents the parameters for the Price Target Consensus API
type PriceTargetConsensusParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// PriceTargetConsensusResponse represents the response from the Price Target Consensus API
type PriceTargetConsensusResponse struct {
	Symbol          string  `json:"symbol"`
	TargetHigh      float64 `json:"targetHigh"`
	TargetLow       float64 `json:"targetLow"`
	TargetConsensus float64 `json:"targetConsensus"`
	TargetMedian    float64 `json:"targetMedian"`
}

// PriceTargetConsensus retrieves analysts' consensus price targets for stocks
func (c *Client) PriceTargetConsensus(params PriceTargetConsensusParams) ([]PriceTargetConsensusResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []PriceTargetConsensusResponse
	err := c.doRequest(c.BaseURL+"/price-target-consensus", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
