package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// InsiderTradingLatestResponse represents the response from the latest insider trading API
type InsiderTradingLatestResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string `json:"symbol"`
	// Add other fields as needed based on actual API response
}

// GetInsiderTradingLatest retrieves the latest insider trading activity
func (c *Client) GetInsiderTradingLatest(page, limit int) ([]InsiderTradingLatestResponse, error) {
	if page < 0 {
		return nil, fmt.Errorf("page must be non-negative")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be positive")
	}

	url := "https://financialmodelingprep.com/stable/insider-trading/latest"

	resp, err := c.doRequest(url, map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []InsiderTradingLatestResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
