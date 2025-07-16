package go_fmp

import (
	"encoding/json"
	"fmt"
)

// LatestFinancialStatementsParams represents the parameters for the Latest Financial Statements API
type LatestFinancialStatementsParams struct {
	Page  *int `json:"page"`  // Optional: Page number (default: 0)
	Limit *int `json:"limit"` // Optional: Number of results (Maximum 250 records per request)
}

// LatestFinancialStatementsResponse represents the response from the Latest Financial Statements API
type LatestFinancialStatementsResponse struct {
	Symbol       string `json:"symbol"`
	CalendarYear int    `json:"calendarYear"`
	Period       string `json:"period"`
	Date         string `json:"date"`
	DateAdded    string `json:"dateAdded"`
}

// LatestFinancialStatements retrieves the latest financial statements data
func (c *Client) LatestFinancialStatements(params LatestFinancialStatementsParams) ([]LatestFinancialStatementsResponse, error) {
	urlParams := map[string]string{}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
	}

	if params.Limit != nil {
		if *params.Limit > 250 {
			return nil, fmt.Errorf("limit cannot exceed 250")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/latest-financial-statements", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []LatestFinancialStatementsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
