package go_fmp

import (
	"fmt"
)

// IncomeStatementGrowthParams represents the parameters for the Income Statement Growth API
type IncomeStatementGrowthParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
	Period string `json:"period"` // Optional: Period type - "Q1,Q2,Q3,Q4,FY,annual,quarter"
}

// IncomeStatementGrowthResponse represents the response from the Income Statement Growth API
type IncomeStatementGrowthResponse struct {
	Symbol                                    string  `json:"symbol"`
	Date                                      string  `json:"date"`
	FiscalYear                                string  `json:"fiscalYear"`
	Period                                    string  `json:"period"`
	ReportedCurrency                          string  `json:"reportedCurrency"`
	GrowthRevenue                             float64 `json:"growthRevenue"`
	GrowthCostOfRevenue                       float64 `json:"growthCostOfRevenue"`
	GrowthGrossProfit                         float64 `json:"growthGrossProfit"`
	GrowthGrossProfitRatio                    float64 `json:"growthGrossProfitRatio"`
	GrowthResearchAndDevelopmentExpenses      float64 `json:"growthResearchAndDevelopmentExpenses"`
	GrowthGeneralAndAdministrativeExpenses    float64 `json:"growthGeneralAndAdministrativeExpenses"`
	GrowthSellingAndMarketingExpenses         float64 `json:"growthSellingAndMarketingExpenses"`
	GrowthOtherExpenses                       float64 `json:"growthOtherExpenses"`
	GrowthOperatingExpenses                   float64 `json:"growthOperatingExpenses"`
	GrowthCostAndExpenses                     float64 `json:"growthCostAndExpenses"`
	GrowthInterestIncome                      float64 `json:"growthInterestIncome"`
	GrowthInterestExpense                     float64 `json:"growthInterestExpense"`
	GrowthDepreciationAndAmortization         float64 `json:"growthDepreciationAndAmortization"`
	GrowthEBITDA                              float64 `json:"growthEBITDA"`
	GrowthOperatingIncome                     float64 `json:"growthOperatingIncome"`
	GrowthIncomeBeforeTax                     float64 `json:"growthIncomeBeforeTax"`
	GrowthIncomeTaxExpense                    float64 `json:"growthIncomeTaxExpense"`
	GrowthNetIncome                           float64 `json:"growthNetIncome"`
	GrowthEPS                                 float64 `json:"growthEPS"`
	GrowthEPSDiluted                          float64 `json:"growthEPSDiluted"`
	GrowthWeightedAverageShsOut               float64 `json:"growthWeightedAverageShsOut"`
	GrowthWeightedAverageShsOutDil            float64 `json:"growthWeightedAverageShsOutDil"`
	GrowthEBIT                                float64 `json:"growthEBIT"`
	GrowthNonOperatingIncomeExcludingInterest float64 `json:"growthNonOperatingIncomeExcludingInterest"`
	GrowthNetInterestIncome                   float64 `json:"growthNetInterestIncome"`
	GrowthTotalOtherIncomeExpensesNet         float64 `json:"growthTotalOtherIncomeExpensesNet"`
	GrowthNetIncomeFromContinuingOperations   float64 `json:"growthNetIncomeFromContinuingOperations"`
	GrowthOtherAdjustmentsToNetIncome         float64 `json:"growthOtherAdjustmentsToNetIncome"`
	GrowthNetIncomeDeductions                 float64 `json:"growthNetIncomeDeductions"`
}

// IncomeStatementGrowth retrieves income statement growth data for a specific stock symbol
func (c *Client) IncomeStatementGrowth(params IncomeStatementGrowthParams) ([]IncomeStatementGrowthResponse, error) {
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

	if params.Period != "" {
		urlParams["period"] = params.Period
	}

	var result []IncomeStatementGrowthResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/income-statement-growth", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
