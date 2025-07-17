package go_fmp

import (
	"fmt"
	"time"
)

// ETFHoldingsParams represents the parameters for the ETF Holdings API
type ETFHoldingsParams struct {
	Symbol string  `json:"symbol"` // Required: ETF symbol (e.g., "SPY")
	Page   *int    `json:"page"`   // Optional: Page number (default: 0)
}

// ETFHoldingsResponse represents the response from the ETF Holdings API
type ETFHoldingsResponse struct {
	Symbol           string    `json:"symbol"`
	Asset            string    `json:"asset"`
	Name             string    `json:"name"`
	CUSIP            string    `json:"cusip"`
	ISIN             string    `json:"isin"`
	SharesNumber     int64     `json:"sharesNumber"`
	WeightPercentage float64   `json:"weightPercentage"`
	MarketValue      float64   `json:"marketValue"`
	UpdatedAt        time.Time `json:"updatedAt"`
	Updated          string    `json:"updated"`
}

// ETFHoldings retrieves a detailed breakdown of the assets held within ETFs and mutual funds
func (c *Client) ETFHoldings(params ETFHoldingsParams) ([]ETFHoldingsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []ETFHoldingsResponse
	err := c.doRequest(c.BaseURL+"/etf/holdings", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
