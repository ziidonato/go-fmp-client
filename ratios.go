package go_fmp

import (
	"fmt"
	"time"
)

// RatiosParams represents the parameters for the Ratios API
type RatiosParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	Period *string `json:"period"` // Optional: Period type ("annual" or "quarter")
	Limit  *int    `json:"limit"`  // Optional: Number of records to return
}

// RatiosResponse represents the response from the Ratios API
type RatiosResponse struct {
	Symbol                  string    `json:"symbol"`
	Date                    time.Time `json:"date"`
	Period                  string    `json:"period"`
	CurrentRatio            float64   `json:"currentRatio"`
	QuickRatio              float64   `json:"quickRatio"`
	CashRatio               float64   `json:"cashRatio"`
	DaysOfSalesOutstanding  float64   `json:"daysOfSalesOutstanding"`
	DaysOfInventoryOutstanding float64 `json:"daysOfInventoryOutstanding"`
	OperatingCycle          float64   `json:"operatingCycle"`
	DaysOfPayablesOutstanding float64 `json:"daysOfPayablesOutstanding"`
	CashConversionCycle     float64   `json:"cashConversionCycle"`
	GrossProfitMargin       float64   `json:"grossProfitMargin"`
	OperatingProfitMargin   float64   `json:"operatingProfitMargin"`
	PretaxProfitMargin      float64   `json:"pretaxProfitMargin"`
	NetProfitMargin         float64   `json:"netProfitMargin"`
	EffectiveTaxRate        float64   `json:"effectiveTaxRate"`
	ReturnOnAssets          float64   `json:"returnOnAssets"`
	ReturnOnEquity          float64   `json:"returnOnEquity"`
	ReturnOnCapitalEmployed float64   `json:"returnOnCapitalEmployed"`
	NetIncomePerEBT         float64   `json:"netIncomePerEBT"`
	EBTPerEBIT              float64   `json:"ebtPerEbit"`
	EBITPerRevenue          float64   `json:"ebitPerRevenue"`
	DebtRatio               float64   `json:"debtRatio"`
	DebtEquityRatio         float64   `json:"debtEquityRatio"`
	LongTermDebtToCapitalization float64 `json:"longTermDebtToCapitalization"`
	TotalDebtToCapitalization float64 `json:"totalDebtToCapitalization"`
	InterestCoverage        float64   `json:"interestCoverage"`
	CashFlowToDebtRatio     float64   `json:"cashFlowToDebtRatio"`
	CompanyEquityMultiplier float64   `json:"companyEquityMultiplier"`
	ReceivablesTurnover     float64   `json:"receivablesTurnover"`
	PayablesTurnover        float64   `json:"payablesTurnover"`
	InventoryTurnover       float64   `json:"inventoryTurnover"`
	FixedAssetTurnover      float64   `json:"fixedAssetTurnover"`
	AssetTurnover           float64   `json:"assetTurnover"`
	OperatingCashFlowPerShare float64 `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare    float64   `json:"freeCashFlowPerShare"`
	CashPerShare            float64   `json:"cashPerShare"`
	PayoutRatio             float64   `json:"payoutRatio"`
	OperatingCashFlowSalesRatio float64 `json:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio float64 `json:"freeCashFlowOperatingCashFlowRatio"`
	CashFlowCoverageRatios  float64   `json:"cashFlowCoverageRatios"`
	ShortTermCoverageRatios float64   `json:"shortTermCoverageRatios"`
	CapitalExpenditureCoverageRatio float64 `json:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio float64 `json:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio     float64   `json:"dividendPayoutRatio"`
	PriceBookValueRatio     float64   `json:"priceBookValueRatio"`
	PriceToBookRatio        float64   `json:"priceToBookRatio"`
	PriceToSalesRatio       float64   `json:"priceToSalesRatio"`
	PriceEarningsRatio      float64   `json:"priceEarningsRatio"`
	PriceToFreeCashFlowsRatio float64 `json:"priceToFreeCashFlowsRatio"`
	PriceToOperatingCashFlowsRatio float64 `json:"priceToOperatingCashFlowsRatio"`
	PriceCashFlowRatio      float64   `json:"priceCashFlowRatio"`
	PriceEarningsToGrowthRatio float64 `json:"priceEarningsToGrowthRatio"`
	PriceSalesRatio         float64   `json:"priceSalesRatio"`
	DividendYield           float64   `json:"dividendYield"`
	EnterpriseValueMultiple float64   `json:"enterpriseValueMultiple"`
	PriceFairValue          float64   `json:"priceFairValue"`
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

	if params.Period != nil {
		urlParams["period"] = *params.Period
	}

	var result []RatiosResponse
	err := c.doRequest(c.BaseURL+"/ratios", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
