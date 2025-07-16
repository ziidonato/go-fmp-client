package go_fmp

import (
	"encoding/json"
	"fmt"
)

// BalanceSheetGrowthBulkParams represents the parameters for the Balance Sheet Growth Bulk API
type BalanceSheetGrowthBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period string `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)
}

// BalanceSheetGrowthBulkResponse represents the response from the Balance Sheet Growth Bulk API
type BalanceSheetGrowthBulkResponse struct {
	Symbol                                        string `json:"symbol"`
	Date                                          string `json:"date"`
	FiscalYear                                    string `json:"fiscalYear"`
	Period                                        string `json:"period"`
	ReportedCurrency                              string `json:"reportedCurrency"`
	GrowthCashAndCashEquivalents                  string `json:"growthCashAndCashEquivalents"`
	GrowthShortTermInvestments                    string `json:"growthShortTermInvestments"`
	GrowthCashAndShortTermInvestments             string `json:"growthCashAndShortTermInvestments"`
	GrowthNetReceivables                          string `json:"growthNetReceivables"`
	GrowthInventory                               string `json:"growthInventory"`
	GrowthOtherCurrentAssets                      string `json:"growthOtherCurrentAssets"`
	GrowthTotalCurrentAssets                      string `json:"growthTotalCurrentAssets"`
	GrowthPropertyPlantEquipmentNet               string `json:"growthPropertyPlantEquipmentNet"`
	GrowthGoodwill                                string `json:"growthGoodwill"`
	GrowthIntangibleAssets                        string `json:"growthIntangibleAssets"`
	GrowthGoodwillAndIntangibleAssets             string `json:"growthGoodwillAndIntangibleAssets"`
	GrowthLongTermInvestments                     string `json:"growthLongTermInvestments"`
	GrowthTaxAssets                               string `json:"growthTaxAssets"`
	GrowthOtherNonCurrentAssets                   string `json:"growthOtherNonCurrentAssets"`
	GrowthTotalNonCurrentAssets                   string `json:"growthTotalNonCurrentAssets"`
	GrowthOtherAssets                             string `json:"growthOtherAssets"`
	GrowthTotalAssets                             string `json:"growthTotalAssets"`
	GrowthAccountPayables                         string `json:"growthAccountPayables"`
	GrowthShortTermDebt                           string `json:"growthShortTermDebt"`
	GrowthTaxPayables                             string `json:"growthTaxPayables"`
	GrowthDeferredRevenue                         string `json:"growthDeferredRevenue"`
	GrowthOtherCurrentLiabilities                 string `json:"growthOtherCurrentLiabilities"`
	GrowthTotalCurrentLiabilities                 string `json:"growthTotalCurrentLiabilities"`
	GrowthLongTermDebt                            string `json:"growthLongTermDebt"`
	GrowthDeferredRevenueNonCurrent               string `json:"growthDeferredRevenueNonCurrent"`
	GrowthDeferredTaxLiabilitiesNonCurrent        string `json:"growthDeferredTaxLiabilitiesNonCurrent"`
	GrowthOtherNonCurrentLiabilities              string `json:"growthOtherNonCurrentLiabilities"`
	GrowthTotalNonCurrentLiabilities              string `json:"growthTotalNonCurrentLiabilities"`
	GrowthOtherLiabilities                        string `json:"growthOtherLiabilities"`
	GrowthTotalLiabilities                        string `json:"growthTotalLiabilities"`
	GrowthPreferredStock                          string `json:"growthPreferredStock"`
	GrowthCommonStock                             string `json:"growthCommonStock"`
	GrowthRetainedEarnings                        string `json:"growthRetainedEarnings"`
	GrowthAccumulatedOtherComprehensiveIncomeLoss string `json:"growthAccumulatedOtherComprehensiveIncomeLoss"`
	GrowthOthertotalStockholdersEquity            string `json:"growthOthertotalStockholdersEquity"`
	GrowthTotalStockholdersEquity                 string `json:"growthTotalStockholdersEquity"`
	GrowthMinorityInterest                        string `json:"growthMinorityInterest"`
	GrowthTotalEquity                             string `json:"growthTotalEquity"`
	GrowthTotalLiabilitiesAndStockholdersEquity   string `json:"growthTotalLiabilitiesAndStockholdersEquity"`
	GrowthTotalInvestments                        string `json:"growthTotalInvestments"`
	GrowthTotalDebt                               string `json:"growthTotalDebt"`
	GrowthNetDebt                                 string `json:"growthNetDebt"`
	GrowthAccountsReceivables                     string `json:"growthAccountsReceivables"`
	GrowthOtherReceivables                        string `json:"growthOtherReceivables"`
	GrowthPrepaids                                string `json:"growthPrepaids"`
	GrowthTotalPayables                           string `json:"growthTotalPayables"`
	GrowthOtherPayables                           string `json:"growthOtherPayables"`
	GrowthAccruedExpenses                         string `json:"growthAccruedExpenses"`
	GrowthCapitalLeaseObligationsCurrent          string `json:"growthCapitalLeaseObligationsCurrent"`
	GrowthAdditionalPaidInCapital                 string `json:"growthAdditionalPaidInCapital"`
	GrowthTreasuryStock                           string `json:"growthTreasuryStock"`
}

// GetBalanceSheetGrowthBulk retrieves growth data across multiple companies' balance sheets
func (c *Client) GetBalanceSheetGrowthBulk(params BalanceSheetGrowthBulkParams) ([]BalanceSheetGrowthBulkResponse, error) {
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

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/balance-sheet-statement-growth-bulk", urlParams)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()



	// Parse the response
	var result []BalanceSheetGrowthBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
