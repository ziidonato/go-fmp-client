package go_fmp

import (
	"fmt"
)

// RatiosTTMParams represents the parameters for the Financial Ratios TTM API
type RatiosTTMParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// RatiosTTMResponse represents the response from the Financial Ratios TTM API
type RatiosTTMResponse struct {
	Symbol                                     string  `json:"symbol"`
	GrossProfitMarginTTM                       float64 `json:"grossProfitMarginTTM"`
	EbitMarginTTM                              float64 `json:"ebitMarginTTM"`
	EbitdaMarginTTM                            float64 `json:"ebitdaMarginTTM"`
	OperatingProfitMarginTTM                   float64 `json:"operatingProfitMarginTTM"`
	PretaxProfitMarginTTM                      float64 `json:"pretaxProfitMarginTTM"`
	ContinuousOperationsProfitMarginTTM        float64 `json:"continuousOperationsProfitMarginTTM"`
	NetProfitMarginTTM                         float64 `json:"netProfitMarginTTM"`
	BottomLineProfitMarginTTM                  float64 `json:"bottomLineProfitMarginTTM"`
	ReceivablesTurnoverTTM                     float64 `json:"receivablesTurnoverTTM"`
	PayablesTurnoverTTM                        float64 `json:"payablesTurnoverTTM"`
	InventoryTurnoverTTM                       float64 `json:"inventoryTurnoverTTM"`
	FixedAssetTurnoverTTM                      float64 `json:"fixedAssetTurnoverTTM"`
	AssetTurnoverTTM                           float64 `json:"assetTurnoverTTM"`
	CurrentRatioTTM                            float64 `json:"currentRatioTTM"`
	QuickRatioTTM                              float64 `json:"quickRatioTTM"`
	SolvencyRatioTTM                           float64 `json:"solvencyRatioTTM"`
	CashRatioTTM                               float64 `json:"cashRatioTTM"`
	PriceToEarningsRatioTTM                    float64 `json:"priceToEarningsRatioTTM"`
	PriceToEarningsGrowthRatioTTM              float64 `json:"priceToEarningsGrowthRatioTTM"`
	ForwardPriceToEarningsGrowthRatioTTM       float64 `json:"forwardPriceToEarningsGrowthRatioTTM"`
	PriceToBookRatioTTM                        float64 `json:"priceToBookRatioTTM"`
	PriceToSalesRatioTTM                       float64 `json:"priceToSalesRatioTTM"`
	PriceToFreeCashFlowRatioTTM                float64 `json:"priceToFreeCashFlowRatioTTM"`
	PriceToOperatingCashFlowRatioTTM           float64 `json:"priceToOperatingCashFlowRatioTTM"`
	DebtToAssetsRatioTTM                       float64 `json:"debtToAssetsRatioTTM"`
	DebtToEquityRatioTTM                       float64 `json:"debtToEquityRatioTTM"`
	DebtToCapitalRatioTTM                      float64 `json:"debtToCapitalRatioTTM"`
	LongTermDebtToCapitalRatioTTM              float64 `json:"longTermDebtToCapitalRatioTTM"`
	FinancialLeverageRatioTTM                  float64 `json:"financialLeverageRatioTTM"`
	WorkingCapitalTurnoverRatioTTM             float64 `json:"workingCapitalTurnoverRatioTTM"`
	OperatingCashFlowRatioTTM                  float64 `json:"operatingCashFlowRatioTTM"`
	OperatingCashFlowSalesRatioTTM             float64 `json:"operatingCashFlowSalesRatioTTM"`
	FreeCashFlowOperatingCashFlowRatioTTM      float64 `json:"freeCashFlowOperatingCashFlowRatioTTM"`
	DebtServiceCoverageRatioTTM                float64 `json:"debtServiceCoverageRatioTTM"`
	InterestCoverageRatioTTM                   float64 `json:"interestCoverageRatioTTM"`
	ShortTermOperatingCashFlowCoverageRatioTTM float64 `json:"shortTermOperatingCashFlowCoverageRatioTTM"`
	OperatingCashFlowCoverageRatioTTM          float64 `json:"operatingCashFlowCoverageRatioTTM"`
	CapitalExpenditureCoverageRatioTTM         float64 `json:"capitalExpenditureCoverageRatioTTM"`
	DividendPaidAndCapexCoverageRatioTTM       float64 `json:"dividendPaidAndCapexCoverageRatioTTM"`
	DividendPayoutRatioTTM                     float64 `json:"dividendPayoutRatioTTM"`
	DividendYieldTTM                           float64 `json:"dividendYieldTTM"`
	DividendYieldPercentageTTM                 float64 `json:"dividendYieldPercentageTTM"`
	RevenuePerShareTTM                         float64 `json:"revenuePerShareTTM"`
	NetIncomePerShareTTM                       float64 `json:"netIncomePerShareTTM"`
	InterestDebtPerShareTTM                    float64 `json:"interestDebtPerShareTTM"`
	CashPerShareTTM                            float64 `json:"cashPerShareTTM"`
	BookValuePerShareTTM                       float64 `json:"bookValuePerShareTTM"`
	TangibleBookValuePerShareTTM               float64 `json:"tangibleBookValuePerShareTTM"`
	ShareholdersEquityPerShareTTM              float64 `json:"shareholdersEquityPerShareTTM"`
	OperatingCashFlowPerShareTTM               float64 `json:"operatingCashFlowPerShareTTM"`
	CapexPerShareTTM                           float64 `json:"capexPerShareTTM"`
	FreeCashFlowPerShareTTM                    float64 `json:"freeCashFlowPerShareTTM"`
	NetIncomePerEBTTTM                         float64 `json:"netIncomePerEBTTTM"`
	EbtPerEbitTTM                              float64 `json:"ebtPerEbitTTM"`
	PriceToFairValueTTM                        float64 `json:"priceToFairValueTTM"`
	DebtToMarketCapTTM                         float64 `json:"debtToMarketCapTTM"`
	EffectiveTaxRateTTM                        float64 `json:"effectiveTaxRateTTM"`
	EnterpriseValueMultipleTTM                 float64 `json:"enterpriseValueMultipleTTM"`
}

// RatiosTTM retrieves trailing twelve months financial ratios for a specific stock symbol
func (c *Client) RatiosTTM(params RatiosTTMParams) ([]RatiosTTMResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []RatiosTTMResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/ratios-ttm", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, err
}
