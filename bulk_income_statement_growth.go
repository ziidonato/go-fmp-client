package go_fmp

import (
	"encoding/json"
	"fmt"
)

// IncomeStatementGrowthBulkParams represents the parameters for the Bulk Income Statement Growth API
type IncomeStatementGrowthBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period string `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)
}

// IncomeStatementGrowthBulkResponse represents the response from the Bulk Income Statement Growth API
type IncomeStatementGrowthBulkResponse struct {
	Symbol                                    string `json:"symbol"`
	Date                                      string `json:"date"`
	FiscalYear                                string `json:"fiscalYear"`
	Period                                    string `json:"period"`
	ReportedCurrency                          string `json:"reportedCurrency"`
	GrowthRevenue                             string `json:"growthRevenue"`
	GrowthCostOfRevenue                       string `json:"growthCostOfRevenue"`
	GrowthGrossProfit                         string `json:"growthGrossProfit"`
	GrowthGrossProfitRatio                    string `json:"growthGrossProfitRatio"`
	GrowthResearchAndDevelopmentExpenses      string `json:"growthResearchAndDevelopmentExpenses"`
	GrowthGeneralAndAdministrativeExpenses    string `json:"growthGeneralAndAdministrativeExpenses"`
	GrowthSellingAndMarketingExpenses         string `json:"growthSellingAndMarketingExpenses"`
	GrowthOtherExpenses                       string `json:"growthOtherExpenses"`
	GrowthOperatingExpenses                   string `json:"growthOperatingExpenses"`
	GrowthCostAndExpenses                     string `json:"growthCostAndExpenses"`
	GrowthInterestIncome                      string `json:"growthInterestIncome"`
	GrowthInterestExpense                     string `json:"growthInterestExpense"`
	GrowthDepreciationAndAmortization         string `json:"growthDepreciationAndAmortization"`
	GrowthEBITDA                              string `json:"growthEBITDA"`
	GrowthOperatingIncome                     string `json:"growthOperatingIncome"`
	GrowthIncomeBeforeTax                     string `json:"growthIncomeBeforeTax"`
	GrowthIncomeTaxExpense                    string `json:"growthIncomeTaxExpense"`
	GrowthNetIncome                           string `json:"growthNetIncome"`
	GrowthEPS                                 string `json:"growthEPS"`
	GrowthEPSDiluted                          string `json:"growthEPSDiluted"`
	GrowthWeightedAverageShsOut               string `json:"growthWeightedAverageShsOut"`
	GrowthWeightedAverageShsOutDil            string `json:"growthWeightedAverageShsOutDil"`
	GrowthEBIT                                string `json:"growthEBIT"`
	GrowthNonOperatingIncomeExcludingInterest string `json:"growthNonOperatingIncomeExcludingInterest"`
	GrowthNetInterestIncome                   string `json:"growthNetInterestIncome"`
	GrowthTotalOtherIncomeExpensesNet         string `json:"growthTotalOtherIncomeExpensesNet"`
	GrowthNetIncomeFromContinuingOperations   string `json:"growthNetIncomeFromContinuingOperations"`
	GrowthOtherAdjustmentsToNetIncome         string `json:"growthOtherAdjustmentsToNetIncome"`
	GrowthNetIncomeDeductions                 string `json:"growthNetIncomeDeductions"`
}

// GetIncomeStatementGrowthBulk retrieves growth data for income statements across multiple companies
func (c *Client) GetIncomeStatementGrowthBulk(params IncomeStatementGrowthBulkParams) ([]IncomeStatementGrowthBulkResponse, error) {
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

	return doRequest[[]IncomeStatementGrowthBulkResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
