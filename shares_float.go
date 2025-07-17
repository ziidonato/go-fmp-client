package go_fmp

import (
	"fmt"
	"time"
)

// SharesFloatParams represents the parameters for the Shares Float API
type SharesFloatParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// SharesFloatResponse represents the response from the Shares Float API
type SharesFloatResponse struct {
	Symbol               string    `json:"symbol"`
	Date                 time.Time `json:"date"`
	FloatShares          float64   `json:"floatShares"`
	OutstandingShares    float64   `json:"outstandingShares"`
	PercentageOfFloat    float64   `json:"percentageOfFloat"`
}

// SharesFloat retrieves the total number of publicly traded shares for any company
func (c *Client) SharesFloat(params SharesFloatParams) ([]SharesFloatResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []SharesFloatResponse
	err := c.doRequest(c.BaseURL+"/shares-float", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
