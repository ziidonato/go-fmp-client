package go_fmp

import (
	"fmt"
	"strconv"
)

// TechnicalIndicatorSMAResponse represents the response from the simple moving average API
type TechnicalIndicatorSMAResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string  `json:"symbol"`
	Value  float64 `json:"value"`
	// Add other fields as needed based on actual API response
}

// GetTechnicalIndicatorSMA retrieves simple moving average technical indicator
func (c *Client) GetTechnicalIndicatorSMA(symbol string, periodLength int, timeframe TimeFrame) ([]TechnicalIndicatorSMAResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if periodLength <= 0 {
		return nil, fmt.Errorf("periodLength must be positive")
	}
	if timeframe == "" {
		return nil, fmt.Errorf("timeframe is required")
	}

	url := c.BaseURL + "/technical-indicators/sma"
	params := map[string]string{
		"symbol":       symbol,
		"periodLength": strconv.Itoa(periodLength),
		"timeframe":    timeframe.String(),
	}
	var result []TechnicalIndicatorSMAResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
