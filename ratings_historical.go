package go_fmp

import (
	"fmt"
)

// RatingsHistoricalParams represents the parameters for the Historical Ratings API
type RatingsHistoricalParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 10000 records per request)
}

// RatingsHistoricalResponse represents the response from the Historical Ratings API
type RatingsHistoricalResponse struct {
	Symbol                  string `json:"symbol"`
	Date                    string `json:"date"`
	Rating                  string `json:"rating"`
	OverallScore            int    `json:"overallScore"`
	DiscountedCashFlowScore int    `json:"discountedCashFlowScore"`
	ReturnOnEquityScore     int    `json:"returnOnEquityScore"`
	ReturnOnAssetsScore     int    `json:"returnOnAssetsScore"`
	DebtToEquityScore       int    `json:"debtToEquityScore"`
	PriceToEarningsScore    int    `json:"priceToEarningsScore"`
	PriceToBookScore        int    `json:"priceToBookScore"`
}

// RatingsHistorical retrieves historical financial ratings for stock symbols
func (c *Client) RatingsHistorical(params RatingsHistoricalParams) ([]RatingsHistoricalResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		if *params.Limit > 10000 {
			return nil, fmt.Errorf("limit cannot exceed 10000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	var result []RatingsHistoricalResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/ratings-historical", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
