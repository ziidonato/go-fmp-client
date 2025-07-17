package go_fmp

import (
	"fmt"
	"time"
)

// KeyMetricsParams represents the parameters for the Key Metrics API
type KeyMetricsParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return (default: 40)
}

// KeyMetricsResponse represents the response from the Key Metrics API
type KeyMetricsResponse struct {
	Symbol                                 string  `json:"symbol"`
	Date                                   time.Time  `json:"date"`
	CalendarYear                           string  `json:"calendarYear"`
	Period                                 string  `json:"period"`
	RevenuePerShare                        float64 `json:"revenuePerShare"`
	NetIncomePerShare                      float64 `json:"netIncomePerShare"`
	OperatingCashFlowPerShare              float64 `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare                   float64 `json:"freeCashFlowPerShare"`
	CashPerShare                           float64 `json:"cashPerShare"`
	BookValuePerShare                      float64 `json:"bookValuePerShare"`
	TangibleBookValuePerShare              float64 `json:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare             float64 `json:"shareholdersEquityPerShare"`
	InterestDebtPerShare                   float64 `json:"interestDebtPerShare"`
	MarketCap                              float64 `json:"marketCap"`
	EnterpriseValue                        float64 `json:"enterpriseValue"`
	PeRatio                                float64 `json:"peRatio"`
	PriceToSalesRatio                      float64 `json:"priceToSalesRatio"`
	PocfRatio                              float64 `json:"pocfratio"`
	PfcfRatio                              float64 `json:"pfcfRatio"`
	PbRatio                                float64 `json:"pbRatio"`
	PtbRatio                               float64 `json:"ptbRatio"`
	EvToSales                              float64 `json:"evToSales"`
	EnterpriseValueOverEBITDA              float64 `json:"enterpriseValueOverEBITDA"`
	EvToOperatingCashFlow                  float64 `json:"evToOperatingCashFlow"`
	EvToFreeCashFlow                       float64 `json:"evToFreeCashFlow"`
	EarningsYield                          float64 `json:"earningsYield"`
	FreeCashFlowYield                      float64 `json:"freeCashFlowYield"`
	DebtToEquity                           float64 `json:"debtToEquity"`
	DebtToAssets                           float64 `json:"debtToAssets"`
	NetDebtToEBITDA                        float64 `json:"netDebtToEBITDA"`
	CurrentRatio                           float64 `json:"currentRatio"`
	InterestCoverage                       float64 `json:"interestCoverage"`
	IncomeQuality                          float64 `json:"incomeQuality"`
	DividendYield                          float64 `json:"dividendYield"`
	PayoutRatio                            float64 `json:"payoutRatio"`
	SalesGeneralAndAdministrativeToRevenue float64 `json:"salesGeneralAndAdministrativeToRevenue"`
	ResearchAndDevelopmentToRevenue        float64 `json:"researchAndDdevelopementToRevenue"`
	IntangiblesToTotalAssets               float64 `json:"intangiblesToTotalAssets"`
	CapexToOperatingCashFlow               float64 `json:"capexToOperatingCashFlow"`
	CapexToRevenue                         float64 `json:"capexToRevenue"`
	CapexToDepreciation                    float64 `json:"capexToDepreciation"`
	StockBasedCompensationToRevenue        float64 `json:"stockBasedCompensationToRevenue"`
	GrahamNumber                           float64 `json:"grahamNumber"`
	Roic                                   float64 `json:"roic"`
	ReturnOnTangibleAssets                 float64 `json:"returnOnTangibleAssets"`
	GrahamNetNet                           float64 `json:"grahamNetNet"`
	WorkingCapital                         float64 `json:"workingCapital"`
	TangibleAssetValue                     float64 `json:"tangibleAssetValue"`
	NetCurrentAssetValue                   float64 `json:"netCurrentAssetValue"`
	InvestedCapital                        float64 `json:"investedCapital"`
	AverageReceivables                     float64 `json:"averageReceivables"`
	AveragePayables                        float64 `json:"averagePayables"`
	AverageInventory                       float64 `json:"averageInventory"`
	DaysSalesOutstanding                   float64 `json:"daysSalesOutstanding"`
	DaysPayablesOutstanding                float64 `json:"daysPayablesOutstanding"`
	DaysOfInventoryOnHand                  float64 `json:"daysOfInventoryOnHand"`
	ReceivablesTurnover                    float64 `json:"receivablesTurnover"`
	PayablesTurnover                       float64 `json:"payablesTurnover"`
	InventoryTurnover                      float64 `json:"inventoryTurnover"`
	Roe                                    float64 `json:"roe"`
	CapexPerShare                          float64 `json:"capexPerShare"`
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

	if params.Period != nil {
		urlParams["period"] = *params.Period
	}

	var result []KeyMetricsResponse
	err := c.doRequest(c.BaseURL+"/key-metrics", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
