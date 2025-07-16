package go_fmp

import (
	"fmt"
	"strconv"
)

// TechnicalIndicatorWilliamsResponse represents the response from the Williams API
type TechnicalIndicatorWilliamsResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string  `json:"symbol"`
	Value  float64 `json:"value"`
	// Add other fields as needed based on actual API response
}

// GetTechnicalIndicatorWilliams retrieves Williams %R technical indicator
func (c *Client) GetTechnicalIndicatorWilliams(symbol string, periodLength int, timeframe string) ([]TechnicalIndicatorWilliamsResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if periodLength <= 0 {
		return nil, fmt.Errorf("periodLength must be positive")
	}
	if timeframe == "" {
		return nil, fmt.Errorf("timeframe is required")
	}

	url := "https://financialmodelingprep.com/stable/technical-indicators/williams"
	params := map[string]string{
		"symbol":       symbol,
		"periodLength": strconv.Itoa(periodLength),
		"timeframe":    timeframe,
	}
	var result []TechnicalIndicatorWilliamsResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
