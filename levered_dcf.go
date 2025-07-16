package go_fmp

import (
	"encoding/json"
	"fmt"
)

// LeveredDCFParams represents the parameters for the Levered DCF API
type LeveredDCFParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// LeveredDCFResponse represents the response from the Levered DCF API
type LeveredDCFResponse struct {
	Symbol     string  `json:"symbol"`
	Date       string  `json:"date"`
	DCF        float64 `json:"dcf"`
	StockPrice float64 `json:"Stock Price"`
}

// LeveredDCF analyzes a company's value incorporating the impact of debt
func (c *Client) LeveredDCF(params LeveredDCFParams) ([]LeveredDCFResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/levered-discounted-cash-flow", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []LeveredDCFResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
