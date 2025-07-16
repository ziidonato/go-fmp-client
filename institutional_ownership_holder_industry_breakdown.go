package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipHolderIndustryBreakdownParams represents the parameters for the Institutional Ownership Holder Industry Breakdown API
type InstitutionalOwnershipHolderIndustryBreakdownParams struct {
	CIK     string `json:"cik"`     // Required: CIK number (e.g., "0001067983")
	Year    string `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string `json:"quarter"` // Required: Quarter (e.g., "3")
}

// InstitutionalOwnershipHolderIndustryBreakdownResponse represents the response from the Institutional Ownership Holder Industry Breakdown API
type InstitutionalOwnershipHolderIndustryBreakdownResponse struct {
	Date                     string  `json:"date"`
	CIK                      string  `json:"cik"`
	InvestorName             string  `json:"investorName"`
	IndustryTitle            string  `json:"industryTitle"`
	Weight                   float64 `json:"weight"`
	LastWeight               float64 `json:"lastWeight"`
	ChangeInWeight           float64 `json:"changeInWeight"`
	ChangeInWeightPercentage float64 `json:"changeInWeightPercentage"`
	Performance              int64   `json:"performance"`
	PerformancePercentage    float64 `json:"performancePercentage"`
	LastPerformance          int64   `json:"lastPerformance"`
	ChangeInPerformance      int64   `json:"changeInPerformance"`
}

// InstitutionalOwnershipHolderIndustryBreakdown retrieves industry breakdown for institutional holders
func (c *Client) InstitutionalOwnershipHolderIndustryBreakdown(params InstitutionalOwnershipHolderIndustryBreakdownParams) ([]InstitutionalOwnershipHolderIndustryBreakdownResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if params.Quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"cik":     params.CIK,
		"year":    params.Year,
		"quarter": params.Quarter,
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/institutional-ownership/holder-industry-breakdown", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []InstitutionalOwnershipHolderIndustryBreakdownResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
