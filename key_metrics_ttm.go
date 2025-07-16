package go_fmp

import (
	"fmt"
)

// KeyMetricsTTMParams represents the parameters for the Key Metrics TTM API
type KeyMetricsTTMParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// KeyMetricsTTMResponse represents the response from the Key Metrics TTM API
type KeyMetricsTTMResponse struct {
	Symbol                                    string  `json:"symbol"`
	MarketCap                                 int64   `json:"marketCap"`
	EnterpriseValueTTM                        int64   `json:"enterpriseValueTTM"`
	EvToSalesTTM                              float64 `json:"evToSalesTTM"`
	EvToOperatingCashFlowTTM                  float64 `json:"evToOperatingCashFlowTTM"`
	EvToFreeCashFlowTTM                       float64 `json:"evToFreeCashFlowTTM"`
	EvToEBITDATTM                             float64 `json:"evToEBITDATTM"`
	NetDebtToEBITDATTM                        float64 `json:"netDebtToEBITDATTM"`
	CurrentRatioTTM                           float64 `json:"currentRatioTTM"`
	IncomeQualityTTM                          float64 `json:"incomeQualityTTM"`
	GrahamNumberTTM                           float64 `json:"grahamNumberTTM"`
	GrahamNetNetTTM                           float64 `json:"grahamNetNetTTM"`
	TaxBurdenTTM                              float64 `json:"taxBurdenTTM"`
	InterestBurdenTTM                         float64 `json:"interestBurdenTTM"`
	WorkingCapitalTTM                         int64   `json:"workingCapitalTTM"`
	InvestedCapitalTTM                        int64   `json:"investedCapitalTTM"`
	ReturnOnAssetsTTM                         float64 `json:"returnOnAssetsTTM"`
	OperatingReturnOnAssetsTTM                float64 `json:"operatingReturnOnAssetsTTM"`
	ReturnOnTangibleAssetsTTM                 float64 `json:"returnOnTangibleAssetsTTM"`
	ReturnOnEquityTTM                         float64 `json:"returnOnEquityTTM"`
	ReturnOnInvestedCapitalTTM                float64 `json:"returnOnInvestedCapitalTTM"`
	ReturnOnCapitalEmployedTTM                float64 `json:"returnOnCapitalEmployedTTM"`
	EarningsYieldTTM                          float64 `json:"earningsYieldTTM"`
	FreeCashFlowYieldTTM                      float64 `json:"freeCashFlowYieldTTM"`
	CapexToOperatingCashFlowTTM               float64 `json:"capexToOperatingCashFlowTTM"`
	CapexToDepreciationTTM                    float64 `json:"capexToDepreciationTTM"`
	CapexToRevenueTTM                         float64 `json:"capexToRevenueTTM"`
	SalesGeneralAndAdministrativeToRevenueTTM float64 `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopementToRevenueTTM       float64 `json:"researchAndDevelopementToRevenueTTM"`
	StockBasedCompensationToRevenueTTM        float64 `json:"stockBasedCompensationToRevenueTTM"`
}

// KeyMetricsTTM retrieves trailing twelve months key financial metrics for a specific stock symbol
func (c *Client) KeyMetricsTTM(params KeyMetricsTTMParams) ([]KeyMetricsTTMResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []KeyMetricsTTMResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/key-metrics-ttm", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
