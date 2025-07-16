package go_fmp

import (
	"encoding/json"
	"fmt"
)

// IncomeStatementAsReportedParams represents the parameters for the Income Statement As Reported API
type IncomeStatementAsReportedParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
	Period string `json:"period"` // Optional: Period type - "annual,quarter"
}

// IncomeStatementAsReportedResponse represents the response from the Income Statement As Reported API
type IncomeStatementAsReportedResponse struct {
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	CIK                                     string  `json:"cik"`
	FilingDate                              string  `json:"filingDate"`
	AcceptedDate                            string  `json:"acceptedDate"`
	FiscalYear                              string  `json:"fiscalYear"`
	Period                                  string  `json:"period"`
	Revenue                                 int64   `json:"revenue"`
	CostOfRevenue                           int64   `json:"costOfRevenue"`
	GrossProfit                             int64   `json:"grossProfit"`
	ResearchAndDevelopmentExpenses          int64   `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        int64   `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             int64   `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses int64   `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           int64   `json:"otherExpenses"`
	OperatingExpenses                       int64   `json:"operatingExpenses"`
	CostAndExpenses                         int64   `json:"costAndExpenses"`
	NetInterestIncome                       int64   `json:"netInterestIncome"`
	InterestIncome                          int64   `json:"interestIncome"`
	InterestExpense                         int64   `json:"interestExpense"`
	DepreciationAndAmortization             int64   `json:"depreciationAndAmortization"`
	EBITDA                                  int64   `json:"ebitda"`
	EBIT                                    int64   `json:"ebit"`
	NonOperatingIncomeExcludingInterest     int64   `json:"nonOperatingIncomeExcludingInterest"`
	OperatingIncome                         int64   `json:"operatingIncome"`
	TotalOtherIncomeExpensesNet             int64   `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         int64   `json:"incomeBeforeTax"`
	IncomeTaxExpense                        int64   `json:"incomeTaxExpense"`
	NetIncomeFromContinuingOperations       int64   `json:"netIncomeFromContinuingOperations"`
	NetIncomeFromDiscontinuedOperations     int64   `json:"netIncomeFromDiscontinuedOperations"`
	OtherAdjustmentsToNetIncome             int64   `json:"otherAdjustmentsToNetIncome"`
	NetIncome                               int64   `json:"netIncome"`
	NetIncomeDeductions                     int64   `json:"netIncomeDeductions"`
	BottomLineNetIncome                     int64   `json:"bottomLineNetIncome"`
	EPS                                     float64 `json:"eps"`
	EPSDiluted                              float64 `json:"epsDiluted"`
	WeightedAverageShsOut                   int64   `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                int64   `json:"weightedAverageShsOutDil"`
}

// IncomeStatementAsReported retrieves income statement as reported data for a specific stock symbol
func (c *Client) IncomeStatementAsReported(params IncomeStatementAsReportedParams) ([]IncomeStatementAsReportedResponse, error) {
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

	if params.Period != "" {
		urlParams["period"] = params.Period
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/income-statement-as-reported", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []IncomeStatementAsReportedResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
