package go_fmp

import (
	"fmt"
	"time"
)

// BulkCashFlowStatementGrowthParams represents the parameters for the Bulk Cash Flow Statement Growth API
type BulkCashFlowStatementGrowthParams struct {
	Year   int    `json:"year"`   // Required: year (e.g., 2023)
	Period string `json:"period"` // Required: period ("annual" or "quarterly")
}

// BulkCashFlowStatementGrowthResponse represents the response from the Bulk Cash Flow Statement Growth API
type BulkCashFlowStatementGrowthResponse struct {
	Symbol                                              string    `json:"symbol"`
	Date                                                time.Time `json:"date"`
	CalendarYear                                        string    `json:"calendarYear"`
	Period                                              string    `json:"period"`
	GrowthNetIncome                                     float64   `json:"growthNetIncome"`
	GrowthDepreciationAndAmortization                   float64   `json:"growthDepreciationAndAmortization"`
	GrowthDeferredIncomeTax                             float64   `json:"growthDeferredIncomeTax"`
	GrowthStockBasedCompensation                        float64   `json:"growthStockBasedCompensation"`
	GrowthChangeInWorkingCapital                        float64   `json:"growthChangeInWorkingCapital"`
	GrowthAccountsReceivables                           float64   `json:"growthAccountsReceivables"`
	GrowthInventory                                     float64   `json:"growthInventory"`
	GrowthAccountsPayables                              float64   `json:"growthAccountsPayables"`
	GrowthOtherWorkingCapital                           float64   `json:"growthOtherWorkingCapital"`
	GrowthOtherNonCashItems                             float64   `json:"growthOtherNonCashItems"`
	GrowthNetCashProvidedByOperatingActivites           float64   `json:"growthNetCashProvidedByOperatingActivites"`
	GrowthInvestmentsInPropertyPlantAndEquipment        float64   `json:"growthInvestmentsInPropertyPlantAndEquipment"`
	GrowthAcquisitionsNet                               float64   `json:"growthAcquisitionsNet"`
	GrowthPurchasesOfInvestments                        float64   `json:"growthPurchasesOfInvestments"`
	GrowthSalesMaturitiesOfInvestments                  float64   `json:"growthSalesMaturitiesOfInvestments"`
	GrowthOtherInvestingActivites                       float64   `json:"growthOtherInvestingActivites"`
	GrowthNetCashUsedForInvestingActivites              float64   `json:"growthNetCashUsedForInvestingActivites"`
	GrowthDebtRepayment                                 float64   `json:"growthDebtRepayment"`
	GrowthCommonStockIssued                             float64   `json:"growthCommonStockIssued"`
	GrowthCommonStockRepurchased                        float64   `json:"growthCommonStockRepurchased"`
	GrowthDividendsPaid                                 float64   `json:"growthDividendsPaid"`
	GrowthOtherFinancingActivites                       float64   `json:"growthOtherFinancingActivites"`
	GrowthNetCashUsedProvidedByFinancingActivities      float64   `json:"growthNetCashUsedProvidedByFinancingActivities"`
	GrowthEffectOfForexChangesOnCash                    float64   `json:"growthEffectOfForexChangesOnCash"`
	GrowthNetChangeInCash                               float64   `json:"growthNetChangeInCash"`
	GrowthCashAtEndOfPeriod                             float64   `json:"growthCashAtEndOfPeriod"`
	GrowthCashAtBeginningOfPeriod                       float64   `json:"growthCashAtBeginningOfPeriod"`
	GrowthOperatingCashFlow                             float64   `json:"growthOperatingCashFlow"`
	GrowthCapitalExpenditure                            float64   `json:"growthCapitalExpenditure"`
	GrowthFreeCashFlow                                  float64   `json:"growthFreeCashFlow"`
}

// BulkCashFlowStatementGrowth retrieves cash flow statement growth data for multiple companies
func (c *Client) BulkCashFlowStatementGrowth(params BulkCashFlowStatementGrowthParams) ([]BulkCashFlowStatementGrowthResponse, error) {
	urlParams := map[string]string{
		"year":   fmt.Sprintf("%d", params.Year),
		"period": params.Period,
	}

	var result []BulkCashFlowStatementGrowthResponse
	err := c.doRequest(c.BaseURL+"/bulk-cash-flow-statement-growth", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk cash flow statement growth: %w", err)
	}

	return result, nil
}
