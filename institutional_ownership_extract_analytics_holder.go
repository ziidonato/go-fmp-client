package go_fmp

import (
	"fmt"
	"time"
)

// InstitutionalOwnershipExtractAnalyticsHolderParams represents the parameters for the Institutional Ownership Extract Analytics By Holder API
type InstitutionalOwnershipExtractAnalyticsHolderParams struct {
	Symbol  string `json:"symbol"`  // Required: Stock symbol (e.g., "AAPL")
	Year    string `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string `json:"quarter"` // Required: Quarter (e.g., "3")
	Page    *int   `json:"page"`    // Optional: Page number (default: 0)
	Limit   *int   `json:"limit"`   // Optional: Number of results (default: 10)
}

// InstitutionalOwnershipExtractAnalyticsHolderResponse represents the response from the Institutional Ownership Extract Analytics Holder API
type InstitutionalOwnershipExtractAnalyticsHolderResponse struct {
	CIK                            string  `json:"cik"`
	EntityName                     string  `json:"entityName"`
	Symbol                         string  `json:"symbol"`
	Date                           time.Time  `json:"date"`
	InvestorCIK                    string  `json:"investorCik"`
	FilingDate                     time.Time  `json:"filingDate"`
	InvestorEntityName             string  `json:"investorEntityName"`
	ReportedHolding                int64   `json:"reportedHolding"`
	ReportedMarketValue            float64 `json:"reportedMarketValue"`
	Change                         int64   `json:"change"`
	ChangePercent                  float64 `json:"changePercent"`
	TypeOfSecurity                 string  `json:"typeOfSecurity"`
	SecurityCUSIP                  string  `json:"securityCusip"`
	SharesType                     SharesType  `json:"sharesType"`
	PutCallShare                   string  `json:"putCallShare"`
	InvestmentDiscretion           string  `json:"investmentDiscretion"`
	OtherManager                   string  `json:"otherManager"`
	Sole                           int64   `json:"sole"`
	Shared                         int64   `json:"shared"`
	None                           int64   `json:"none"`
	Weight                         float64 `json:"weight"`
	UpdateNotificationDate         string  `json:"updateNotificationDate"`
	IsNew                          bool    `json:"isNew"`
	IsSoldOut                      bool    `json:"isSoldOut"`
	HoldingPeriod                  int     `json:"holdingPeriod"`
	HoldingTurnover                float64 `json:"holdingTurnover"`
	LastReportedMarketValue        float64 `json:"lastReportedMarketValue"`
	LastWeight                     float64 `json:"lastWeight"`
	PercentOfPortfolio             float64 `json:"percentOfPortfolio"`
	Rank                           int     `json:"rank"`
	PercentOwnershipOfSharesOut    float64 `json:"percentOwnershipOfSharesOut"`
	QuarterlyPercentChangeOwnership float64 `json:"quarterlyPercentChangeOwnership"`
	SharesOut                      int64   `json:"sharesOut"`
	MarketCap                      float64 `json:"marketCap"`
	SharePrice                     float64 `json:"sharePrice"`
	Quarter                        string  `json:"quarter"`
	Source                         string  `json:"source"`
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

	var result []InstitutionalOwnershipExtractAnalyticsHolderResponse
	err := c.doRequest(c.BaseURL+"/institutional-ownership/extract-analytics/holder", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
