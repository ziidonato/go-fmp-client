package go_fmp

import (
	"fmt"
	"time"
)

// IncomeStatementGrowthParams represents the parameters for the Income Statement Growth API
type IncomeStatementGrowthParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return (default: 20)
}

// IncomeStatementGrowthResponse represents the response from the Income Statement Growth API
type IncomeStatementGrowthResponse struct {
	Date                                      time.Time  `json:"date"`
	Symbol                                    string  `json:"symbol"`
	ReportedCurrency                          string  `json:"reportedCurrency"`
	CIK                                       string  `json:"cik"`
	FillingDate                               string  `json:"fillingDate"`
	AcceptedDate                              string  `json:"acceptedDate"`
	CalendarYear                              string  `json:"calendarYear"`
	Period                                    string  `json:"period"`
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
	GrowthInterestExpense                     float64 `json:"growthInterestExpense"`
	GrowthDepreciationAndAmortization         float64 `json:"growthDepreciationAndAmortization"`
	GrowthEBITDA                              float64 `json:"growthEBITDA"`
	GrowthEBITDARatio                         float64 `json:"growthEBITDARatio"`
	GrowthOperatingIncome                     float64 `json:"growthOperatingIncome"`
	GrowthOperatingIncomeRatio                float64 `json:"growthOperatingIncomeRatio"`
	GrowthTotalOtherIncomeExpensesNet         float64 `json:"growthTotalOtherIncomeExpensesNet"`
	GrowthIncomeBeforeTax                     float64 `json:"growthIncomeBeforeTax"`
	GrowthIncomeBeforeTaxRatio                float64 `json:"growthIncomeBeforeTaxRatio"`
	GrowthIncomeTaxExpense                    float64 `json:"growthIncomeTaxExpense"`
	GrowthNetIncome                           float64 `json:"growthNetIncome"`
	GrowthNetIncomeRatio                      float64 `json:"growthNetIncomeRatio"`
	GrowthEPS                                 float64 `json:"growthEPS"`
	GrowthEPSDiluted                          float64 `json:"growthEPSDiluted"`
	GrowthWeightedAverageShsOut               float64 `json:"growthWeightedAverageShsOut"`
	GrowthWeightedAverageShsOutDil            float64 `json:"growthWeightedAverageShsOutDil"`
	Link                                      string  `json:"link"`
	FinalLink                                 string  `json:"finalLink"`
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

	if params.Period != nil {
		urlParams["period"] = *params.Period
	}

	var result []IncomeStatementGrowthResponse
	err := c.doRequest(c.BaseURL+"/income-statement-growth", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
