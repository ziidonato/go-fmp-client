package go_fmp

import (
	"encoding/json"
	"fmt"
)

// PeersBulkResponse represents the response from the Stock Peers Bulk API
type PeersBulkResponse struct {
	Symbol string `json:"symbol"`
	Peers  string `json:"peers"`

// GetPeersBulk retrieves peer companies for all stocks in the database
func (c *Client) GetPeersBulk() ([]PeersBulkResponse, error) {
	return doRequest[[]PeersBulkResponse](c, "https://financialmodelingprep.com/stable/peers-bulk", nil)
