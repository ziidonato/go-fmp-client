package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

	// Build the URL
	baseURL := "https://financialmodelingprep.com/stable/income-statement-growth-bulk"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Add query parameters
	q := u.Query()
	q.Set("year", params.Year)
	q.Set("period", params.Period)
	u.RawQuery = q.Encode()

	// Add API key if available
	if c.APIKey != "" {
		q.Set("apikey", c.APIKey)
		u.RawQuery = q.Encode()
	}

	// Create the request
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("User-Agent", "fmp-go-client")
	req.Header.Set("Accept", "application/json")

	// Make the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Parse the response
	var result []IncomeStatementGrowthBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
