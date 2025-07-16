package go_fmp

import (
	"encoding/json"
	"fmt"
)

// PeersBulkResponse represents the response from the Stock Peers Bulk API
type PeersBulkResponse struct {
	Symbol string `json:"symbol"`
	Peers  string `json:"peers"`
}

// GetPeersBulk retrieves peer companies for all stocks in the database
func (c *Client) GetPeersBulk() ([]PeersBulkResponse, error) {
	resp, err := c.doRequest("https://financialmodelingprep.com/stable/peers-bulk", nil)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result []PeersBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
