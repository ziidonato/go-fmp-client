package go_fmp

import (
	"encoding/json"
	"fmt"
)

// OwnerEarningsParams represents the parameters for the Owner Earnings API
type OwnerEarningsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)

// OwnerEarningsResponse represents the response from the Owner Earnings API
type OwnerEarningsResponse struct {
	Symbol                 string  `json:"symbol"`
	ReportedCurrency       string  `json:"reportedCurrency"`
	FiscalYear             string  `json:"fiscalYear"`
	Period                 string  `json:"period"`
	Date                   string  `json:"date"`
	AveragePPE             float64 `json:"averagePPE"`
	MaintenanceCapex       int64   `json:"maintenanceCapex"`
	OwnersEarnings         int64   `json:"ownersEarnings"`
	GrowthCapex            int64   `json:"growthCapex"`
	OwnersEarningsPerShare float64 `json:"ownersEarningsPerShare"`

// OwnerEarnings retrieves owner earnings data for a specific stock symbol
func (c *Client) OwnerEarnings(params OwnerEarningsParams) ([]OwnerEarningsResponse, error) {
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

	return doRequest[[]OwnerEarningsResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
