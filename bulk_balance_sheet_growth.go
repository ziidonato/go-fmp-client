package go_fmp

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// BalanceSheetGrowthBulkParams represents the parameters for the Balance Sheet Growth Bulk API
type BalanceSheetGrowthBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period Period `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)
}

// BalanceSheetGrowthBulkResponse represents the response from the Balance Sheet Growth Bulk API
type BalanceSheetGrowthBulkResponse struct {
	Symbol                                        string `json:"symbol"`
	Date time.Time `json:"date"`
	FiscalYear                                    string `json:"fiscalYear"`
	Period Period `json:"period"`
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

	// Build the URL
	baseURL := c.BaseURL + "/balance-sheet-statement-growth-bulk"
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
	var result []BalanceSheetGrowthBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
