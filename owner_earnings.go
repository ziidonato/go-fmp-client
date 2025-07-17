package go_fmp

import (
	"fmt"
	"time"
)

// OwnerEarningsParams represents the parameters for the Owner Earnings API
type OwnerEarningsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
}

// OwnerEarningsResponse represents the response from the Owner Earnings API
type OwnerEarningsResponse struct {
	Symbol                 string    `json:"symbol"`
	Date                   time.Time `json:"date"`
	AverageReceivables     float64   `json:"averageReceivables"`
	AverageInventory       float64   `json:"averageInventory"`
	AveragePayables        float64   `json:"averagePayables"`
	DaysReceivablesOutstanding float64 `json:"daysReceivablesOutstanding"`
	DaysInventoryOutstanding   float64 `json:"daysInventoryOutstanding"`
	DaysPayablesOutstanding    float64 `json:"daysPayablesOutstanding"`
	NetIncome              float64   `json:"netIncome"`
	DepreciationAndAmortization float64 `json:"depreciationAndAmortization"`
	MaintenanceCapex       float64   `json:"maintenanceCapex"`
	OwnerEarnings          float64   `json:"ownerEarnings"`
	GrowthCapex            float64   `json:"growthCapex"`
	ChangeInWorkingCapital float64   `json:"changeInWorkingCapital"`
}

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

	var result []OwnerEarningsResponse
	err := c.doRequest(c.BaseURL+"/owner-earnings", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
