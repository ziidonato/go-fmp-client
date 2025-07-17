package go_fmp

import (
	"fmt"
	"time"
)

// BulkIncomeStatementGrowthParams represents the parameters for the Bulk Income Statement Growth API
type BulkIncomeStatementGrowthParams struct {
	Year   int    `json:"year"`   // Required: year (e.g., 2023)
	Period string `json:"period"` // Required: period ("annual" or "quarterly")
}

// BulkIncomeStatementGrowthResponse represents the response from the Bulk Income Statement Growth API
type BulkIncomeStatementGrowthResponse struct {
	Symbol                                  string    `json:"symbol"`
	Date                                    time.Time `json:"date"`
	CalendarYear                            string    `json:"calendarYear"`
	Period                                  string    `json:"period"`
	GrowthRevenue                           float64   `json:"growthRevenue"`
	GrowthCostOfRevenue                     float64   `json:"growthCostOfRevenue"`
	GrowthGrossProfit                       float64   `json:"growthGrossProfit"`
	GrowthGrossProfitRatio                  float64   `json:"growthGrossProfitRatio"`
	GrowthResearchAndDevelopmentExpenses    float64   `json:"growthResearchAndDevelopmentExpenses"`
	GrowthGeneralAndAdministrativeExpenses  float64   `json:"growthGeneralAndAdministrativeExpenses"`
	GrowthSellingAndMarketingExpenses       float64   `json:"growthSellingAndMarketingExpenses"`
	GrowthOtherExpenses                     float64   `json:"growthOtherExpenses"`
	GrowthOperatingExpenses                 float64   `json:"growthOperatingExpenses"`
	GrowthCostAndExpenses                   float64   `json:"growthCostAndExpenses"`
	GrowthInterestExpense                   float64   `json:"growthInterestExpense"`
	GrowthDepreciationAndAmortization       float64   `json:"growthDepreciationAndAmortization"`
	GrowthEBITDA                            float64   `json:"growthEBITDA"`
	GrowthEBITDARatio                       float64   `json:"growthEBITDARatio"`
	GrowthOperatingIncome                   float64   `json:"growthOperatingIncome"`
	GrowthOperatingIncomeRatio              float64   `json:"growthOperatingIncomeRatio"`
	GrowthTotalOtherIncomeExpensesNet       float64   `json:"growthTotalOtherIncomeExpensesNet"`
	GrowthIncomeBeforeTax                   float64   `json:"growthIncomeBeforeTax"`
	GrowthIncomeBeforeTaxRatio              float64   `json:"growthIncomeBeforeTaxRatio"`
	GrowthIncomeTaxExpense                  float64   `json:"growthIncomeTaxExpense"`
	GrowthNetIncome                         float64   `json:"growthNetIncome"`
	GrowthNetIncomeRatio                    float64   `json:"growthNetIncomeRatio"`
	GrowthEPS                               float64   `json:"growthEPS"`
	GrowthEPSDiluted                        float64   `json:"growthEPSDiluted"`
	GrowthWeightedAverageShsOut             float64   `json:"growthWeightedAverageShsOut"`
	GrowthWeightedAverageShsOutDil          float64   `json:"growthWeightedAverageShsOutDil"`
}

// BulkIncomeStatementGrowth retrieves income statement growth data for multiple companies
func (c *Client) BulkIncomeStatementGrowth(params BulkIncomeStatementGrowthParams) ([]BulkIncomeStatementGrowthResponse, error) {
	urlParams := map[string]string{
		"year":   fmt.Sprintf("%d", params.Year),
		"period": params.Period,
	}

	var result []BulkIncomeStatementGrowthResponse
	err := c.doRequest(c.BaseURL+"/bulk-income-statement-growth", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk income statement growth: %w", err)
	}

	return result, nil
}
