package go_fmp

import (
	"encoding/json"
	"fmt"
)

// FundDisclosureHoldersSearchParams represents the parameters for the Mutual Fund & ETF Disclosure Name Search API
type FundDisclosureHoldersSearchParams struct {
	Name string `json:"name"` // Required: Fund/ETF name (e.g., "Federated Hermes Government Income Securities, Inc.")
}

// FundDisclosureHoldersSearchResponse represents the response from the Mutual Fund & ETF Disclosure Name Search API
type FundDisclosureHoldersSearchResponse struct {
	Symbol              string `json:"symbol"`
	CIK                 string `json:"cik"`
	ClassID             string `json:"classId"`
	SeriesID            string `json:"seriesId"`
	EntityName          string `json:"entityName"`
	EntityOrgType       string `json:"entityOrgType"`
	SeriesName          string `json:"seriesName"`
	ClassName           string `json:"className"`
	ReportingFileNumber string `json:"reportingFileNumber"`
	Address             string `json:"address"`
	City                string `json:"city"`
	ZipCode             string `json:"zipCode"`
	State               string `json:"state"`
}

// FundDisclosureHoldersSearch searches for mutual fund and ETF disclosures by name
func (c *Client) FundDisclosureHoldersSearch(params FundDisclosureHoldersSearchParams) ([]FundDisclosureHoldersSearchResponse, error) {
	if params.Name == "" {
		return nil, fmt.Errorf("name parameter is required")
	}

	urlParams := map[string]string{
		"name": params.Name,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/funds/disclosure-holders-search", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []FundDisclosureHoldersSearchResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
