package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// IncomeStatementBulkParams represents the parameters for the Bulk Income Statement API
type IncomeStatementBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period string `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)
}

// IncomeStatementBulkResponse represents the response from the Bulk Income Statement API
type IncomeStatementBulkResponse struct {
	Date                                    string `json:"date"`
	Symbol                                  string `json:"symbol"`
	ReportedCurrency                        string `json:"reportedCurrency"`
	CIK                                     string `json:"cik"`
	FilingDate                              string `json:"filingDate"`
	AcceptedDate                            string `json:"acceptedDate"`
	FiscalYear                              string `json:"fiscalYear"`
	Period                                  string `json:"period"`
	Revenue                                 string `json:"revenue"`
	CostOfRevenue                           string `json:"costOfRevenue"`
	GrossProfit                             string `json:"grossProfit"`
	ResearchAndDevelopmentExpenses          string `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses        string `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses             string `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdministrativeExpenses string `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                           string `json:"otherExpenses"`
	OperatingExpenses                       string `json:"operatingExpenses"`
	CostAndExpenses                         string `json:"costAndExpenses"`
	NetInterestIncome                       string `json:"netInterestIncome"`
	InterestIncome                          string `json:"interestIncome"`
	InterestExpense                         string `json:"interestExpense"`
	DepreciationAndAmortization             string `json:"depreciationAndAmortization"`
	EBITDA                                  string `json:"ebitda"`
	EBIT                                    string `json:"ebit"`
	NonOperatingIncomeExcludingInterest     string `json:"nonOperatingIncomeExcludingInterest"`
	OperatingIncome                         string `json:"operatingIncome"`
	TotalOtherIncomeExpensesNet             string `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                         string `json:"incomeBeforeTax"`
	IncomeTaxExpense                        string `json:"incomeTaxExpense"`
	NetIncomeFromContinuingOperations       string `json:"netIncomeFromContinuingOperations"`
	NetIncomeFromDiscontinuedOperations     string `json:"netIncomeFromDiscontinuedOperations"`
	OtherAdjustmentsToNetIncome             string `json:"otherAdjustmentsToNetIncome"`
	NetIncome                               string `json:"netIncome"`
	NetIncomeDeductions                     string `json:"netIncomeDeductions"`
	BottomLineNetIncome                     string `json:"bottomLineNetIncome"`
	EPS                                     string `json:"eps"`
	EPSDiluted                              string `json:"epsDiluted"`
	WeightedAverageShsOut                   string `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil                string `json:"weightedAverageShsOutDil"`
}

// GetIncomeStatementBulk retrieves detailed income statement data in bulk
func (c *Client) GetIncomeStatementBulk(params IncomeStatementBulkParams) ([]IncomeStatementBulkResponse, error) {
	// Validate required parameters
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}
	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	// Build the URL
	baseURL := "https://financialmodelingprep.com/stable/income-statement-bulk"
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
	var result []IncomeStatementBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
