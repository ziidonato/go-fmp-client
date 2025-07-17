package go_fmp

import (
	"fmt"
	"strconv"
)

// TechnicalIndicatorADXResponse represents the response from the average directional index API
type TechnicalIndicatorADXResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string  `json:"symbol"`
	Value  float64 `json:"value"`
	// Add other fields as needed based on actual API response
}

// GetTechnicalIndicatorADX retrieves average directional index technical indicator
func (c *Client) GetTechnicalIndicatorADX(symbol string, periodLength int, timeframe TimeFrame) ([]TechnicalIndicatorADXResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if periodLength <= 0 {
		return nil, fmt.Errorf("periodLength must be positive")
	}
	if timeframe == "" {
		return nil, fmt.Errorf("timeframe is required")
	}

	url := c.BaseURL + "/technical-indicators/adx"

	var result []TechnicalIndicatorADXResponse
	err := c.doRequest(url, map[string]string{
		"symbol":       symbol,
		"periodLength": strconv.Itoa(periodLength),
		"timeframe":    timeframe.String(),
	}, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
