package go_fmp

import (
	"encoding/json"
	"fmt"
)

// RevenueProductSegmentationParams represents the parameters for the Revenue Product Segmentation API
type RevenueProductSegmentationParams struct {
	Symbol    string `json:"symbol"`    // Required: Stock symbol (e.g., "AAPL")
	Period    string `json:"period"`    // Optional: Period type - "annual,quarter"
	Structure string `json:"structure"` // Optional: Structure type - "flat"
}

// RevenueProductSegmentationResponse represents the response from the Revenue Product Segmentation API
type RevenueProductSegmentationResponse struct {
	Symbol           string                 `json:"symbol"`
	FiscalYear       int                    `json:"fiscalYear"`
	Period           string                 `json:"period"`
	ReportedCurrency *string                `json:"reportedCurrency"`
	Date             string                 `json:"date"`
	Data             map[string]interface{} `json:"data"`
}

// RevenueProductSegmentation retrieves revenue product segmentation data for a specific stock symbol
func (c *Client) RevenueProductSegmentation(params RevenueProductSegmentationParams) ([]RevenueProductSegmentationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Period != " {
		urlParams["period"] = params.Period
	}

	if params.Structure != " {
		urlParams["structure"] = params.Structure
	}

	return doRequest[[]RevenueProductSegmentationResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
