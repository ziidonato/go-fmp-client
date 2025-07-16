package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CashFlowStatementAsReportedParams represents the parameters for the Cash Flow Statement As Reported API
type CashFlowStatementAsReportedParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
	Period string `json:"period"` // Optional: Period type - "annual,quarter"
}

// CashFlowStatementAsReportedResponse represents the response from the Cash Flow Statement As Reported API
type CashFlowStatementAsReportedResponse struct {
	Date                                   string `json:"date"`
	Symbol                                 string `json:"symbol"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	CIK                                    string `json:"cik"`
	FilingDate                             string `json:"filingDate"`
	AcceptedDate                           string `json:"acceptedDate"`
	FiscalYear                             string `json:"fiscalYear"`
	Period                                 string `json:"period"`
	NetIncome                              int64  `json:"netIncome"`
	DepreciationAndAmortization            int64  `json:"depreciationAndAmortization"`
	DeferredIncomeTax                      int64  `json:"deferredIncomeTax"`
	StockBasedCompensation                 int64  `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                 int64  `json:"changeInWorkingCapital"`
	AccountsReceivables                    int64  `json:"accountsReceivables"`
	Inventory                              int64  `json:"inventory"`
	AccountsPayables                       int64  `json:"accountsPayables"`
	OtherWorkingCapital                    int64  `json:"otherWorkingCapital"`
	OtherNonCashItems                      int64  `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities   int64  `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment int64  `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                        int64  `json:"acquisitionsNet"`
	PurchasesOfInvestments                 int64  `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments           int64  `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivities               int64  `json:"otherInvestingActivities"`
	NetCashProvidedByInvestingActivities   int64  `json:"netCashProvidedByInvestingActivities"`
	NetDebtIssuance                        int64  `json:"netDebtIssuance"`
	LongTermNetDebtIssuance                int64  `json:"longTermNetDebtIssuance"`
	ShortTermNetDebtIssuance               int64  `json:"shortTermNetDebtIssuance"`
	NetStockIssuance                       int64  `json:"netStockIssuance"`
	NetCommonStockIssuance                 int64  `json:"netCommonStockIssuance"`
	CommonStockIssuance                    int64  `json:"commonStockIssuance"`
	CommonStockRepurchased                 int64  `json:"commonStockRepurchased"`
	NetPreferredStockIssuance              int64  `json:"netPreferredStockIssuance"`
	NetDividendsPaid                       int64  `json:"netDividendsPaid"`
	CommonDividendsPaid                    int64  `json:"commonDividendsPaid"`
	PreferredDividendsPaid                 int64  `json:"preferredDividendsPaid"`
	OtherFinancingActivities               int64  `json:"otherFinancingActivities"`
	NetCashProvidedByFinancingActivities   int64  `json:"netCashProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash             int64  `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                        int64  `json:"netChangeInCash"`
	CashAtEndOfPeriod                      int64  `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                int64  `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                      int64  `json:"operatingCashFlow"`
	CapitalExpenditure                     int64  `json:"capitalExpenditure"`
	FreeCashFlow                           int64  `json:"freeCashFlow"`
	IncomeTaxesPaid                        int64  `json:"incomeTaxesPaid"`
	InterestPaid                           int64  `json:"interestPaid"`
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

	if params.Period != "" {
		urlParams["period"] = params.Period
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/cash-flow-statement-as-reported", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CashFlowStatementAsReportedResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
