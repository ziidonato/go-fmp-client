package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InsiderTradingTransactionTypeResponse represents the response from the all insider transaction types API
type InsiderTradingTransactionTypeResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Type string `json:"type"`
	// Add other fields as needed based on actual API response
}

// GetInsiderTradingTransactionType retrieves a comprehensive list of insider transaction types
func (c *Client) GetInsiderTradingTransactionType() ([]InsiderTradingTransactionTypeResponse, error) {
	url := "https://financialmodelingprep.com/stable/insider-trading-transaction-type"

	resp, err := c.doRequest(url, map[string]string{})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []InsiderTradingTransactionTypeResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
