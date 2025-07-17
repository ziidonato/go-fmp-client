package go_fmp

import (
	"fmt"
	"time"
)

// BulkCashFlowStatementParams represents the parameters for the Bulk Cash Flow Statement API
type BulkCashFlowStatementParams struct {
	Year   int    `json:"year"`   // Required: year (e.g., 2023)
	Period string `json:"period"` // Required: period ("annual" or "quarterly")
}

// BulkCashFlowStatementResponse represents the response from the Bulk Cash Flow Statement API
type BulkCashFlowStatementResponse struct {
	Symbol                                       string    `json:"symbol"`
	Date                                         time.Time `json:"date"`
	ReportedCurrency                             string    `json:"reportedCurrency"`
	CIK                                          string    `json:"cik"`
	FilingDate                                   time.Time `json:"filingDate"`
	AcceptedDate                                 time.Time `json:"acceptedDate"`
	CalendarYear                                 string    `json:"calendarYear"`
	Period                                       string    `json:"period"`
	NetIncome                                    float64   `json:"netIncome"`
	DepreciationAndAmortization                  float64   `json:"depreciationAndAmortization"`
	DeferredIncomeTax                            float64   `json:"deferredIncomeTax"`
	StockBasedCompensation                       float64   `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                       float64   `json:"changeInWorkingCapital"`
	AccountsReceivables                          float64   `json:"accountsReceivables"`
	Inventory                                    float64   `json:"inventory"`
	AccountsPayables                             float64   `json:"accountsPayables"`
	OtherWorkingCapital                          float64   `json:"otherWorkingCapital"`
	OtherNonCashItems                            float64   `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities         float64   `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment       float64   `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                              float64   `json:"acquisitionsNet"`
	PurchasesOfInvestments                       float64   `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments                 float64   `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivites                      float64   `json:"otherInvestingActivites"`
	NetCashUsedForInvestingActivites             float64   `json:"netCashUsedForInvestingActivites"`
	DebtRepayment                                float64   `json:"debtRepayment"`
	CommonStockIssued                            float64   `json:"commonStockIssued"`
	CommonStockRepurchased                       float64   `json:"commonStockRepurchased"`
	DividendsPaid                                float64   `json:"dividendsPaid"`
	OtherFinancingActivites                      float64   `json:"otherFinancingActivites"`
	NetCashUsedProvidedByFinancingActivities     float64   `json:"netCashUsedProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash                   float64   `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                              float64   `json:"netChangeInCash"`
	CashAtEndOfPeriod                            float64   `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                      float64   `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                            float64   `json:"operatingCashFlow"`
	CapitalExpenditure                           float64   `json:"capitalExpenditure"`
	FreeCashFlow                                 float64   `json:"freeCashFlow"`
	Link                                         string    `json:"link"`
	FinalLink                                    string    `json:"finalLink"`
}

// BulkCashFlowStatement retrieves cash flow statement data for multiple companies
func (c *Client) BulkCashFlowStatement(params BulkCashFlowStatementParams) ([]BulkCashFlowStatementResponse, error) {
	urlParams := map[string]string{
		"year":   fmt.Sprintf("%d", params.Year),
		"period": params.Period,
	}

	var result []BulkCashFlowStatementResponse
	err := c.doRequest(c.BaseURL+"/bulk-cash-flow-statement", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk cash flow statement: %w", err)
	}

	return result, nil
}
