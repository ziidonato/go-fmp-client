package go_fmp

import (
	"fmt"
)

// BulkPeersParams represents the parameters for the Bulk Peers API
type BulkPeersParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkPeersResponse represents the response from the Bulk Peers API
type BulkPeersResponse struct {
	Symbol string   `json:"symbol"`
	Peers  []string `json:"peers"`
}

// BulkPeers retrieves peer companies for multiple stocks
func (c *Client) BulkPeers(params BulkPeersParams) ([]BulkPeersResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkPeersResponse
	err := c.doRequest(c.BaseURL+"/bulk-peers", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk peers: %w", err)
	}

	return result, nil
}
