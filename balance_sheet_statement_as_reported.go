package go_fmp

import (
	"fmt"
	"time"
)

// BalanceSheetStatementAsReportedParams represents the parameters for the Balance Sheet Statement As Reported API
type BalanceSheetStatementAsReportedParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// BalanceSheetStatementAsReportedResponse represents the response from the Balance Sheet Statement As Reported API
type BalanceSheetStatementAsReportedResponse struct {
	Date                                    time.Time `json:"date"`
	Symbol                                  string `json:"symbol"`
	ReportedCurrency                        string `json:"reportedCurrency"`
	CIK                                     string `json:"cik"`
	FilingDate                              time.Time `json:"filingDate"`
	AcceptedDate                            time.Time `json:"acceptedDate"`
	CalendarYear                            string `json:"calendarYear"`
	Period                                  string `json:"period"`
	CashAndCashEquivalents                  int64 `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    int64 `json:"shortTermInvestments"`
	CashAndShortTermInvestments             int64 `json:"cashAndShortTermInvestments"`
	NetReceivables                          int64 `json:"netReceivables"`
	Inventory                               int64 `json:"inventory"`
	OtherCurrentAssets                      int64 `json:"otherCurrentAssets"`
	TotalCurrentAssets                      int64 `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               int64 `json:"propertyPlantEquipmentNet"`
	Goodwill                                int64 `json:"goodwill"`
	IntangibleAssets                        int64 `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             int64 `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     int64 `json:"longTermInvestments"`
	TaxAssets                               int64 `json:"taxAssets"`
	OtherNonCurrentAssets                   int64 `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   int64 `json:"totalNonCurrentAssets"`
	OtherAssets                             int64 `json:"otherAssets"`
	TotalAssets                             int64 `json:"totalAssets"`
	AccountPayables                         int64 `json:"accountPayables"`
	ShortTermDebt                           int64 `json:"shortTermDebt"`
	TaxPayables                             int64 `json:"taxPayables"`
	DeferredRevenue                         int64 `json:"deferredRevenue"`
	OtherCurrentLiabilities                 int64 `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 int64 `json:"totalCurrentLiabilities"`
	LongTermDebt                            int64 `json:"longTermDebt"`
	DeferredRevenueNonCurrent               int64 `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        int64 `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              int64 `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              int64 `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        int64 `json:"otherLiabilities"`
	CapitalLeaseObligations                 int64 `json:"capitalLeaseObligations"`
	TotalLiabilities                        int64 `json:"totalLiabilities"`
	PreferredStock                          int64 `json:"preferredStock"`
	CommonStock                             int64 `json:"commonStock"`
	RetainedEarnings                        int64 `json:"retainedEarnings"`
	AccumulatedOtherComprehensiveIncomeLoss int64 `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OtherTotalStockholdersEquity            int64 `json:"otherTotalStockholdersEquity"`
	TotalStockholdersEquity                 int64 `json:"totalStockholdersEquity"`
	TotalLiabilitiesAndStockholdersEquity   int64 `json:"totalLiabilitiesAndStockholdersEquity"`
	MinorityInterest                        int64 `json:"minorityInterest"`
	TotalEquity                             int64 `json:"totalEquity"`
	TotalLiabilitiesAndTotalEquity          int64 `json:"totalLiabilitiesAndTotalEquity"`
	TotalInvestments                        int64 `json:"totalInvestments"`
	TotalDebt                               int64 `json:"totalDebt"`
	NetDebt                                 int64 `json:"netDebt"`
	Link                                    string `json:"link"`
	FinalLink                               string `json:"finalLink"`
}

// BalanceSheetStatementAsReported retrieves balance sheet statement as reported data for a specific stock symbol
func (c *Client) BalanceSheetStatementAsReported(params BalanceSheetStatementAsReportedParams) ([]BalanceSheetStatementAsReportedResponse, error) {
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

	var result []BalanceSheetStatementAsReportedResponse
	err := c.doRequest(c.BaseURL+"/balance-sheet-statement-as-reported", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
