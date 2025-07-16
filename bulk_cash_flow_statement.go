package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CashFlowStatementBulkParams represents the parameters for the Cash Flow Statement Bulk API
type CashFlowStatementBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period string `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)

// CashFlowStatementBulkResponse represents the response from the Cash Flow Statement Bulk API
type CashFlowStatementBulkResponse struct {
	Date                                   string `json:"date"`
	Symbol                                 string `json:"symbol"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	CIK                                    string `json:"cik"`
	FilingDate                             string `json:"filingDate"`
	AcceptedDate                           string `json:"acceptedDate"`
	FiscalYear                             string `json:"fiscalYear"`
	Period                                 string `json:"period"`
	NetIncome                              string `json:"netIncome"`
	DepreciationAndAmortization            string `json:"depreciationAndAmortization"`
	DeferredIncomeTax                      string `json:"deferredIncomeTax"`
	StockBasedCompensation                 string `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                 string `json:"changeInWorkingCapital"`
	AccountsReceivables                    string `json:"accountsReceivables"`
	Inventory                              string `json:"inventory"`
	AccountsPayables                       string `json:"accountsPayables"`
	OtherWorkingCapital                    string `json:"otherWorkingCapital"`
	OtherNonCashItems                      string `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities   string `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment string `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                        string `json:"acquisitionsNet"`
	PurchasesOfInvestments                 string `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments           string `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivities               string `json:"otherInvestingActivities"`
	NetCashProvidedByInvestingActivities   string `json:"netCashProvidedByInvestingActivities"`
	NetDebtIssuance                        string `json:"netDebtIssuance"`
	LongTermNetDebtIssuance                string `json:"longTermNetDebtIssuance"`
	ShortTermNetDebtIssuance               string `json:"shortTermNetDebtIssuance"`
	NetStockIssuance                       string `json:"netStockIssuance"`
	NetCommonStockIssuance                 string `json:"netCommonStockIssuance"`
	CommonStockIssuance                    string `json:"commonStockIssuance"`
	CommonStockRepurchased                 string `json:"commonStockRepurchased"`
	NetPreferredStockIssuance              string `json:"netPreferredStockIssuance"`
	NetDividendsPaid                       string `json:"netDividendsPaid"`
	CommonDividendsPaid                    string `json:"commonDividendsPaid"`
	PreferredDividendsPaid                 string `json:"preferredDividendsPaid"`
	OtherFinancingActivities               string `json:"otherFinancingActivities"`
	NetCashProvidedByFinancingActivities   string `json:"netCashProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash             string `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                        string `json:"netChangeInCash"`
	CashAtEndOfPeriod                      string `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                string `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                      string `json:"operatingCashFlow"`
	CapitalExpenditure                     string `json:"capitalExpenditure"`
	FreeCashFlow                           string `json:"freeCashFlow"`
	IncomeTaxesPaid                        string `json:"incomeTaxesPaid"`
	InterestPaid                           string `json:"interestPaid"`

// GetCashFlowStatementBulk retrieves detailed cash flow reports for a wide range of companies
func (c *Client) GetCashFlowStatementBulk(params CashFlowStatementBulkParams) ([]CashFlowStatementBulkResponse, error) {
	// Validate required parameters
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}
	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	urlParams := map[string]string{
		"year":   params.Year,
		"period": params.Period,
	}

	return doRequest[[]CashFlowStatementBulkResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
