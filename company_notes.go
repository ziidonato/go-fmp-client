package go_fmp

import (
	"fmt"
)

// CompanyNotesParams represents the parameters for the Company Notes API
type CompanyNotesParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// CompanyNotesResponse represents the response from the Company Notes API
type CompanyNotesResponse struct {
	CIK      string `json:"cik"`
	Symbol   string `json:"symbol"`
	Title    string `json:"title"`
	Exchange string `json:"exchange"`
}

// CompanyNotes retrieves detailed information about company-issued notes
func (c *Client) CompanyNotes(params CompanyNotesParams) ([]CompanyNotesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []CompanyNotesResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/company-notes", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
