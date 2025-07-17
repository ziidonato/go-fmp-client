package go_fmp

import (
	"time"
	"fmt"
)

// ESGDisclosuresParams represents the parameters for the ESG Investment Search API
type ESGDisclosuresParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// ESGDisclosuresResponse represents the response from the ESG Investment Search API
type ESGDisclosuresResponse struct {
	Date time.Time `json:"date"`
	AcceptedDate time.Time `json:"acceptedDate"`
	Symbol             string  `json:"symbol"`
	CIK                string  `json:"cik"`
	CompanyName        string  `json:"companyName"`
	FormType FormType `json:"formType"`
	EnvironmentalScore float64 `json:"environmentalScore"`
	SocialScore        float64 `json:"socialScore"`
	GovernanceScore    float64 `json:"governanceScore"`
	ESGScore           float64 `json:"ESGScore"`
	URL                string  `json:"url"`
}

// ESGDisclosures retrieves ESG disclosures for companies and funds
func (c *Client) ESGDisclosures(params ESGDisclosuresParams) ([]ESGDisclosuresResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []ESGDisclosuresResponse
	err := c.doRequest(c.BaseURL+"/esg-disclosures", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
