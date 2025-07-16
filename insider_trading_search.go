package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// InsiderTradingSearchResponse represents the response from the search insider trades API
type InsiderTradingSearchResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string `json:"symbol"`
	// Add other fields as needed based on actual API response
}

// GetInsiderTradingSearch searches insider trading activity by company or symbol
func (c *Client) GetInsiderTradingSearch(page, limit int) ([]InsiderTradingSearchResponse, error) {
	if page < 0 {
		return nil, fmt.Errorf("page must be non-negative")
	}
	if limit <= 0 {
		return nil, fmt.Errorf("limit must be positive")
	}

	url := "https://financialmodelingprep.com/stable/insider-trading/search"

	return doRequest[[]InsiderTradingSearchResponse](c, url, map[string]string{
		"page":  strconv.Itoa(page)
}
