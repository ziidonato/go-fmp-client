package go_fmp

import (
	"encoding/json"
	"fmt"
)

// RevenueGeographicSegmentationParams represents the parameters for the Revenue Geographic Segmentation API
type RevenueGeographicSegmentationParams struct {
	Symbol    string `json:"symbol"`    // Required: Stock symbol (e.g., "AAPL")
	Period    string `json:"period"`    // Optional: Period type - "annual,quarter"
	Structure string `json:"structure"` // Optional: Structure type - "flat"
}

// RevenueGeographicSegmentationResponse represents the response from the Revenue Geographic Segmentation API
type RevenueGeographicSegmentationResponse struct {
	Symbol           string                 `json:"symbol"`
	FiscalYear       int                    `json:"fiscalYear"`
	Period           string                 `json:"period"`
	ReportedCurrency *string                `json:"reportedCurrency"`
	Date             string                 `json:"date"`
	Data             map[string]interface{} `json:"data"`
}

// RevenueGeographicSegmentation retrieves revenue geographic segmentation data for a specific stock symbol
func (c *Client) RevenueGeographicSegmentation(params RevenueGeographicSegmentationParams) ([]RevenueGeographicSegmentationResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Period != "" {
		urlParams["period"] = params.Period
	}

	if params.Structure != "" {
		urlParams["structure"] = params.Structure
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/revenue-geographic-segmentation", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []RevenueGeographicSegmentationResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
