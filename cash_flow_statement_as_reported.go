package go_fmp

import (
	"fmt"
	"time"
)

// CashFlowStatementAsReportedParams represents the parameters for the Cash Flow Statement As Reported API
type CashFlowStatementAsReportedParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// CashFlowStatementAsReportedResponse represents the response from the Cash Flow Statement As Reported API
type CashFlowStatementAsReportedResponse struct {
	Date                                   time.Time `json:"date"`
	Symbol                                 string `json:"symbol"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	CIK                                    string `json:"cik"`
	FilingDate                             time.Time `json:"filingDate"`
	AcceptedDate                           time.Time `json:"acceptedDate"`
	CalendarYear                           string `json:"calendarYear"`
	Period                                 string `json:"period"`
	NetIncome                              int64 `json:"netIncome"`
	DepreciationAndAmortization            int64 `json:"depreciationAndAmortization"`
	DeferredIncomeTax                      int64 `json:"deferredIncomeTax"`
	StockBasedCompensation                 int64 `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                 int64 `json:"changeInWorkingCapital"`
	AccountsReceivables                    int64 `json:"accountsReceivables"`
	Inventory                              int64 `json:"inventory"`
	AccountsPayables                       int64 `json:"accountsPayables"`
	OtherWorkingCapital                    int64 `json:"otherWorkingCapital"`
	OtherNonCashItems                      int64 `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities   int64 `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment int64 `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                        int64 `json:"acquisitionsNet"`
	PurchasesOfInvestments                 int64 `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments           int64 `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivites                int64 `json:"otherInvestingActivites"`
	NetCashUsedForInvestingActivites       int64 `json:"netCashUsedForInvestingActivites"`
	DebtRepayment                          int64 `json:"debtRepayment"`
	CommonStockIssued                      int64 `json:"commonStockIssued"`
	CommonStockRepurchased                 int64 `json:"commonStockRepurchased"`
	DividendsPaid                          int64 `json:"dividendsPaid"`
	OtherFinancingActivites                int64 `json:"otherFinancingActivites"`
	NetCashUsedProvidedByFinancingActivities int64 `json:"netCashUsedProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash             int64 `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                        int64 `json:"netChangeInCash"`
	CashAtEndOfPeriod                      int64 `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                int64 `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                      int64 `json:"operatingCashFlow"`
	CapitalExpenditure                     int64 `json:"capitalExpenditure"`
	FreeCashFlow                           int64 `json:"freeCashFlow"`
	Link                                   string `json:"link"`
	FinalLink                              string `json:"finalLink"`
}

// CashFlowStatementAsReported retrieves cash flow statement as reported data for a specific stock symbol
func (c *Client) CashFlowStatementAsReported(params CashFlowStatementAsReportedParams) ([]CashFlowStatementAsReportedResponse, error) {
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

	var result []CashFlowStatementAsReportedResponse
	err := c.doRequest(c.BaseURL+"/cash-flow-statement-as-reported", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
