package go_fmp

import (
	"encoding/json"
	"fmt"
)

// InstitutionalOwnershipExtractAnalyticsHolderParams represents the parameters for the Institutional Ownership Extract Analytics By Holder API
type InstitutionalOwnershipExtractAnalyticsHolderParams struct {
	Symbol  string `json:"symbol"`  // Required: Stock symbol (e.g., "AAPL")
	Year    string `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string `json:"quarter"` // Required: Quarter (e.g., "3")
	Page    *int   `json:"page"`    // Optional: Page number (default: 0)
	Limit   *int   `json:"limit"`   // Optional: Number of results (default: 10)
}

// InstitutionalOwnershipExtractAnalyticsHolderResponse represents the response from the Institutional Ownership Extract Analytics By Holder API
type InstitutionalOwnershipExtractAnalyticsHolderResponse struct {
	Date                           string  `json:"date"`
	CIK                            string  `json:"cik"`
	FilingDate                     string  `json:"filingDate"`
	InvestorName                   string  `json:"investorName"`
	Symbol                         string  `json:"symbol"`
	SecurityName                   string  `json:"securityName"`
	TypeOfSecurity                 string  `json:"typeOfSecurity"`
	SecurityCusip                  string  `json:"securityCusip"`
	SharesType                     string  `json:"sharesType"`
	PutCallShare                   string  `json:"putCallShare"`
	InvestmentDiscretion           string  `json:"investmentDiscretion"`
	IndustryTitle                  string  `json:"industryTitle"`
	Weight                         float64 `json:"weight"`
	LastWeight                     float64 `json:"lastWeight"`
	ChangeInWeight                 float64 `json:"changeInWeight"`
	ChangeInWeightPercentage       float64 `json:"changeInWeightPercentage"`
	MarketValue                    int64   `json:"marketValue"`
	LastMarketValue                int64   `json:"lastMarketValue"`
	ChangeInMarketValue            int64   `json:"changeInMarketValue"`
	ChangeInMarketValuePercentage  float64 `json:"changeInMarketValuePercentage"`
	SharesNumber                   int64   `json:"sharesNumber"`
	LastSharesNumber               int64   `json:"lastSharesNumber"`
	ChangeInSharesNumber           int64   `json:"changeInSharesNumber"`
	ChangeInSharesNumberPercentage float64 `json:"changeInSharesNumberPercentage"`
	QuarterEndPrice                float64 `json:"quarterEndPrice"`
	AvgPricePaid                   float64 `json:"avgPricePaid"`
	IsNew                          bool    `json:"isNew"`
	IsSoldOut                      bool    `json:"isSoldOut"`
	Ownership                      float64 `json:"ownership"`
	LastOwnership                  float64 `json:"lastOwnership"`
	ChangeInOwnership              float64 `json:"changeInOwnership"`
	ChangeInOwnershipPercentage    float64 `json:"changeInOwnershipPercentage"`
	HoldingPeriod                  int     `json:"holdingPeriod"`
	FirstAdded                     string  `json:"firstAdded"`
	Performance                    int64   `json:"performance"`
	PerformancePercentage          float64 `json:"performancePercentage"`
	LastPerformance                int64   `json:"lastPerformance"`
	ChangeInPerformance            int64   `json:"changeInPerformance"`
	IsCountedForPerformance        bool    `json:"isCountedForPerformance"`
}

// InstitutionalOwnershipExtractAnalyticsHolder retrieves analytical breakdown of institutional filings by holder
func (c *Client) InstitutionalOwnershipExtractAnalyticsHolder(params InstitutionalOwnershipExtractAnalyticsHolderParams) ([]InstitutionalOwnershipExtractAnalyticsHolderResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if params.Quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"symbol":  params.Symbol,
		"year":    params.Year,
		"quarter": params.Quarter,
	}

	if params.Page != nil {
		urlParams["page"] = fmt.Sprintf("%d", *params.Page)
	}

	if params.Limit != nil {
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/institutional-ownership/extract-analytics/holder", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []InstitutionalOwnershipExtractAnalyticsHolderResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
