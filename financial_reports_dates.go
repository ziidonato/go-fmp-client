package go_fmp

import (
	"encoding/json"
	"fmt"
)

// FinancialReportsDatesParams represents the parameters for the Financial Reports Dates API
type FinancialReportsDatesParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// FinancialReportsDatesResponse represents the response from the Financial Reports Dates API
type FinancialReportsDatesResponse struct {
	Symbol     string `json:"symbol"`
	FiscalYear int    `json:"fiscalYear"`
	Period     string `json:"period"`
	LinkXlsx   string `json:"linkXlsx"`
	LinkJson   string `json:"linkJson"`
}

// FinancialReportsDates retrieves financial reports dates for a specific stock symbol
func (c *Client) FinancialReportsDates(params FinancialReportsDatesParams) ([]FinancialReportsDatesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/financial-reports-dates", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []FinancialReportsDatesResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
