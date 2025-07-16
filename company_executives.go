package go_fmp

import (
	"fmt"
)

// CompanyExecutivesParams represents the parameters for the Company Executives API
type CompanyExecutivesParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Active *string `json:"active"` // Required: Filter for active executives (e.g., "true")
}

// CompanyExecutivesResponse represents the response from the Company Executives API
type CompanyExecutivesResponse struct {
	Title       string `json:"title"`
	Name        string `json:"name"`
	Pay         *int   `json:"pay"`
	CurrencyPay string `json:"currencyPay"`
	Gender      string `json:"gender"`
	YearBorn    *int   `json:"yearBorn"`
	Active      *bool  `json:"active"`
}

// CompanyExecutives retrieves detailed information on company executives
func (c *Client) CompanyExecutives(params CompanyExecutivesParams) ([]CompanyExecutivesResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Active == nil {
		return nil, fmt.Errorf("active parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
		"active": *params.Active,
	}

	var result []CompanyExecutivesResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/key-executives", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
