package go_fmp

import (
	"fmt"
	"time"
)

// IncomeStatementGrowthParams represents the parameters for the Income Statement Growth API
type IncomeStatementGrowthParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return (default: 20)
}

// IncomeStatementGrowthResponse represents the response from the Income Statement Growth API
type IncomeStatementGrowthResponse struct {
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	CIK                                     string  `json:"cik"`
	FilingDate                              time.Time  `json:"filingDate"`
	AcceptedDate                            time.Time  `json:"acceptedDate"`
	CalendarYear                            string  `json:"calendarYear"`
	Period                                  string  `json:"period"`
	RevenueGrowth                           float64 `json:"revenueGrowth"`
	CostOfRevenueGrowth                     float64 `json:"costOfRevenueGrowth"`
	GrossProfitGrowth                       float64 `json:"grossProfitGrowth"`
	GrossProfitRatioGrowth                  float64 `json:"grossProfitRatioGrowth"`
	ResearchAndDevelopmentExpensesGrowth    float64 `json:"researchAndDevelopmentExpensesGrowth"`
	GeneralAndAdministrativeExpensesGrowth  float64 `json:"generalAndAdministrativeExpensesGrowth"`
	SellingAndMarketingExpensesGrowth       float64 `json:"sellingAndMarketingExpensesGrowth"`
	OtherExpensesGrowth                     float64 `json:"otherExpensesGrowth"`
	OperatingExpensesGrowth                 float64 `json:"operatingExpensesGrowth"`
	CostAndExpensesGrowth                   float64 `json:"costAndExpensesGrowth"`
	InterestExpenseGrowth                   float64 `json:"interestExpenseGrowth"`
	DepreciationAndAmortizationGrowth       float64 `json:"depreciationAndAmortizationGrowth"`
	EBITDAGrowth                            float64 `json:"ebitdaGrowth"`
	EBITDARatioGrowth                       float64 `json:"ebitdaratioGrowth"`
	OperatingIncomeGrowth                   float64 `json:"operatingIncomeGrowth"`
	OperatingIncomeRatioGrowth              float64 `json:"operatingIncomeRatioGrowth"`
	TotalOtherIncomeExpensesNetGrowth       float64 `json:"totalOtherIncomeExpensesNetGrowth"`
	IncomeBeforeTaxGrowth                   float64 `json:"incomeBeforeTaxGrowth"`
	IncomeBeforeTaxRatioGrowth              float64 `json:"incomeBeforeTaxRatioGrowth"`
	IncomeTaxExpenseGrowth                  float64 `json:"incomeTaxExpenseGrowth"`
	NetIncomeGrowth                         float64 `json:"netIncomeGrowth"`
	NetIncomeRatioGrowth                    float64 `json:"netIncomeRatioGrowth"`
	EPSGrowth                               float64 `json:"epsGrowth"`
	EPSDilutedGrowth                        float64 `json:"epsdilutedGrowth"`
	WeightedAverageShsOutGrowth             float64 `json:"weightedAverageShsOutGrowth"`
	WeightedAverageShsOutDilGrowth          float64 `json:"weightedAverageShsOutDilGrowth"`
	Link                                      string  `json:"link"`
	FinalLink                                 string  `json:"finalLink"`
}

// IncomeStatementGrowth retrieves income statement growth data for a specific stock symbol
func (c *Client) IncomeStatementGrowth(params IncomeStatementGrowthParams) ([]IncomeStatementGrowthResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		if *params.Limit > 1000 {
			return nil, fmt.Errorf("limit cannot exceed 1000")
		}
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	if params.Period != nil {
		urlParams["period"] = *params.Period
	}

	var result []IncomeStatementGrowthResponse
	err := c.doRequest(c.BaseURL+"/income-statement-growth", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
