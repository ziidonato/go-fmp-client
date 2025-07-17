package go_fmp

import (
	"fmt"
	"time"
)

// BalanceSheetStatementTTMParams represents the parameters for the Balance Sheet Statement TTM API
type BalanceSheetStatementTTMParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
}

// BalanceSheetStatementTTMResponse represents the response from the Balance Sheet Statement TTM API
type BalanceSheetStatementTTMResponse struct {
	Date                                    time.Time `json:"date"`
	Symbol                                  string `json:"symbol"`
	ReportedCurrency                        string `json:"reportedCurrency"`
	CIK                                     string `json:"cik"`
	FilingDate                              time.Time `json:"filingDate"`
	AcceptedDate                            time.Time `json:"acceptedDate"`
	CalendarYear                            string `json:"calendarYear"`
	Period                                  string `json:"period"`
	CashAndCashEquivalentsTTM               float64 `json:"cashAndCashEquivalentsTTM"`
	ShortTermInvestmentsTTM                 float64 `json:"shortTermInvestmentsTTM"`
	CashAndShortTermInvestmentsTTM          float64 `json:"cashAndShortTermInvestmentsTTM"`
	NetReceivablesTTM                       float64 `json:"netReceivablesTTM"`
	InventoryTTM                            float64 `json:"inventoryTTM"`
	OtherCurrentAssetsTTM                   float64 `json:"otherCurrentAssetsTTM"`
	TotalCurrentAssetsTTM                   float64 `json:"totalCurrentAssetsTTM"`
	PropertyPlantEquipmentNetTTM            float64 `json:"propertyPlantEquipmentNetTTM"`
	GoodwillTTM                             float64 `json:"goodwillTTM"`
	IntangibleAssetsTTM                     float64 `json:"intangibleAssetsTTM"`
	GoodwillAndIntangibleAssetsTTM          float64 `json:"goodwillAndIntangibleAssetsTTM"`
	LongTermInvestmentsTTM                  float64 `json:"longTermInvestmentsTTM"`
	TaxAssetsTTM                            float64 `json:"taxAssetsTTM"`
	OtherNonCurrentAssetsTTM                float64 `json:"otherNonCurrentAssetsTTM"`
	TotalNonCurrentAssetsTTM                float64 `json:"totalNonCurrentAssetsTTM"`
	OtherAssetsTTM                          float64 `json:"otherAssetsTTM"`
	TotalAssetsTTM                          float64 `json:"totalAssetsTTM"`
	AccountPayablesTTM                      float64 `json:"accountPayablesTTM"`
	ShortTermDebtTTM                        float64 `json:"shortTermDebtTTM"`
	TaxPayablesTTM                          float64 `json:"taxPayablesTTM"`
	DeferredRevenueTTM                      float64 `json:"deferredRevenueTTM"`
	OtherCurrentLiabilitiesTTM              float64 `json:"otherCurrentLiabilitiesTTM"`
	TotalCurrentLiabilitiesTTM              float64 `json:"totalCurrentLiabilitiesTTM"`
	LongTermDebtTTM                         float64 `json:"longTermDebtTTM"`
	DeferredRevenueNonCurrentTTM            float64 `json:"deferredRevenueNonCurrentTTM"`
	DeferredTaxLiabilitiesNonCurrentTTM     float64 `json:"deferredTaxLiabilitiesNonCurrentTTM"`
	OtherNonCurrentLiabilitiesTTM           float64 `json:"otherNonCurrentLiabilitiesTTM"`
	TotalNonCurrentLiabilitiesTTM           float64 `json:"totalNonCurrentLiabilitiesTTM"`
	OtherLiabilitiesTTM                     float64 `json:"otherLiabilitiesTTM"`
	CapitalLeaseObligationsTTM              float64 `json:"capitalLeaseObligationsTTM"`
	TotalLiabilitiesTTM                     float64 `json:"totalLiabilitiesTTM"`
	PreferredStockTTM                       float64 `json:"preferredStockTTM"`
	CommonStockTTM                          float64 `json:"commonStockTTM"`
	RetainedEarningsTTM                     float64 `json:"retainedEarningsTTM"`
	AccumulatedOtherComprehensiveIncomeLossTTM float64 `json:"accumulatedOtherComprehensiveIncomeLossTTM"`
	OtherTotalStockholdersEquityTTM         float64 `json:"otherTotalStockholdersEquityTTM"`
	TotalStockholdersEquityTTM              float64 `json:"totalStockholdersEquityTTM"`
	TotalLiabilitiesAndStockholdersEquityTTM float64 `json:"totalLiabilitiesAndStockholdersEquityTTM"`
	MinorityInterestTTM                     float64 `json:"minorityInterestTTM"`
	TotalEquityTTM                          float64 `json:"totalEquityTTM"`
	TotalLiabilitiesAndTotalEquityTTM       float64 `json:"totalLiabilitiesAndTotalEquityTTM"`
	TotalInvestmentsTTM                     float64 `json:"totalInvestmentsTTM"`
	TotalDebtTTM                            float64 `json:"totalDebtTTM"`
	NetDebtTTM                              float64 `json:"netDebtTTM"`
	Link                                    string `json:"link"`
	FinalLink                               string `json:"finalLink"`
}

// BalanceSheetStatementTTM retrieves trailing twelve months balance sheet statement data for a specific stock symbol
func (c *Client) BalanceSheetStatementTTM(params BalanceSheetStatementTTMParams) ([]BalanceSheetStatementTTMResponse, error) {
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

	var result []BalanceSheetStatementTTMResponse
	err := c.doRequest(c.BaseURL+"/balance-sheet-statement-ttm", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
