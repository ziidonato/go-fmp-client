package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CrowdfundingOfferingsSearchParams represents the parameters for the Crowdfunding Campaign Search API
type CrowdfundingOfferingsSearchParams struct {
	Name string `json:"name"` // Required: Company name, campaign name, or platform (e.g., "enotap")
}

// CrowdfundingOfferingsSearchResponse represents the response from the Crowdfunding Campaign Search API
type CrowdfundingOfferingsSearchResponse struct {
	CIK  string  `json:"cik"`
	Name string  `json:"name"`
	Date *string `json:"date"`
}

// CrowdfundingOfferingsSearch searches for crowdfunding campaigns by company name, campaign name, or platform
func (c *Client) CrowdfundingOfferingsSearch(params CrowdfundingOfferingsSearchParams) ([]CrowdfundingOfferingsSearchResponse, error) {
	if params.Name == "" {
		return nil, fmt.Errorf("name parameter is required")
	}

	urlParams := map[string]string{
		"name": params.Name,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/crowdfunding-offerings-search", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CrowdfundingOfferingsSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
