package go_fmp

import (
	"fmt"
)

// BulkRatiosTTMParams represents the parameters for the Bulk Ratios TTM API
type BulkRatiosTTMParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkRatiosTTMResponse represents the response from the Bulk Ratios TTM API
type BulkRatiosTTMResponse struct {
	Symbol                              string  `json:"symbol"`
	DividendYielTTM                     float64 `json:"dividendYielTTM"`
	DividendYielPercentageTTM           float64 `json:"dividendYielPercentageTTM"`
	PeRatioTTM                          float64 `json:"peRatioTTM"`
	PegRatioTTM                         float64 `json:"pegRatioTTM"`
	PayoutRatioTTM                      float64 `json:"payoutRatioTTM"`
	CurrentRatioTTM                     float64 `json:"currentRatioTTM"`
	QuickRatioTTM                       float64 `json:"quickRatioTTM"`
	CashRatioTTM                        float64 `json:"cashRatioTTM"`
	DaysOfSalesOutstandingTTM           float64 `json:"daysOfSalesOutstandingTTM"`
	DaysOfInventoryOutstandingTTM       float64 `json:"daysOfInventoryOutstandingTTM"`
	OperatingCycleTTM                   float64 `json:"operatingCycleTTM"`
	DaysOfPayablesOutstandingTTM        float64 `json:"daysOfPayablesOutstandingTTM"`
	CashConversionCycleTTM              float64 `json:"cashConversionCycleTTM"`
	GrossProfitMarginTTM                float64 `json:"grossProfitMarginTTM"`
	OperatingProfitMarginTTM            float64 `json:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM               float64 `json:"pretaxProfitMarginTTM"`
	NetProfitMarginTTM                  float64 `json:"netProfitMarginTTM"`
	EffectiveTaxRateTTM                 float64 `json:"effectiveTaxRateTTM"`
	ReturnOnAssetsTTM                   float64 `json:"returnOnAssetsTTM"`
	ReturnOnEquityTTM                   float64 `json:"returnOnEquityTTM"`
	ReturnOnCapitalEmployedTTM          float64 `json:"returnOnCapitalEmployedTTM"`
	NetIncomePerEBTTTM                  float64 `json:"netIncomePerEBTTTM"`
	EbtPerEbitTTM                       float64 `json:"ebtPerEbitTTM"`
	EbitPerRevenueTTM                   float64 `json:"ebitPerRevenueTTM"`
	DebtRatioTTM                        float64 `json:"debtRatioTTM"`
	DebtEquityRatioTTM                  float64 `json:"debtEquityRatioTTM"`
	LongTermDebtToCapitalizationTTM     float64 `json:"longTermDebtToCapitalizationTTM"`
	TotalDebtToCapitalizationTTM        float64 `json:"totalDebtToCapitalizationTTM"`
	InterestCoverageTTM                 float64 `json:"interestCoverageTTM"`
	CashFlowToDebtRatioTTM              float64 `json:"cashFlowToDebtRatioTTM"`
	CompanyEquityMultiplierTTM          float64 `json:"companyEquityMultiplierTTM"`
	ReceivablesTurnoverTTM              float64 `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                 float64 `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                float64 `json:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM               float64 `json:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM                    float64 `json:"assetTurnoverTTM"`
	OperatingCashFlowPerShareTTM        float64 `json:"operatingCashFlowPerShareTTM"`
	FreeCashFlowPerShareTTM             float64 `json:"freeCashFlowPerShareTTM"`
	CashPerShareTTM                     float64 `json:"cashPerShareTTM"`
	OperatingCashFlowSalesRatioTTM      float64 `json:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM float64 `json:"freeCashFlowOperatingCashFlowRatioTTM"`
	CashFlowCoverageRatiosTTM           float64 `json:"cashFlowCoverageRatiosTTM"`
	ShortTermCoverageRatiosTTM          float64 `json:"shortTermCoverageRatiosTTM"`
	CapitalExpenditureCoverageRatioTTM  float64 `json:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTM float64 `json:"dividendPaidAndCapexCoverageRatioTTM"`
	PriceBookValueRatioTTM              float64 `json:"priceBookValueRatioTTM"`
	PriceToBookRatioTTM                 float64 `json:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                float64 `json:"priceToSalesRatioTTM"`
	PriceEarningsRatioTTM               float64 `json:"priceEarningsRatioTTM"`
	PriceToFreeCashFlowsRatioTTM        float64 `json:"priceToFreeCashFlowsRatioTTM"`
	PriceToOperatingCashFlowsRatioTTM   float64 `json:"priceToOperatingCashFlowsRatioTTM"`
	PriceCashFlowRatioTTM               float64 `json:"priceCashFlowRatioTTM"`
	PriceEarningsToGrowthRatioTTM       float64 `json:"priceEarningsToGrowthRatioTTM"`
	PriceSalesRatioTTM                  float64 `json:"priceSalesRatioTTM"`
	EnterpriseValueMultipleTTM          float64 `json:"enterpriseValueMultipleTTM"`
	PriceFairValueTTM                   float64 `json:"priceFairValueTTM"`
	DividendPerShareTTM                 float64 `json:"dividendPerShareTTM"`
}

// BulkRatiosTTM retrieves trailing twelve months financial ratios for multiple stocks
func (c *Client) BulkRatiosTTM(params BulkRatiosTTMParams) ([]BulkRatiosTTMResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkRatiosTTMResponse
	err := c.doRequest(c.BaseURL+"/bulk-ratios-ttm", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk ratios TTM: %w", err)
	}

	return result, nil
}
