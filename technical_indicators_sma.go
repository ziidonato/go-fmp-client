package go_fmp

import (
	"encoding/json"
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

// GetTechnicalIndicatorSMA retrieves simple moving average technical indicator
func (c *Client) GetTechnicalIndicatorSMA(symbol string, periodLength int, timeframe string) ([]TechnicalIndicatorSMAResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if periodLength <= 0 {
		return nil, fmt.Errorf("periodLength must be positive")
	}
	if timeframe == "" {
		return nil, fmt.Errorf("timeframe is required")
	}

	url := "https://financialmodelingprep.com/stable/technical-indicators/sma"

	return doRequest[[]TechnicalIndicatorSMAResponse](c, url, map[string]string{
		"symbol":       symbol,
		"periodLength": strconv.Itoa(periodLength)
	}
