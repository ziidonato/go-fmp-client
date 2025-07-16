package go_fmp

import (
	"encoding/json"
	"fmt"
)

// RatingsSnapshotParams represents the parameters for the Ratings Snapshot API
type RatingsSnapshotParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (default: 1)

// RatingsSnapshotResponse represents the response from the Ratings Snapshot API
type RatingsSnapshotResponse struct {
	Symbol                  string `json:"symbol"`
	Rating                  string `json:"rating"`
	OverallScore            int    `json:"overallScore"`
	DiscountedCashFlowScore int    `json:"discountedCashFlowScore"`
	ReturnOnEquityScore     int    `json:"returnOnEquityScore"`
	ReturnOnAssetsScore     int    `json:"returnOnAssetsScore"`
	DebtToEquityScore       int    `json:"debtToEquityScore"`
	PriceToEarningsScore    int    `json:"priceToEarningsScore"`
	PriceToBookScore        int    `json:"priceToBookScore"`

// RatingsSnapshot retrieves a comprehensive snapshot of financial ratings for stock symbols
func (c *Client) RatingsSnapshot(params RatingsSnapshotParams) ([]RatingsSnapshotResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	return doRequest[[]RatingsSnapshotResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
