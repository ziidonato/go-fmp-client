package go_fmp

import (
	"fmt"
)

// ESGRatingsParams represents the parameters for the ESG Ratings API
type ESGRatingsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// ESGRatingsResponse represents the response from the ESG Ratings API
type ESGRatingsResponse struct {
	Symbol        string `json:"symbol"`
	CIK           string `json:"cik"`
	CompanyName   string `json:"companyName"`
	Industry      string `json:"industry"`
	FiscalYear    int    `json:"fiscalYear"`
	ESGRiskRating string `json:"ESGRiskRating"`
	IndustryRank  string `json:"industryRank"`
}

// ESGRatings retrieves comprehensive ESG ratings for companies and funds
func (c *Client) ESGRatings(params ESGRatingsParams) ([]ESGRatingsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []ESGRatingsResponse
	err := c.doRequest(c.BaseURL+"/esg-ratings", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
