package go_fmp

import (
	"encoding/json"
)

// FinancialStatementSymbolListResponse represents the response from the Financial Statement Symbols List API
type FinancialStatementSymbolListResponse struct {
	Symbol            string `json:"symbol"`
	CompanyName       string `json:"companyName"`
	TradingCurrency   string `json:"tradingCurrency"`
	ReportingCurrency string `json:"reportingCurrency"`

// FinancialStatementSymbolList retrieves a comprehensive list of companies with available financial statements
func (c *Client) FinancialStatementSymbolList() ([]FinancialStatementSymbolListResponse, error) {
	return doRequest[[]FinancialStatementSymbolListResponse](c, "https://financialmodelingprep.com/stable/financial-statement-symbol-list", nil)
