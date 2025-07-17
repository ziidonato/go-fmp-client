package go_fmp

import (
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// CashFlowStatementBulkParams represents the parameters for the Cash Flow Statement Bulk API
type CashFlowStatementBulkParams struct {
	Year   string `json:"year"`   // Required: year (e.g., "2024")
	Period Period `json:"period"` // Required: period (Q1,Q2,Q3,Q4,FY)
}

// CashFlowStatementBulkResponse represents the response from the Cash Flow Statement Bulk API
type CashFlowStatementBulkResponse struct {
	Date time.Time `json:"date"`
	Symbol                                 string `json:"symbol"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	CIK                                    string `json:"cik"`
	FilingDate                             string `json:"filingDate"`
	AcceptedDate                           string `json:"acceptedDate"`
	FiscalYear                             string `json:"fiscalYear"`
	Period Period `json:"period"`
	NetIncome                              string `json:"netIncome"`
	DepreciationAndAmortization            string `json:"depreciationAndAmortization"`
	DeferredIncomeTax                      string `json:"deferredIncomeTax"`
	StockBasedCompensation                 string `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                 string `json:"changeInWorkingCapital"`
	AccountsReceivables                    string `json:"accountsReceivables"`
	Inventory                              string `json:"inventory"`
	AccountsPayables                       string `json:"accountsPayables"`
	OtherWorkingCapital                    string `json:"otherWorkingCapital"`
	OtherNonCashItems                      string `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities   string `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment string `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                        string `json:"acquisitionsNet"`
	PurchasesOfInvestments                 string `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments           string `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivities               string `json:"otherInvestingActivities"`
	NetCashProvidedByInvestingActivities   string `json:"netCashProvidedByInvestingActivities"`
	NetDebtIssuance                        string `json:"netDebtIssuance"`
	LongTermNetDebtIssuance                string `json:"longTermNetDebtIssuance"`
	ShortTermNetDebtIssuance               string `json:"shortTermNetDebtIssuance"`
	NetStockIssuance                       string `json:"netStockIssuance"`
	NetCommonStockIssuance                 string `json:"netCommonStockIssuance"`
	CommonStockIssuance                    string `json:"commonStockIssuance"`
	CommonStockRepurchased                 string `json:"commonStockRepurchased"`
	NetPreferredStockIssuance              string `json:"netPreferredStockIssuance"`
	NetDividendsPaid                       string `json:"netDividendsPaid"`
	CommonDividendsPaid                    string `json:"commonDividendsPaid"`
	PreferredDividendsPaid                 string `json:"preferredDividendsPaid"`
	OtherFinancingActivities               string `json:"otherFinancingActivities"`
	NetCashProvidedByFinancingActivities   string `json:"netCashProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash             string `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                        string `json:"netChangeInCash"`
	CashAtEndOfPeriod                      string `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                string `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                      string `json:"operatingCashFlow"`
	CapitalExpenditure                     string `json:"capitalExpenditure"`
	FreeCashFlow                           string `json:"freeCashFlow"`
	IncomeTaxesPaid                        string `json:"incomeTaxesPaid"`
	InterestPaid                           string `json:"interestPaid"`
}

// GetCashFlowStatementBulk retrieves detailed cash flow reports for a wide range of companies
func (c *Client) GetCashFlowStatementBulk(params CashFlowStatementBulkParams) ([]CashFlowStatementBulkResponse, error) {
	// Validate required parameters
	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}
	if params.Period == "" {
		return nil, fmt.Errorf("period parameter is required")
	}

	// Build the URL
	baseURL := c.BaseURL + "/cash-flow-statement-bulk"
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
	var result []CashFlowStatementBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
