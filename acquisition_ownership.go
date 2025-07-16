package go_fmp

import (
	"encoding/json"
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

	return c.doRequest[[]AcquisitionOwnershipResponse]("https://financialmodelingprep.com/stable/acquisition-of-beneficial-ownership", map[string]string{
		"symbol": symbol,
	})
}
