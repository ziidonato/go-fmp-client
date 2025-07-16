package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AcquisitionOwnershipResponse represents the response from the acquisition ownership API
type AcquisitionOwnershipResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string `json:"symbol"`
	// Add other fields as needed based on actual API response
}

// GetAcquisitionOwnership tracks changes in stock ownership during acquisitions
func (c *Client) GetAcquisitionOwnership(symbol string) ([]AcquisitionOwnershipResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}

	url := "https://financialmodelingprep.com/stable/acquisition-of-beneficial-ownership"

	resp, err := c.get(url, map[string]string{
		"symbol": symbol,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	var result []AcquisitionOwnershipResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
