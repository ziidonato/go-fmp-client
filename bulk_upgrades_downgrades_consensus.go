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

// UpgradesDowngradesConsensusBulk retrieves analyst ratings across all symbols
func (c *Client) UpgradesDowngradesConsensusBulk() ([]UpgradesDowngradesConsensusBulkResponse, error) {
	return doRequest[[]UpgradesDowngradesConsensusBulkResponse](c, "https://financialmodelingprep.com/stable/upgrades-downgrades-consensus-bulk", nil)
