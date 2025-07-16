package go_fmp

import (
	"encoding/json"
	"fmt"
)

// EarningsSurprisesBulkParams represents the parameters for the Earnings Surprises Bulk API
type EarningsSurprisesBulkParams struct {
	Year string `json:"year"` // Required: year (e.g., "2024")
}

// EarningsSurprisesBulkResponse represents the response from the Earnings Surprises Bulk API
type EarningsSurprisesBulkResponse struct {
	Symbol       string `json:"symbol"`
	Date         string `json:"date"`
	EPSActual    string `json:"epsActual"`
	EPSEstimated string `json:"epsEstimated"`
	LastUpdated  string `json:"lastUpdated"`
}

// GetEarningsSurprisesBulk retrieves bulk data on annual earnings surprises
func (c *Client) GetEarningsSurprisesBulk(params EarningsSurprisesBulkParams) ([]EarningsSurprisesBulkResponse, error) {
	// Validate required parameters
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	urlParams := map[string]string{
		"year": params.Year,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/earnings-surprises-bulk", urlParams)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()



	// Parse the response
	var result []EarningsSurprisesBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
