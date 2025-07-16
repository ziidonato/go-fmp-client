package go_fmp

import (
	"encoding/json"
	"fmt"
)

// RatiosParams represents the parameters for the Financial Ratios API
type RatiosParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Limit  *int   `json:"limit"`  // Optional: Number of results (Maximum 1000 records per request)
	Period string `json:"period"` // Optional: Period type - "Q1,Q2,Q3,Q4,FY,annual,quarter"
}

// RatiosResponse represents the response from the Financial Ratios API
type RatiosResponse struct {
	Symbol                                  string  `json:"symbol"`
	Date                                    string  `json:"date"`
	FiscalYear                              string  `json:"fiscalYear"`
	Period                                  string  `json:"period"`
	ReportedCurrency                        string  `json:"reportedCurrency"`
	GrossProfitMargin                       float64 `json:"grossProfitMargin"`
	EbitMargin                              float64 `json:"ebitMargin"`
	EbitdaMargin                            float64 `json:"ebitdaMargin"`
	OperatingProfitMargin                   float64 `json:"operatingProfitMargin"`
	PretaxProfitMargin                      float64 `json:"pretaxProfitMargin"`
	ContinuousOperationsProfitMargin        float64 `json:"continuousOperationsProfitMargin"`
	NetProfitMargin                         float64 `json:"netProfitMargin"`
	BottomLineProfitMargin                  float64 `json:"bottomLineProfitMargin"`
	ReceivablesTurnover                     float64 `json:"receivablesTurnover"`
	PayablesTurnover                        float64 `json:"payablesTurnover"`
	InventoryTurnover                       float64 `json:"inventoryTurnover"`
	FixedAssetTurnover                      float64 `json:"fixedAssetTurnover"`
	AssetTurnover                           float64 `json:"assetTurnover"`
	CurrentRatio                            float64 `json:"currentRatio"`
	QuickRatio                              float64 `json:"quickRatio"`
	SolvencyRatio                           float64 `json:"solvencyRatio"`
	CashRatio                               float64 `json:"cashRatio"`
	PriceToEarningsRatio                    float64 `json:"priceToEarningsRatio"`
	PriceToEarningsGrowthRatio              float64 `json:"priceToEarningsGrowthRatio"`
	ForwardPriceToEarningsGrowthRatio       float64 `json:"forwardPriceToEarningsGrowthRatio"`
	PriceToBookRatio                        float64 `json:"priceToBookRatio"`
	PriceToSalesRatio                       float64 `json:"priceToSalesRatio"`
	PriceToFreeCashFlowRatio                float64 `json:"priceToFreeCashFlowRatio"`
	PriceToOperatingCashFlowRatio           float64 `json:"priceToOperatingCashFlowRatio"`
	DebtToAssetsRatio                       float64 `json:"debtToAssetsRatio"`
	DebtToEquityRatio                       float64 `json:"debtToEquityRatio"`
	DebtToCapitalRatio                      float64 `json:"debtToCapitalRatio"`
	LongTermDebtToCapitalRatio              float64 `json:"longTermDebtToCapitalRatio"`
	FinancialLeverageRatio                  float64 `json:"financialLeverageRatio"`
	WorkingCapitalTurnoverRatio             float64 `json:"workingCapitalTurnoverRatio"`
	OperatingCashFlowRatio                  float64 `json:"operatingCashFlowRatio"`
	OperatingCashFlowSalesRatio             float64 `json:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio      float64 `json:"freeCashFlowOperatingCashFlowRatio"`
	DebtServiceCoverageRatio                float64 `json:"debtServiceCoverageRatio"`
	InterestCoverageRatio                   float64 `json:"interestCoverageRatio"`
	ShortTermOperatingCashFlowCoverageRatio float64 `json:"shortTermOperatingCashFlowCoverageRatio"`
	OperatingCashFlowCoverageRatio          float64 `json:"operatingCashFlowCoverageRatio"`
	CapitalExpenditureCoverageRatio         float64 `json:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio       float64 `json:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio                     float64 `json:"dividendPayoutRatio"`
	DividendYield                           float64 `json:"dividendYield"`
	DividendYieldPercentage                 float64 `json:"dividendYieldPercentage"`
	RevenuePerShare                         float64 `json:"revenuePerShare"`
	NetIncomePerShare                       float64 `json:"netIncomePerShare"`
	InterestDebtPerShare                    float64 `json:"interestDebtPerShare"`
	CashPerShare                            float64 `json:"cashPerShare"`
	BookValuePerShare                       float64 `json:"bookValuePerShare"`
	TangibleBookValuePerShare               float64 `json:"tangibleBookValuePerShare"`
	ShareholdersEquityPerShare              float64 `json:"shareholdersEquityPerShare"`
	OperatingCashFlowPerShare               float64 `json:"operatingCashFlowPerShare"`
	CapexPerShare                           float64 `json:"capexPerShare"`
	FreeCashFlowPerShare                    float64 `json:"freeCashFlowPerShare"`
	NetIncomePerEBT                         float64 `json:"netIncomePerEBT"`
	EbtPerEbit                              float64 `json:"ebtPerEbit"`
	PriceToFairValue                        float64 `json:"priceToFairValue"`
	DebtToMarketCap                         float64 `json:"debtToMarketCap"`
	EffectiveTaxRate                        float64 `json:"effectiveTaxRate"`
	EnterpriseValueMultiple                 float64 `json:"enterpriseValueMultiple"`
}

// Ratios retrieves financial ratios for a specific stock symbol
func (c *Client) Ratios(params RatiosParams) ([]RatiosResponse, error) {
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
		urlParams["period"] = params.Period
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/ratios", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []RatiosResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
