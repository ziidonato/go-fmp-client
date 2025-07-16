package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BalanceSheetStatementParams represents the parameters for the Balance Sheet Statement API
type BalanceSheetStatementParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
	Period string `json:"period"` // Optional: Period type - "Q1,Q2,Q3,Q4,FY,annual,quarter"
}

// BalanceSheetStatementResponse represents the response from the Balance Sheet Statement API
type BalanceSheetStatementResponse struct {
	Date                                    string `json:"date"`
	Symbol                                  string `json:"symbol"`
	ReportedCurrency                        string `json:"reportedCurrency"`
	CIK                                     string `json:"cik"`
	FilingDate                              string `json:"filingDate"`
	AcceptedDate                            string `json:"acceptedDate"`
	FiscalYear                              string `json:"fiscalYear"`
	Period                                  string `json:"period"`
	CashAndCashEquivalents                  int64  `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    int64  `json:"shortTermInvestments"`
	CashAndShortTermInvestments             int64  `json:"cashAndShortTermInvestments"`
	NetReceivables                          int64  `json:"netReceivables"`
	AccountsReceivables                     int64  `json:"accountsReceivables"`
	OtherReceivables                        int64  `json:"otherReceivables"`
	Inventory                               int64  `json:"inventory"`
	Prepaids                                int64  `json:"prepaids"`
	OtherCurrentAssets                      int64  `json:"otherCurrentAssets"`
	TotalCurrentAssets                      int64  `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               int64  `json:"propertyPlantEquipmentNet"`
	Goodwill                                int64  `json:"goodwill"`
	IntangibleAssets                        int64  `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             int64  `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     int64  `json:"longTermInvestments"`
	TaxAssets                               int64  `json:"taxAssets"`
	OtherNonCurrentAssets                   int64  `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   int64  `json:"totalNonCurrentAssets"`
	OtherAssets                             int64  `json:"otherAssets"`
	TotalAssets                             int64  `json:"totalAssets"`
	TotalPayables                           int64  `json:"totalPayables"`
	AccountPayables                         int64  `json:"accountPayables"`
	OtherPayables                           int64  `json:"otherPayables"`
	AccruedExpenses                         int64  `json:"accruedExpenses"`
	ShortTermDebt                           int64  `json:"shortTermDebt"`
	CapitalLeaseObligationsCurrent          int64  `json:"capitalLeaseObligationsCurrent"`
	TaxPayables                             int64  `json:"taxPayables"`
	DeferredRevenue                         int64  `json:"deferredRevenue"`
	OtherCurrentLiabilities                 int64  `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 int64  `json:"totalCurrentLiabilities"`
	LongTermDebt                            int64  `json:"longTermDebt"`
	DeferredRevenueNonCurrent               int64  `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        int64  `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              int64  `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              int64  `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        int64  `json:"otherLiabilities"`
	CapitalLeaseObligations                 int64  `json:"capitalLeaseObligations"`
	TotalLiabilities                        int64  `json:"totalLiabilities"`
	TreasuryStock                           int64  `json:"treasuryStock"`
	PreferredStock                          int64  `json:"preferredStock"`
	CommonStock                             int64  `json:"commonStock"`
	RetainedEarnings                        int64  `json:"retainedEarnings"`
	AdditionalPaidInCapital                 int64  `json:"additionalPaidInCapital"`
	AccumulatedOtherComprehensiveIncomeLoss int64  `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OtherTotalStockholdersEquity            int64  `json:"otherTotalStockholdersEquity"`
	TotalStockholdersEquity                 int64  `json:"totalStockholdersEquity"`
	TotalEquity                             int64  `json:"totalEquity"`
	MinorityInterest                        int64  `json:"minorityInterest"`
	TotalLiabilitiesAndTotalEquity          int64  `json:"totalLiabilitiesAndTotalEquity"`
	TotalInvestments                        int64  `json:"totalInvestments"`
	TotalDebt                               int64  `json:"totalDebt"`
	NetDebt                                 int64  `json:"netDebt"`
}

// BalanceSheetStatement retrieves balance sheet statement data for a specific stock symbol
func (c *Client) BalanceSheetStatement(params BalanceSheetStatementParams) ([]BalanceSheetStatementResponse, error) {
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

	if params.Period != " {
		urlParams["period"] = params.Period
	}

	return doRequest[[]BalanceSheetStatementResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
