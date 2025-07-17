package go_fmp

import (
	"fmt"
)

// BulkKeyMetricsTTMParams represents the parameters for the Bulk Key Metrics TTM API
type BulkKeyMetricsTTMParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkKeyMetricsTTMResponse represents the response from the Bulk Key Metrics TTM API
type BulkKeyMetricsTTMResponse struct {
	Symbol                              string  `json:"symbol"`
	RevenuePerShareTTM                  float64 `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM                float64 `json:"netIncomePerShareTTM"`
	OperatingCashFlowPerShareTTM        float64 `json:"operatingCashFlowPerShareTTM"`
	FreeCashFlowPerShareTTM             float64 `json:"freeCashFlowPerShareTTM"`
	CashPerShareTTM                     float64 `json:"cashPerShareTTM"`
	BookValuePerShareTTM                float64 `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM        float64 `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM       float64 `json:"shareholdersEquityPerShareTTM"`
	InterestDebtPerShareTTM             float64 `json:"interestDebtPerShareTTM"`
	MarketCapTTM                        float64 `json:"marketCapTTM"`
	EnterpriseValueTTM                  float64 `json:"enterpriseValueTTM"`
	PeRatioTTM                          float64 `json:"peRatioTTM"`
	PriceToSalesRatioTTM                float64 `json:"priceToSalesRatioTTM"`
	PocfratioTTM                        float64 `json:"pocfratioTTM"`
	PfcfRatioTTM                        float64 `json:"pfcfRatioTTM"`
	PbRatioTTM                          float64 `json:"pbRatioTTM"`
	PtbRatioTTM                         float64 `json:"ptbRatioTTM"`
	EvToSalesTTM                        float64 `json:"evToSalesTTM"`
	EnterpriseValueOverEBITDATTM        float64 `json:"enterpriseValueOverEBITDATTM"`
	EvToOperatingCashFlowTTM            float64 `json:"evToOperatingCashFlowTTM"`
	EvToFreeCashFlowTTM                 float64 `json:"evToFreeCashFlowTTM"`
	EarningsYieldTTM                    float64 `json:"earningsYieldTTM"`
	FreeCashFlowYieldTTM                float64 `json:"freeCashFlowYieldTTM"`
	DebtToEquityTTM                     float64 `json:"debtToEquityTTM"`
	DebtToAssetsTTM                     float64 `json:"debtToAssetsTTM"`
	NetDebtToEBITDATTM                  float64 `json:"netDebtToEBITDATTM"`
	CurrentRatioTTM                     float64 `json:"currentRatioTTM"`
	InterestCoverageTTM                 float64 `json:"interestCoverageTTM"`
	IncomeQualityTTM                    float64 `json:"incomeQualityTTM"`
	DividendYieldTTM                    float64 `json:"dividendYieldTTM"`
	DividendYieldPercentageTTM          float64 `json:"dividendYieldPercentageTTM"`
	PayoutRatioTTM                      float64 `json:"payoutRatioTTM"`
	SalesGeneralAndAdministrativeToRevenueTTM float64 `json:"salesGeneralAndAdministrativeToRevenueTTM"`
	ResearchAndDevelopmentToRevenueTTM  float64 `json:"researchAndDevelopmentToRevenueTTM"`
	IntangiblesToTotalAssetsTTM         float64 `json:"intangiblesToTotalAssetsTTM"`
	CapexToOperatingCashFlowTTM         float64 `json:"capexToOperatingCashFlowTTM"`
	CapexToRevenueTTM                   float64 `json:"capexToRevenueTTM"`
	CapexToDepreciationTTM              float64 `json:"capexToDepreciationTTM"`
	StockBasedCompensationToRevenueTTM  float64 `json:"stockBasedCompensationToRevenueTTM"`
	GrahamNumberTTM                     float64 `json:"grahamNumberTTM"`
	RoicTTM                             float64 `json:"roicTTM"`
	ReturnOnTangibleAssetsTTM           float64 `json:"returnOnTangibleAssetsTTM"`
	GrahamNetNetTTM                     float64 `json:"grahamNetNetTTM"`
	WorkingCapitalTTM                   float64 `json:"workingCapitalTTM"`
	TangibleAssetValueTTM               float64 `json:"tangibleAssetValueTTM"`
	NetCurrentAssetValueTTM             float64 `json:"netCurrentAssetValueTTM"`
	InvestedCapitalTTM                  float64 `json:"investedCapitalTTM"`
	AverageReceivablesTTM               float64 `json:"averageReceivablesTTM"`
	AveragePayablesTTM                  float64 `json:"averagePayablesTTM"`
	AverageInventoryTTM                 float64 `json:"averageInventoryTTM"`
	DaysSalesOutstandingTTM             float64 `json:"daysSalesOutstandingTTM"`
	DaysPayablesOutstandingTTM          float64 `json:"daysPayablesOutstandingTTM"`
	DaysOfInventoryOnHandTTM            float64 `json:"daysOfInventoryOnHandTTM"`
	ReceivablesTurnoverTTM              float64 `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                 float64 `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                float64 `json:"inventoryTurnoverTTM"`
	RoeTTM                              float64 `json:"roeTTM"`
	CapexPerShareTTM                    float64 `json:"capexPerShareTTM"`
	DividendPerShareTTM                 float64 `json:"dividendPerShareTTM"`
	DebtToMarketCapTTM                  float64 `json:"debtToMarketCapTTM"`
}

// BulkKeyMetricsTTM retrieves trailing twelve months key metrics for multiple stocks
func (c *Client) BulkKeyMetricsTTM(params BulkKeyMetricsTTMParams) ([]BulkKeyMetricsTTMResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkKeyMetricsTTMResponse
	err := c.doRequest(c.BaseURL+"/bulk-key-metrics-ttm", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk key metrics TTM: %w", err)
	}

	return result, nil
}
