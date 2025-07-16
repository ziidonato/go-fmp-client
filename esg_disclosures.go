package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ESGDisclosuresParams represents the parameters for the ESG Investment Search API
type ESGDisclosuresParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// ESGDisclosuresResponse represents the response from the ESG Investment Search API
type ESGDisclosuresResponse struct {
	Date               string  `json:"date"`
	AcceptedDate       string  `json:"acceptedDate"`
	Symbol             string  `json:"symbol"`
	CIK                string  `json:"cik"`
	CompanyName        string  `json:"companyName"`
	FormType           string  `json:"formType"`
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

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/esg-disclosures", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ESGDisclosuresResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
