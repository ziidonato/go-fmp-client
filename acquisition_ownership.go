package go_fmp

import (
	"fmt"
)

// AcquisitionOwnershipResponse represents the response from the acquisition ownership API
type AcquisitionOwnershipResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string `json:"symbol"`
	// Add other fields as needed based on actual API response
}

// AcquisitionOwnership tracks changes in stock ownership during acquisitions
func (c *Client) AcquisitionOwnership(symbol string) ([]AcquisitionOwnershipResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := c.BaseURL + "/acquisition-of-beneficial-ownership"

	var result []AcquisitionOwnershipResponse
	if err := c.doRequest(url, map[string]string{"symbol": symbol}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
