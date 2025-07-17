package go_fmp

import (
	"fmt"
	"time"
)

// KeyMetricsParams represents the parameters for the Key Metrics API
type KeyMetricsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
	Period Period `json:"period"` // Optional: Period type - "Q1,Q2,Q3,Q4,FY,annual,quarter"
}

// KeyMetricsResponse represents the response from the Key Metrics API
type KeyMetricsResponse struct {
	Symbol                                 string           `json:"symbol"`
	Date                                   time.Time        `json:"date"`
	FiscalYear                             string           `json:"fiscalYear"`
	Period                                 Period           `json:"period"`
	ReportedCurrency                       ReportedCurrency `json:"reportedCurrency"`
	MarketCap                              int64   `json:"marketCap"`
	EnterpriseValue                        int64   `json:"enterpriseValue"`
	EvToSales                              float64 `json:"evToSales"`
	EvToOperatingCashFlow                  float64 `json:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64 `json:"evToFreeCashFlow"`
	EvToEBITDA                             float64 `json:"evToEBITDA"`
	NetDebtToEBITDA                        float64 `json:"netDebtToEBITDA"`
	CurrentRatio                           float64 `json:"currentRatio"`
	IncomeQuality                          float64 `json:"incomeQuality"`
	GrahamNumber                           float64 `json:"grahamNumber"`
	GrahamNetNet                           float64 `json:"grahamNetNet"`
	TaxBurden                              float64 `json:"taxBurden"`
	InterestBurden                         float64 `json:"interestBurden"`
	WorkingCapital                         int64   `json:"workingCapital"`
	InvestedCapital                        int64   `json:"investedCapital"`
	ReturnOnAssets                         float64 `json:"returnOnAssets"`
	OperatingReturnOnAssets                float64 `json:"operatingReturnOnAssets"`
	ReturnOnTangibleAssets                 float64 `json:"returnOnTangibleAssets"`
	ReturnOnEquity                         float64 `json:"returnOnEquity"`
	ReturnOnInvestedCapital                float64 `json:"returnOnInvestedCapital"`
	ReturnOnCapitalEmployed                float64 `json:"returnOnCapitalEmployed"`
	EarningsYield                          float64 `json:"earningsYield"`
	FreeCashFlowYield                      float64 `json:"freeCashFlowYield"`
	CapexToOperatingCashFlow               float64 `json:"capexToOperatingCashFlow"`
	CapexToDepreciation                    float64 `json:"capexToDepreciation"`
	CapexToRevenue                         float64 `json:"capexToRevenue"`
	SalesGeneralAndAdministrativeToRevenue float64 `json:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDevelopementToRevenue       float64 `json:"researchAndDevelopementToRevenue"`
	StockBasedCompensationToRevenue        float64 `json:"stockBasedCompensationToRevenue"`
	IntangiblesToTotalAssets               float64 `json:"intangiblesToTotalAssets"`
	AverageReceivables                     int64   `json:"averageReceivables"`
	AveragePayables                        int64   `json:"averagePayables"`
	AverageInventory                       int64   `json:"averageInventory"`
	DaysOfSalesOutstanding                 float64 `json:"daysOfSalesOutstanding"`
	DaysOfPayablesOutstanding              float64 `json:"daysOfPayablesOutstanding"`
	DaysOfInventoryOutstanding             float64 `json:"daysOfInventoryOutstanding"`
	OperatingCycle                         float64 `json:"operatingCycle"`
	CashConversionCycle                    float64 `json:"cashConversionCycle"`
	FreeCashFlowToEquity                   int64   `json:"freeCashFlowToEquity"`
	FreeCashFlowToFirm                     float64 `json:"freeCashFlowToFirm"`
	TangibleAssetValue                     int64   `json:"tangibleAssetValue"`
	NetCurrentAssetValue                   int64   `json:"netCurrentAssetValue"`
}

// KeyMetrics retrieves key financial metrics for a specific stock symbol
func (c *Client) KeyMetrics(params KeyMetricsParams) ([]KeyMetricsResponse, error) {
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
		urlParams["period"] = string(params.Period)
	}

	var result []KeyMetricsResponse
	err := c.doRequest(c.BaseURL+"/key-metrics", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
