package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CustomDCFAdvancedParams represents the parameters for the Custom DCF Advanced API
type CustomDCFAdvancedParams struct {
	Symbol                                     string   `json:"symbol"`                                     // Required: Stock symbol (e.g., "AAPL")
	RevenueGrowthPct                           *float64 `json:"revenueGrowthPct"`                           // Optional: Revenue growth percentage
	EBITDAPct                                  *float64 `json:"ebitdaPct"`                                  // Optional: EBITDA percentage
	DepreciationAndAmortizationPct             *float64 `json:"depreciationAndAmortizationPct"`             // Optional: Depreciation and amortization percentage
	CashAndShortTermInvestmentsPct             *float64 `json:"cashAndShortTermInvestmentsPct"`             // Optional: Cash and short-term investments percentage
	ReceivablesPct                             *float64 `json:"receivablesPct"`                             // Optional: Receivables percentage
	InventoriesPct                             *float64 `json:"inventoriesPct"`                             // Optional: Inventories percentage
	PayablePct                                 *float64 `json:"payablePct"`                                 // Optional: Payable percentage
	EBITPct                                    *float64 `json:"ebitPct"`                                    // Optional: EBIT percentage
	CapitalExpenditurePct                      *float64 `json:"capitalExpenditurePct"`                      // Optional: Capital expenditure percentage
	OperatingCashFlowPct                       *float64 `json:"operatingCashFlowPct"`                       // Optional: Operating cash flow percentage
	SellingGeneralAndAdministrativeExpensesPct *float64 `json:"sellingGeneralAndAdministrativeExpensesPct"` // Optional: SG&A expenses percentage
	TaxRate                                    *float64 `json:"taxRate"`                                    // Optional: Tax rate
	LongTermGrowthRate                         *float64 `json:"longTermGrowthRate"`                         // Optional: Long-term growth rate
	CostOfDebt                                 *float64 `json:"costOfDebt"`                                 // Optional: Cost of debt
	CostOfEquity                               *float64 `json:"costOfEquity"`                               // Optional: Cost of equity
	MarketRiskPremium                          *float64 `json:"marketRiskPremium"`                          // Optional: Market risk premium
	Beta                                       *float64 `json:"beta"`                                       // Optional: Beta
	RiskFreeRate                               *float64 `json:"riskFreeRate"`                               // Optional: Risk-free rate

// CustomDCFAdvancedResponse represents the response from the Custom DCF Advanced API
type CustomDCFAdvancedResponse struct {
	Year                         string  `json:"year"`
	Symbol                       string  `json:"symbol"`
	Revenue                      int64   `json:"revenue"`
	RevenuePercentage            float64 `json:"revenuePercentage"`
	EBITDA                       int64   `json:"ebitda"`
	EBITDAPercentage             float64 `json:"ebitdaPercentage"`
	EBIT                         int64   `json:"ebit"`
	EBITPercentage               float64 `json:"ebitPercentage"`
	Depreciation                 int64   `json:"depreciation"`
	DepreciationPercentage       float64 `json:"depreciationPercentage"`
	TotalCash                    int64   `json:"totalCash"`
	TotalCashPercentage          float64 `json:"totalCashPercentage"`
	Receivables                  int64   `json:"receivables"`
	ReceivablesPercentage        float64 `json:"receivablesPercentage"`
	Inventories                  int64   `json:"inventories"`
	InventoriesPercentage        float64 `json:"inventoriesPercentage"`
	Payable                      int64   `json:"payable"`
	PayablePercentage            float64 `json:"payablePercentage"`
	CapitalExpenditure           int64   `json:"capitalExpenditure"`
	CapitalExpenditurePercentage float64 `json:"capitalExpenditurePercentage"`
	Price                        float64 `json:"price"`
	Beta                         float64 `json:"beta"`
	DilutedSharesOutstanding     int64   `json:"dilutedSharesOutstanding"`
	CostOfDebt                   float64 `json:"costofDebt"`
	TaxRate                      float64 `json:"taxRate"`
	AfterTaxCostOfDebt           float64 `json:"afterTaxCostOfDebt"`
	RiskFreeRate                 float64 `json:"riskFreeRate"`
	MarketRiskPremium            float64 `json:"marketRiskPremium"`
	CostOfEquity                 float64 `json:"costOfEquity"`
	TotalDebt                    int64   `json:"totalDebt"`
	TotalEquity                  int64   `json:"totalEquity"`
	TotalCapital                 int64   `json:"totalCapital"`
	DebtWeighting                float64 `json:"debtWeighting"`
	EquityWeighting              float64 `json:"equityWeighting"`
	WACC                         float64 `json:"wacc"`
	TaxRateCash                  int64   `json:"taxRateCash"`
	EBIAT                        int64   `json:"ebiat"`
	UFCF                         int64   `json:"ufcf"`
	SumPvUfcf                    int64   `json:"sumPvUfcf"`
	LongTermGrowthRate           float64 `json:"longTermGrowthRate"`
	TerminalValue                int64   `json:"terminalValue"`
	PresentTerminalValue         int64   `json:"presentTerminalValue"`
	EnterpriseValue              int64   `json:"enterpriseValue"`
	NetDebt                      int64   `json:"netDebt"`
	EquityValue                  int64   `json:"equityValue"`
	EquityValuePerShare          float64 `json:"equityValuePerShare"`
	FreeCashFlowT1               int64   `json:"freeCashFlowT1"`

// CustomDCFAdvanced runs a tailored Discounted Cash Flow analysis with detailed inputs
func (c *Client) CustomDCFAdvanced(params CustomDCFAdvancedParams) ([]CustomDCFAdvancedResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	// Add optional parameters if provided
	if params.RevenueGrowthPct != nil {
		urlParams["revenueGrowthPct"] = fmt.Sprintf("%f", *params.RevenueGrowthPct)
	}
	if params.EBITDAPct != nil {
		urlParams["ebitdaPct"] = fmt.Sprintf("%f", *params.EBITDAPct)
	}
	if params.DepreciationAndAmortizationPct != nil {
		urlParams["depreciationAndAmortizationPct"] = fmt.Sprintf("%f", *params.DepreciationAndAmortizationPct)
	}
	if params.CashAndShortTermInvestmentsPct != nil {
		urlParams["cashAndShortTermInvestmentsPct"] = fmt.Sprintf("%f", *params.CashAndShortTermInvestmentsPct)
	}
	if params.ReceivablesPct != nil {
		urlParams["receivablesPct"] = fmt.Sprintf("%f", *params.ReceivablesPct)
	}
	if params.InventoriesPct != nil {
		urlParams["inventoriesPct"] = fmt.Sprintf("%f", *params.InventoriesPct)
	}
	if params.PayablePct != nil {
		urlParams["payablePct"] = fmt.Sprintf("%f", *params.PayablePct)
	}
	if params.EBITPct != nil {
		urlParams["ebitPct"] = fmt.Sprintf("%f", *params.EBITPct)
	}
	if params.CapitalExpenditurePct != nil {
		urlParams["capitalExpenditurePct"] = fmt.Sprintf("%f", *params.CapitalExpenditurePct)
	}
	if params.OperatingCashFlowPct != nil {
		urlParams["operatingCashFlowPct"] = fmt.Sprintf("%f", *params.OperatingCashFlowPct)
	}
	if params.SellingGeneralAndAdministrativeExpensesPct != nil {
		urlParams["sellingGeneralAndAdministrativeExpensesPct"] = fmt.Sprintf("%f", *params.SellingGeneralAndAdministrativeExpensesPct)
	}
	if params.TaxRate != nil {
		urlParams["taxRate"] = fmt.Sprintf("%f", *params.TaxRate)
	}
	if params.LongTermGrowthRate != nil {
		urlParams["longTermGrowthRate"] = fmt.Sprintf("%f", *params.LongTermGrowthRate)
	}
	if params.CostOfDebt != nil {
		urlParams["costOfDebt"] = fmt.Sprintf("%f", *params.CostOfDebt)
	}
	if params.CostOfEquity != nil {
		urlParams["costOfEquity"] = fmt.Sprintf("%f", *params.CostOfEquity)
	}
	if params.MarketRiskPremium != nil {
		urlParams["marketRiskPremium"] = fmt.Sprintf("%f", *params.MarketRiskPremium)
	}
	if params.Beta != nil {
		urlParams["beta"] = fmt.Sprintf("%f", *params.Beta)
	}
	if params.RiskFreeRate != nil {
		urlParams["riskFreeRate"] = fmt.Sprintf("%f", *params.RiskFreeRate)
	}

	return doRequest[[]CustomDCFAdvancedResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
