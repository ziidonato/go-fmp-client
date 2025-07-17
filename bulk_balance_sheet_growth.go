package go_fmp

import (
	"fmt"
	"time"
)

// BulkBalanceSheetGrowthParams represents the parameters for the Bulk Balance Sheet Growth API
type BulkBalanceSheetGrowthParams struct {
	Year   int    `json:"year"`   // Required: year (e.g., 2023)
	Period string `json:"period"` // Required: period ("annual" or "quarterly")
}

// BulkBalanceSheetGrowthResponse represents the response from the Bulk Balance Sheet Growth API
type BulkBalanceSheetGrowthResponse struct {
	Symbol                                        string    `json:"symbol"`
	Date                                          time.Time `json:"date"`
	CalendarYear                                  string    `json:"calendarYear"`
	Period                                        string    `json:"period"`
	GrowthCashAndCashEquivalents                  float64   `json:"growthCashAndCashEquivalents"`
	GrowthShortTermInvestments                    float64   `json:"growthShortTermInvestments"`
	GrowthCashAndShortTermInvestments             float64   `json:"growthCashAndShortTermInvestments"`
	GrowthNetReceivables                          float64   `json:"growthNetReceivables"`
	GrowthInventory                               float64   `json:"growthInventory"`
	GrowthOtherCurrentAssets                      float64   `json:"growthOtherCurrentAssets"`
	GrowthTotalCurrentAssets                      float64   `json:"growthTotalCurrentAssets"`
	GrowthPropertyPlantEquipmentNet               float64   `json:"growthPropertyPlantEquipmentNet"`
	GrowthGoodwill                                float64   `json:"growthGoodwill"`
	GrowthIntangibleAssets                        float64   `json:"growthIntangibleAssets"`
	GrowthGoodwillAndIntangibleAssets             float64   `json:"growthGoodwillAndIntangibleAssets"`
	GrowthLongTermInvestments                     float64   `json:"growthLongTermInvestments"`
	GrowthTaxAssets                               float64   `json:"growthTaxAssets"`
	GrowthOtherNonCurrentAssets                   float64   `json:"growthOtherNonCurrentAssets"`
	GrowthTotalNonCurrentAssets                   float64   `json:"growthTotalNonCurrentAssets"`
	GrowthOtherAssets                             float64   `json:"growthOtherAssets"`
	GrowthTotalAssets                             float64   `json:"growthTotalAssets"`
	GrowthAccountPayables                         float64   `json:"growthAccountPayables"`
	GrowthShortTermDebt                           float64   `json:"growthShortTermDebt"`
	GrowthTaxPayables                             float64   `json:"growthTaxPayables"`
	GrowthDeferredRevenue                         float64   `json:"growthDeferredRevenue"`
	GrowthOtherCurrentLiabilities                 float64   `json:"growthOtherCurrentLiabilities"`
	GrowthTotalCurrentLiabilities                 float64   `json:"growthTotalCurrentLiabilities"`
	GrowthLongTermDebt                            float64   `json:"growthLongTermDebt"`
	GrowthDeferredRevenueNonCurrent               float64   `json:"growthDeferredRevenueNonCurrent"`
	GrowthDeferredTaxLiabilitiesNonCurrent        float64   `json:"growthDeferredTaxLiabilitiesNonCurrent"`
	GrowthOtherNonCurrentLiabilities              float64   `json:"growthOtherNonCurrentLiabilities"`
	GrowthTotalNonCurrentLiabilities              float64   `json:"growthTotalNonCurrentLiabilities"`
	GrowthOtherLiabilities                        float64   `json:"growthOtherLiabilities"`
	GrowthTotalLiabilities                        float64   `json:"growthTotalLiabilities"`
	GrowthCommonStock                             float64   `json:"growthCommonStock"`
	GrowthRetainedEarnings                        float64   `json:"growthRetainedEarnings"`
	GrowthAccumulatedOtherComprehensiveIncomeLoss float64   `json:"growthAccumulatedOtherComprehensiveIncomeLoss"`
	GrowthOtherTotalStockholdersEquity            float64   `json:"growthOthertotalStockholdersEquity"`
	GrowthTotalStockholdersEquity                 float64   `json:"growthTotalStockholdersEquity"`
	GrowthTotalLiabilitiesAndStockholdersEquity   float64   `json:"growthTotalLiabilitiesAndStockholdersEquity"`
	GrowthTotalInvestments                        float64   `json:"growthTotalInvestments"`
	GrowthTotalDebt                               float64   `json:"growthTotalDebt"`
	GrowthNetDebt                                 float64   `json:"growthNetDebt"`
}

// BulkBalanceSheetGrowth retrieves balance sheet growth data for multiple companies
func (c *Client) BulkBalanceSheetGrowth(params BulkBalanceSheetGrowthParams) ([]BulkBalanceSheetGrowthResponse, error) {
	urlParams := map[string]string{
		"year":   fmt.Sprintf("%d", params.Year),
		"period": params.Period,
	}

	var result []BulkBalanceSheetGrowthResponse
	err := c.doRequest(c.BaseURL+"/bulk-balance-sheet-growth", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk balance sheet growth: %w", err)
	}

	return result, nil
}
