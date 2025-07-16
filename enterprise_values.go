package go_fmp

import (
	"encoding/json"
	"fmt"
)

// EnterpriseValuesParams represents the parameters for the Enterprise Values API
type EnterpriseValuesParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
	Period string `json:"period"` // Optional: Period type - "Q1,Q2,Q3,Q4,FY,annual,quarter"
}

// EnterpriseValuesResponse represents the response from the Enterprise Values API
type EnterpriseValuesResponse struct {
	Symbol                      string  `json:"symbol"`
	Date                        string  `json:"date"`
	StockPrice                  float64 `json:"stockPrice"`
	NumberOfShares              int64   `json:"numberOfShares"`
	MarketCapitalization        int64   `json:"marketCapitalization"`
	MinusCashAndCashEquivalents int64   `json:"minusCashAndCashEquivalents"`
	AddTotalDebt                int64   `json:"addTotalDebt"`
	EnterpriseValue             int64   `json:"enterpriseValue"`
}

// EnterpriseValues retrieves enterprise value data for a specific stock symbol
func (c *Client) EnterpriseValues(params EnterpriseValuesParams) ([]EnterpriseValuesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		if *params.Limit > 1000 {
			return nil, fmt.Errorf("limit cannot exceed 1000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	if params.Period != "" {
		urlParams["period"] = params.Period
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/enterprise-values", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []EnterpriseValuesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
