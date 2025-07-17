package go_fmp

import (
	"fmt"
)

// BulkUpgradesDowngradesConsensusParams represents the parameters for the Bulk Upgrades Downgrades Consensus API
type BulkUpgradesDowngradesConsensusParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkUpgradesDowngradesConsensusResponse represents the response from the Bulk Upgrades Downgrades Consensus API
type BulkUpgradesDowngradesConsensusResponse struct {
	Symbol              string  `json:"symbol"`
	StrongBuy           float64 `json:"strongBuy"`
	Buy                 float64 `json:"buy"`
	Hold                float64 `json:"hold"`
	Sell                float64 `json:"sell"`
	StrongSell          float64 `json:"strongSell"`
	Consensus           string  `json:"consensus"`
}

// BulkUpgradesDowngradesConsensus retrieves analyst consensus ratings for multiple stocks
func (c *Client) BulkUpgradesDowngradesConsensus(params BulkUpgradesDowngradesConsensusParams) ([]BulkUpgradesDowngradesConsensusResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkUpgradesDowngradesConsensusResponse
	err := c.doRequest(c.BaseURL+"/bulk-upgrades-downgrades-consensus", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk upgrades downgrades consensus: %w", err)
	}

	return result, nil
}
