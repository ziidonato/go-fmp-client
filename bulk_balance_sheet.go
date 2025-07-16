package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BalanceSheetBulkParams represents the parameters for the Bulk Balance Sheet Statement API
type BalanceSheetBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period string `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)

// BalanceSheetBulkResponse represents the response from the Bulk Balance Sheet Statement API
type BalanceSheetBulkResponse struct {
	Date                                    string `json:"date"`
	Symbol                                  string `json:"symbol"`
	ReportedCurrency                        string `json:"reportedCurrency"`
	CIK                                     string `json:"cik"`
	FilingDate                              string `json:"filingDate"`
	AcceptedDate                            string `json:"acceptedDate"`
	FiscalYear                              string `json:"fiscalYear"`
	Period                                  string `json:"period"`
	CashAndCashEquivalents                  string `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    string `json:"shortTermInvestments"`
	CashAndShortTermInvestments             string `json:"cashAndShortTermInvestments"`
	NetReceivables                          string `json:"netReceivables"`
	AccountsReceivables                     string `json:"accountsReceivables"`
	OtherReceivables                        string `json:"otherReceivables"`
	Inventory                               string `json:"inventory"`
	Prepaids                                string `json:"prepaids"`
	OtherCurrentAssets                      string `json:"otherCurrentAssets"`
	TotalCurrentAssets                      string `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               string `json:"propertyPlantEquipmentNet"`
	Goodwill                                string `json:"goodwill"`
	IntangibleAssets                        string `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             string `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     string `json:"longTermInvestments"`
	TaxAssets                               string `json:"taxAssets"`
	OtherNonCurrentAssets                   string `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   string `json:"totalNonCurrentAssets"`
	OtherAssets                             string `json:"otherAssets"`
	TotalAssets                             string `json:"totalAssets"`
	TotalPayables                           string `json:"totalPayables"`
	AccountPayables                         string `json:"accountPayables"`
	OtherPayables                           string `json:"otherPayables"`
	AccruedExpenses                         string `json:"accruedExpenses"`
	ShortTermDebt                           string `json:"shortTermDebt"`
	CapitalLeaseObligationsCurrent          string `json:"capitalLeaseObligationsCurrent"`
	TaxPayables                             string `json:"taxPayables"`
	DeferredRevenue                         string `json:"deferredRevenue"`
	OtherCurrentLiabilities                 string `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 string `json:"totalCurrentLiabilities"`
	LongTermDebt                            string `json:"longTermDebt"`
	CapitalLeaseObligationsNonCurrent       string `json:"capitalLeaseObligationsNonCurrent"`
	DeferredRevenueNonCurrent               string `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        string `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              string `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              string `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        string `json:"otherLiabilities"`
	CapitalLeaseObligations                 string `json:"capitalLeaseObligations"`
	TotalLiabilities                        string `json:"totalLiabilities"`
	TreasuryStock                           string `json:"treasuryStock"`
	PreferredStock                          string `json:"preferredStock"`
	CommonStock                             string `json:"commonStock"`
	RetainedEarnings                        string `json:"retainedEarnings"`
	AdditionalPaidInCapital                 string `json:"additionalPaidInCapital"`
	AccumulatedOtherComprehensiveIncomeLoss string `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OtherTotalStockholdersEquity            string `json:"otherTotalStockholdersEquity"`
	TotalStockholdersEquity                 string `json:"totalStockholdersEquity"`
	TotalEquity                             string `json:"totalEquity"`
	MinorityInterest                        string `json:"minorityInterest"`
	TotalLiabilitiesAndTotalEquity          string `json:"totalLiabilitiesAndTotalEquity"`
	TotalInvestments                        string `json:"totalInvestments"`
	TotalDebt                               string `json:"totalDebt"`
	NetDebt                                 string `json:"netDebt"`

// BalanceSheetBulk retrieves comprehensive balance sheet data across multiple companies
func (c *Client) BalanceSheetBulk(params BalanceSheetBulkParams) ([]BalanceSheetBulkResponse, error) {
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

	return doRequest[[]BalanceSheetBulkResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
