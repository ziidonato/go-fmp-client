package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// KeyMetricsTTMBulkResponse represents the response from the Key Metrics TTM Bulk API
type KeyMetricsTTMBulkResponse struct {
	Symbol                                    string `json:"symbol"`
	MarketCap                                 string `json:"marketCap"`
	EnterpriseValueTTM                        string `json:"enterpriseValueTTM"`
	EVToSalesTTM                              string `json:"evToSalesTTM"`
	EVToOperatingCashFlowTTM                  string `json:"evToOperatingCashFlowTTM"`
	EVToFreeCashFlowTTM                       string `json:"evToFreeCashFlowTTM"`
	EVToEBITDATTM                             string `json:"evToEBITDATTM"`
	NetDebtToEBITDATTM                        string `json:"netDebtToEBITDATTM"`
	CurrentRatioTTM                           string `json:"currentRatioTTM"`
	IncomeQualityTTM                          string `json:"incomeQualityTTM"`
	GrahamNumberTTM                           string `json:"grahamNumberTTM"`
	GrahamNetNetTTM                           string `json:"grahamNetNetTTM"`
	TaxBurdenTTM                              string `json:"taxBurdenTTM"`
	InterestBurdenTTM                         string `json:"interestBurdenTTM"`
	WorkingCapitalTTM                         string `json:"workingCapitalTTM"`
	InvestedCapitalTTM                        string `json:"investedCapitalTTM"`
	ReturnOnAssetsTTM                         string `json:"returnOnAssetsTTM"`
	OperatingReturnOnAssetsTTM                string `json:"operatingReturnOnAssetsTTM"`
	ReturnOnTangibleAssetsTTM                 string `json:"returnOnTangibleAssetsTTM"`
	ReturnOnEquityTTM                         string `json:"returnOnEquityTTM"`
	ReturnOnInvestedCapitalTTM                string `json:"returnOnInvestedCapitalTTM"`
	ReturnOnCapitalEmployedTTM                string `json:"returnOnCapitalEmployedTTM"`
	EarningsYieldTTM                          string `json:"earningsYieldTTM"`
	FreeCashFlowYieldTTM                      string `json:"freeCashFlowYieldTTM"`
	CapexToOperatingCashFlowTTM               string `json:"capexToOperatingCashFlowTTM"`
	CapexToDepreciationTTM                    string `json:"capexToDepreciationTTM"`
	CapexToRevenueTTM                         string `json:"capexToRevenueTTM"`
	SalesGeneralAndAdministrativeToRevenueTTM string `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopementToRevenueTTM       string `json:"researchAndDevelopementToRevenueTTM"`
	StockBasedCompensationToRevenueTTM        string `json:"stockBasedCompensationToRevenueTTM"`
	IntangiblesToTotalAssetsTTM               string `json:"intangiblesToTotalAssetsTTM"`
	AverageReceivablesTTM                     string `json:"averageReceivablesTTM"`
	AveragePayablesTTM                        string `json:"averagePayablesTTM"`
	AverageInventoryTTM                       string `json:"averageInventoryTTM"`
	DaysOfSalesOutstandingTTM                 string `json:"daysOfSalesOutstandingTTM"`
	DaysOfPayablesOutstandingTTM              string `json:"daysOfPayablesOutstandingTTM"`
	DaysOfInventoryOutstandingTTM             string `json:"daysOfInventoryOutstandingTTM"`
	OperatingCycleTTM                         string `json:"operatingCycleTTM"`
	CashConversionCycleTTM                    string `json:"cashConversionCycleTTM"`
	FreeCashFlowToEquityTTM                   string `json:"freeCashFlowToEquityTTM"`
	FreeCashFlowToFirmTTM                     string `json:"freeCashFlowToFirmTTM"`
	TangibleAssetValueTTM                     string `json:"tangibleAssetValueTTM"`
	NetCurrentAssetValueTTM                   string `json:"netCurrentAssetValueTTM"`
}

// GetKeyMetricsTTMBulk retrieves trailing twelve months data for all companies
func (c *Client) GetKeyMetricsTTMBulk() ([]KeyMetricsTTMBulkResponse, error) {
	// Build the URL
	baseURL := c.BaseURL + "/key-metrics-ttm-bulk"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Add API key if available
	if c.APIKey != "" {
		q := u.Query()
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
	var result []KeyMetricsTTMBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
