package go_fmp

import (
	"fmt"
	"time"
)

// BulkIncomeStatementParams represents the parameters for the Bulk Income Statement API
type BulkIncomeStatementParams struct {
	Year   int    `json:"year"`   // Required: year (e.g., 2023)
	Period string `json:"period"` // Required: period ("annual" or "quarterly")
}

// BulkIncomeStatementResponse represents the response from the Bulk Income Statement API
type BulkIncomeStatementResponse struct {
	Symbol                                  string    `json:"symbol"`
	Date                                    time.Time `json:"date"`
	ReportedCurrency                        string    `json:"reportedCurrency"`
	CIK                                     string    `json:"cik"`
	FilingDate                              time.Time `json:"filingDate"`
	AcceptedDate                            time.Time `json:"acceptedDate"`
	CalendarYear                            string    `json:"calendarYear"`
	Period                                  string    `json:"period"`
	Revenue                                 float64   `json:"revenue"`
	CostOfRevenue                           float64   `json:"costOfRevenue"`
	GrossProfit                             float64   `json:"grossProfit"`
	GrossProfitRatio                        float64   `json:"grossProfitRatio"`
	ResearchAndDevelopmentExpenses          float64   `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        float64   `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             float64   `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses float64   `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           float64   `json:"otherExpenses"`
	OperatingExpenses                       float64   `json:"operatingExpenses"`
	CostAndExpenses                         float64   `json:"costAndExpenses"`
	InterestIncome                          float64   `json:"interestIncome"`
	InterestExpense                         float64   `json:"interestExpense"`
	DepreciationAndAmortization             float64   `json:"depreciationAndAmortization"`
	EBITDA                                  float64   `json:"ebitda"`
	EBITDARatio                             float64   `json:"ebitdaratio"`
	OperatingIncome                         float64   `json:"operatingIncome"`
	OperatingIncomeRatio                    float64   `json:"operatingIncomeRatio"`
	TotalOtherIncomeExpensesNet             float64   `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         float64   `json:"incomeBeforeTax"`
	IncomeBeforeTaxRatio                    float64   `json:"incomeBeforeTaxRatio"`
	IncomeTaxExpense                        float64   `json:"incomeTaxExpense"`
	NetIncome                               float64   `json:"netIncome"`
	NetIncomeRatio                          float64   `json:"netIncomeRatio"`
	EPS                                     float64   `json:"eps"`
	EPSDiluted                              float64   `json:"epsdiluted"`
	WeightedAverageShsOut                   float64   `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                float64   `json:"weightedAverageShsOutDil"`
	Link                                    string    `json:"link"`
	FinalLink                               string    `json:"finalLink"`
}

// BulkIncomeStatement retrieves income statement data for multiple companies
func (c *Client) BulkIncomeStatement(params BulkIncomeStatementParams) ([]BulkIncomeStatementResponse, error) {
	urlParams := map[string]string{
		"year":   fmt.Sprintf("%d", params.Year),
		"period": params.Period,
	}

	var result []BulkIncomeStatementResponse
	err := c.doRequest(c.BaseURL+"/bulk-income-statement", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk income statement: %w", err)
	}

	return result, nil
}
