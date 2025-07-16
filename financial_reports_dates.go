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

	return doRequest[[]FinancialReportsDatesResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
