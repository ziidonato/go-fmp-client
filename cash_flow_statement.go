package go_fmp

import (
	"fmt"
	"time"
)

// CashFlowStatementParams represents the parameters for the Cash Flow Statement API
type CashFlowStatementParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// CashFlowStatementResponse represents the response from the Cash Flow Statement API
type CashFlowStatementResponse struct {
	Date                                   time.Time `json:"date"`
	Symbol                                 string `json:"symbol"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	CIK                                    string `json:"cik"`
	FilingDate                             time.Time `json:"filingDate"`
	AcceptedDate                           time.Time `json:"acceptedDate"`
	CalendarYear                           string `json:"calendarYear"`
	Period                                 string `json:"period"`
	NetIncome                              float64 `json:"netIncome"`
	DepreciationAndAmortization            float64 `json:"depreciationAndAmortization"`
	DeferredIncomeTax                      float64 `json:"deferredIncomeTax"`
	StockBasedCompensation                 float64 `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                 float64 `json:"changeInWorkingCapital"`
	AccountsReceivables                    float64 `json:"accountsReceivables"`
	Inventory                              float64 `json:"inventory"`
	AccountsPayables                       float64 `json:"accountsPayables"`
	OtherWorkingCapital                    float64 `json:"otherWorkingCapital"`
	OtherNonCashItems                      float64 `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities   float64 `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment float64 `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                        float64 `json:"acquisitionsNet"`
	PurchasesOfInvestments                 float64 `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments           float64 `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivites                float64 `json:"otherInvestingActivites"`
	NetCashUsedForInvestingActivites       float64 `json:"netCashUsedForInvestingActivites"`
	DebtRepayment                          float64 `json:"debtRepayment"`
	CommonStockIssued                      float64 `json:"commonStockIssued"`
	CommonStockRepurchased                 float64 `json:"commonStockRepurchased"`
	DividendsPaid                          float64 `json:"dividendsPaid"`
	OtherFinancingActivites                float64 `json:"otherFinancingActivites"`
	NetCashUsedProvidedByFinancingActivities float64 `json:"netCashUsedProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash             float64 `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                        float64 `json:"netChangeInCash"`
	CashAtEndOfPeriod                      float64 `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                float64 `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                      float64 `json:"operatingCashFlow"`
	CapitalExpenditure                     float64 `json:"capitalExpenditure"`
	FreeCashFlow                           float64 `json:"freeCashFlow"`
	Link                                   string `json:"link"`
	FinalLink                              string `json:"finalLink"`
}

// CashFlowStatement retrieves cash flow statement data for a specific stock symbol
func (c *Client) CashFlowStatement(params CashFlowStatementParams) ([]CashFlowStatementResponse, error) {
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

	var result []CashFlowStatementResponse
	err := c.doRequest(c.BaseURL+"/cash-flow-statement", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
