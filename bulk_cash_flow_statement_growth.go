package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// CashFlowStatementGrowthBulkParams represents the parameters for the Cash Flow Statement Growth Bulk API
type CashFlowStatementGrowthBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period string `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)
}

// CashFlowStatementGrowthBulkResponse represents the response from the Cash Flow Statement Growth Bulk API
type CashFlowStatementGrowthBulkResponse struct {
	Symbol                                         string `json:"symbol"`
	Date                                           string `json:"date"`
	FiscalYear                                     string `json:"fiscalYear"`
	Period                                         string `json:"period"`
	ReportedCurrency                               string `json:"reportedCurrency"`
	GrowthNetIncome                                string `json:"growthNetIncome"`
	GrowthDepreciationAndAmortization              string `json:"growthDepreciationAndAmortization"`
	GrowthDeferredIncomeTax                        string `json:"growthDeferredIncomeTax"`
	GrowthStockBasedCompensation                   string `json:"growthStockBasedCompensation"`
	GrowthChangeInWorkingCapital                   string `json:"growthChangeInWorkingCapital"`
	GrowthAccountsReceivables                      string `json:"growthAccountsReceivables"`
	GrowthInventory                                string `json:"growthInventory"`
	GrowthAccountsPayables                         string `json:"growthAccountsPayables"`
	GrowthOtherWorkingCapital                      string `json:"growthOtherWorkingCapital"`
	GrowthOtherNonCashItems                        string `json:"growthOtherNonCashItems"`
	GrowthNetCashProvidedByOperatingActivites      string `json:"growthNetCashProvidedByOperatingActivites"`
	GrowthInvestmentsInPropertyPlantAndEquipment   string `json:"growthInvestmentsInPropertyPlantAndEquipment"`
	GrowthAcquisitionsNet                          string `json:"growthAcquisitionsNet"`
	GrowthPurchasesOfInvestments                   string `json:"growthPurchasesOfInvestments"`
	GrowthSalesMaturitiesOfInvestments             string `json:"growthSalesMaturitiesOfInvestments"`
	GrowthOtherInvestingActivites                  string `json:"growthOtherInvestingActivites"`
	GrowthNetCashUsedForInvestingActivites         string `json:"growthNetCashUsedForInvestingActivites"`
	GrowthDebtRepayment                            string `json:"growthDebtRepayment"`
	GrowthCommonStockIssued                        string `json:"growthCommonStockIssued"`
	GrowthCommonStockRepurchased                   string `json:"growthCommonStockRepurchased"`
	GrowthDividendsPaid                            string `json:"growthDividendsPaid"`
	GrowthOtherFinancingActivites                  string `json:"growthOtherFinancingActivites"`
	GrowthNetCashUsedProvidedByFinancingActivities string `json:"growthNetCashUsedProvidedByFinancingActivities"`
	GrowthEffectOfForexChangesOnCash               string `json:"growthEffectOfForexChangesOnCash"`
	GrowthNetChangeInCash                          string `json:"growthNetChangeInCash"`
	GrowthCashAtEndOfPeriod                        string `json:"growthCashAtEndOfPeriod"`
	GrowthCashAtBeginningOfPeriod                  string `json:"growthCashAtBeginningOfPeriod"`
	GrowthOperatingCashFlow                        string `json:"growthOperatingCashFlow"`
	GrowthCapitalExpenditure                       string `json:"growthCapitalExpenditure"`
	GrowthFreeCashFlow                             string `json:"growthFreeCashFlow"`
	GrowthNetDebtIssuance                          string `json:"growthNetDebtIssuance"`
	GrowthLongTermNetDebtIssuance                  string `json:"growthLongTermNetDebtIssuance"`
	GrowthShortTermNetDebtIssuance                 string `json:"growthShortTermNetDebtIssuance"`
	GrowthNetStockIssuance                         string `json:"growthNetStockIssuance"`
	GrowthPreferredDividendsPaid                   string `json:"growthPreferredDividendsPaid"`
	GrowthIncomeTaxesPaid                          string `json:"growthIncomeTaxesPaid"`
	GrowthInterestPaid                             string `json:"growthInterestPaid"`
}

// GetCashFlowStatementGrowthBulk retrieves bulk growth data for cash flow statements
func (c *Client) GetCashFlowStatementGrowthBulk(params CashFlowStatementGrowthBulkParams) ([]CashFlowStatementGrowthBulkResponse, error) {
	// Validate required parameters
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}
	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	// Build the URL
	baseURL := c.BaseURL + "/cash-flow-statement-growth-bulk"
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
	var result []CashFlowStatementGrowthBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
