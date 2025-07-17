package go_fmp

import (
	"fmt"
	"time"
)

// IncomeStatementParams represents the parameters for the Income Statement API
type IncomeStatementParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *Period `json:"period"` // Optional: Period type (PeriodAnnual or PeriodQuarter)
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// IncomeStatementResponse represents the response from the Income Statement API
type IncomeStatementResponse struct {
	Date                                    time.Time  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	CIK                                     string  `json:"cik"`
	FilingDate                              time.Time  `json:"filingDate"`
	AcceptedDate                            time.Time  `json:"acceptedDate"`
	CalendarYear                            string  `json:"calendarYear"`
	Period                                  string  `json:"period"`
	Revenue                                 float64 `json:"revenue"`
	CostOfRevenue                           float64 `json:"costOfRevenue"`
	GrossProfit                             float64 `json:"grossProfit"`
	GrossProfitRatio                        float64 `json:"grossProfitRatio"`
	ResearchAndDevelopmentExpenses          float64 `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        float64 `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             float64 `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses float64 `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           float64 `json:"otherExpenses"`
	OperatingExpenses                       float64 `json:"operatingExpenses"`
	CostAndExpenses                         float64 `json:"costAndExpenses"`
	InterestIncome                          float64 `json:"interestIncome"`
	InterestExpense                         float64 `json:"interestExpense"`
	DepreciationAndAmortization             float64 `json:"depreciationAndAmortization"`
	EBITDA                                  float64 `json:"ebitda"`
	EBITDARatio                             float64 `json:"ebitdaratio"`
	OperatingIncome                         float64 `json:"operatingIncome"`
	OperatingIncomeRatio                    float64 `json:"operatingIncomeRatio"`
	TotalOtherIncomeExpensesNet             float64 `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         float64 `json:"incomeBeforeTax"`
	IncomeBeforeTaxRatio                    float64 `json:"incomeBeforeTaxRatio"`
	IncomeTaxExpense                        float64 `json:"incomeTaxExpense"`
	NetIncome                               float64 `json:"netIncome"`
	NetIncomeRatio                          float64 `json:"netIncomeRatio"`
	EPS                                     float64 `json:"eps"`
	EPSDiluted                              float64 `json:"epsdiluted"`
	WeightedAverageShsOut                   float64 `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                float64 `json:"weightedAverageShsOutDil"`
	Link                                    string  `json:"link"`
	FinalLink                               string  `json:"finalLink"`
}

// IncomeStatement retrieves income statement data for a specific stock symbol
func (c *Client) IncomeStatement(params IncomeStatementParams) ([]IncomeStatementResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.Limit != nil {
		urlParams["limit"] = fmt.Sprintf("%d", *params.Limit)
	}

	if params.Period != nil {
		urlParams["period"] = string(*params.Period)
	}

	var result []IncomeStatementResponse
	err := c.doRequest(c.BaseURL+"/income-statement", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
