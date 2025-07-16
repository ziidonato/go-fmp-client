package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// TechnicalIndicatorDEMAResponse represents the response from the double exponential moving average API
type TechnicalIndicatorDEMAResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Symbol string  `json:"symbol"`
	Value  float64 `json:"value"`
	// Add other fields as needed based on actual API response
}

// GetTechnicalIndicatorDEMA retrieves double exponential moving average technical indicator
func (c *Client) GetTechnicalIndicatorDEMA(symbol string, periodLength int, timeframe string) ([]TechnicalIndicatorDEMAResponse, error) {
	if symbol == "" {
		return nil, fmt.Errorf("symbol is required")
	}
	if periodLength <= 0 {
		return nil, fmt.Errorf("periodLength must be positive")
	}
	if timeframe == "" {
		return nil, fmt.Errorf("timeframe is required")
	}

	url := "https://financialmodelingprep.com/stable/technical-indicators/dema"

	resp, err := c.doRequest(url, map[string]string{
		"symbol":       symbol,
		"periodLength": strconv.Itoa(periodLength),
		"timeframe":    timeframe,
	})
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()


	var result []TechnicalIndicatorDEMAResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
