package go_fmp

import (
	"fmt"
	"time"
)

// LeveredDCFParams represents the parameters for the Levered DCF API
type LeveredDCFParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// LeveredDCFResponse represents the response from the Levered DCF API
type LeveredDCFResponse struct {
	Symbol     string    `json:"symbol"`
	Date       time.Time `json:"date"`
	DCF        float64   `json:"dcf"`
	Price      float64   `json:"price"`
}

// LeveredDCF analyzes a company's value incorporating the impact of debt
func (c *Client) LeveredDCF(params LeveredDCFParams) ([]LeveredDCFResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []LeveredDCFResponse
	err := c.doRequest(c.BaseURL+"/levered-discounted-cash-flow", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
