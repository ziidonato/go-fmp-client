package go_fmp

import (
	"encoding/json"
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

	resp, err := c.get("https://financialmodelingprep.com/stable/esg-ratings", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ESGRatingsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
