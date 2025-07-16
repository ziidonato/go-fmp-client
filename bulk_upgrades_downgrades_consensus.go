package go_fmp

import (
	"encoding/json"
	"fmt"
)

// UpgradesDowngradesConsensusBulkResponse represents the response from the Upgrades Downgrades Consensus Bulk API
type UpgradesDowngradesConsensusBulkResponse struct {
	Symbol     string `json:"symbol"`
	StrongBuy  string `json:"strongBuy"`
	Buy        string `json:"buy"`
	Hold       string `json:"hold"`
	Sell       string `json:"sell"`
	StrongSell string `json:"strongSell"`
	Consensus  string `json:"consensus"`
}

// UpgradesDowngradesConsensusBulk retrieves analyst ratings across all symbols
func (c *Client) UpgradesDowngradesConsensusBulk() ([]UpgradesDowngradesConsensusBulkResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/upgrades-downgrades-consensus-bulk", nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()



	// Parse the response
	var result []UpgradesDowngradesConsensusBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
