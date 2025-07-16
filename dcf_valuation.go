package go_fmp

import (
	"encoding/json"
	"fmt"
)

// DCFValuationParams represents the parameters for the DCF Valuation API
type DCFValuationParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")

// DCFValuationResponse represents the response from the DCF Valuation API
type DCFValuationResponse struct {
	Symbol     string  `json:"symbol"`
	Date       string  `json:"date"`
	DCF        float64 `json:"dcf"`
	StockPrice float64 `json:"Stock Price"`

// DCFValuation estimates the intrinsic value of a company using discounted cash flow analysis
func (c *Client) DCFValuation(params DCFValuationParams) ([]DCFValuationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]DCFValuationResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
