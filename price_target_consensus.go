package go_fmp

import (
	"encoding/json"
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

	resp, err := c.get("https://financialmodelingprep.com/stable/price-target-consensus", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []PriceTargetConsensusResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
