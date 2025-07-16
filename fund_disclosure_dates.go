package go_fmp

import (
	"encoding/json"
	"fmt"
)

// FundDisclosureDatesParams represents the parameters for the Fund & ETF Disclosures by Date API
type FundDisclosureDatesParams struct {
	Symbol string  `json:"symbol"` // Required: Fund/ETF symbol (e.g., "VWO")
	CIK    *string `json:"cik"`    // Optional: CIK number

// FundDisclosureDatesResponse represents the response from the Fund & ETF Disclosures by Date API
type FundDisclosureDatesResponse struct {
	Date    string `json:"date"`
	Year    int    `json:"year"`
	Quarter int    `json:"quarter"`

// FundDisclosureDates retrieves disclosure dates for mutual funds and ETFs
func (c *Client) FundDisclosureDates(params FundDisclosureDatesParams) ([]FundDisclosureDatesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.CIK != nil {
		urlParams["cik"] = *params.CIK
	}

	return doRequest[[]FundDisclosureDatesResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
