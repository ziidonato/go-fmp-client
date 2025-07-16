package go_fmp

import (
	"encoding/json"
	"fmt"
)

// FundDisclosureHoldersLatestParams represents the parameters for the Mutual Fund & ETF Disclosure API
type FundDisclosureHoldersLatestParams struct {
	Symbol string `json:"symbol"` // Required: Symbol (e.g., "AAPL")

// FundDisclosureHoldersLatestResponse represents the response from the Mutual Fund & ETF Disclosure API
type FundDisclosureHoldersLatestResponse struct {
	CIK           string  `json:"cik"`
	Holder        string  `json:"holder"`
	Shares        int64   `json:"shares"`
	DateReported  string  `json:"dateReported"`
	Change        int64   `json:"change"`
	WeightPercent float64 `json:"weightPercent"`

// FundDisclosureHoldersLatest retrieves the latest disclosures from mutual funds and ETFs
func (c *Client) FundDisclosureHoldersLatest(params FundDisclosureHoldersLatestParams) ([]FundDisclosureHoldersLatestResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]FundDisclosureHoldersLatestResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
