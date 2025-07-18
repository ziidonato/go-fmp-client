package fmp_client

// DCFParams represents parameters for DCF valuation endpoints.
type DCFParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// CustomDCFParams represents parameters for custom DCF calculations.
type CustomDCFParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// RevenueGrowthPct is the revenue growth percentage (optional).
	RevenueGrowthPct *float64 `json:"revenueGrowthPct,omitempty"`
	// EbitdaPct is the EBITDA percentage (optional).
	EbitdaPct *float64 `json:"ebitdaPct,omitempty"`
	// DepreciationAndAmortizationPct is the D&A percentage (optional).
	DepreciationAndAmortizationPct *float64 `json:"depreciationAndAmortizationPct,omitempty"`
	// CashAndShortTermInvestmentsPct is the cash percentage (optional).
	CashAndShortTermInvestmentsPct *float64 `json:"cashAndShortTermInvestmentsPct,omitempty"`
	// ReceivablesPct is the receivables percentage (optional).
	ReceivablesPct *float64 `json:"receivablesPct,omitempty"`
	// InventoriesPct is the inventories percentage (optional).
	InventoriesPct *float64 `json:"inventoriesPct,omitempty"`
	// PayablePct is the payables percentage (optional).
	PayablePct *float64 `json:"payablePct,omitempty"`
	// EbitPct is the EBIT percentage (optional).
	EbitPct *float64 `json:"ebitPct,omitempty"`
	// CapitalExpenditurePct is the capex percentage (optional).
	CapitalExpenditurePct *float64 `json:"capitalExpenditurePct,omitempty"`
	// OperatingCashFlowPct is the operating cash flow percentage (optional).
	OperatingCashFlowPct *float64 `json:"operatingCashFlowPct,omitempty"`
	// SellingGeneralAndAdministrativeExpensesPct is the SG&A percentage (optional).
	SellingGeneralAndAdministrativeExpensesPct *float64 `json:"sellingGeneralAndAdministrativeExpensesPct,omitempty"`
	// TaxRate is the tax rate (optional).
	TaxRate *float64 `json:"taxRate,omitempty"`
	// LongTermGrowthRate is the terminal growth rate (optional).
	LongTermGrowthRate *float64 `json:"longTermGrowthRate,omitempty"`
	// CostOfDebt is the cost of debt percentage (optional).
	CostOfDebt *float64 `json:"costOfDebt,omitempty"`
	// CostOfEquity is the cost of equity percentage (optional).
	CostOfEquity *float64 `json:"costOfEquity,omitempty"`
	// MarketRiskPremium is the market risk premium (optional).
	MarketRiskPremium *float64 `json:"marketRiskPremium,omitempty"`
	// Beta is the stock's beta coefficient (optional).
	Beta *float64 `json:"beta,omitempty"`
	// RiskFreeRate is the risk-free rate (optional).
	RiskFreeRate *float64 `json:"riskFreeRate,omitempty"`
}

// DCFResponse represents basic DCF valuation results.
type DCFResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the valuation date.
	Date string `json:"date"`
	// DCF is the discounted cash flow value per share.
	DCF float64 `json:"dcf"`
	// StockPrice is the current stock price.
	StockPrice float64 `json:"Stock Price"`
}

// CustomDCFResponse represents detailed custom DCF calculation results.
type CustomDCFResponse struct {
	// Year is the projection year.
	Year string `json:"year"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Revenue is the projected revenue.
	Revenue int64 `json:"revenue"`
	// RevenuePercentage is the revenue growth percentage.
	RevenuePercentage float64 `json:"revenuePercentage"`
	// EBITDA is the projected EBITDA.
	EBITDA int64 `json:"ebitda"`
	// EBITDAPercentage is the EBITDA margin.
	EBITDAPercentage float64 `json:"ebitdaPercentage"`
	// EBIT is the projected EBIT.
	EBIT int64 `json:"ebit"`
	// EBITPercentage is the EBIT margin.
	EBITPercentage float64 `json:"ebitPercentage"`
	// Depreciation is the depreciation amount.
	Depreciation int64 `json:"depreciation"`
	// DepreciationPercentage is the depreciation percentage.
	DepreciationPercentage float64 `json:"depreciationPercentage"`
	// TotalCash is the total cash.
	TotalCash int64 `json:"totalCash"`
	// TotalCashPercentage is the cash percentage.
	TotalCashPercentage float64 `json:"totalCashPercentage"`
	// Receivables is the receivables amount.
	Receivables int64 `json:"receivables"`
	// ReceivablesPercentage is the receivables percentage.
	ReceivablesPercentage float64 `json:"receivablesPercentage"`
	// Inventories is the inventories amount.
	Inventories int64 `json:"inventories"`
	// InventoriesPercentage is the inventories percentage.
	InventoriesPercentage float64 `json:"inventoriesPercentage"`
	// Payable is the payables amount.
	Payable int64 `json:"payable"`
	// PayablePercentage is the payables percentage.
	PayablePercentage float64 `json:"payablePercentage"`
	// CapitalExpenditure is the capex amount.
	CapitalExpenditure int64 `json:"capitalExpenditure"`
	// CapitalExpenditurePercentage is the capex percentage.
	CapitalExpenditurePercentage float64 `json:"capitalExpenditurePercentage"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// Beta is the beta coefficient.
	Beta float64 `json:"beta"`
	// DilutedSharesOutstanding is the share count.
	DilutedSharesOutstanding int64 `json:"dilutedSharesOutstanding"`
	// CostOfDebt is the cost of debt.
	CostOfDebt float64 `json:"costofDebt"`
	// TaxRate is the tax rate.
	TaxRate float64 `json:"taxRate"`
	// AfterTaxCostOfDebt is the after-tax cost of debt.
	AfterTaxCostOfDebt float64 `json:"afterTaxCostOfDebt"`
	// RiskFreeRate is the risk-free rate.
	RiskFreeRate float64 `json:"riskFreeRate"`
	// MarketRiskPremium is the market risk premium.
	MarketRiskPremium float64 `json:"marketRiskPremium"`
	// CostOfEquity is the cost of equity.
	CostOfEquity float64 `json:"costOfEquity"`
	// TotalDebt is the total debt.
	TotalDebt int64 `json:"totalDebt"`
	// TotalEquity is the total equity.
	TotalEquity int64 `json:"totalEquity"`
	// TotalCapital is the total capital.
	TotalCapital int64 `json:"totalCapital"`
	// DebtWeighting is the debt weight percentage.
	DebtWeighting float64 `json:"debtWeighting"`
	// EquityWeighting is the equity weight percentage.
	EquityWeighting float64 `json:"equityWeighting"`
	// WACC is the weighted average cost of capital.
	WACC float64 `json:"wacc"`
	// TaxRateCash is the cash tax amount.
	TaxRateCash int64 `json:"taxRateCash"`
	// EBIAT is the EBIT after tax.
	EBIAT int64 `json:"ebiat"`
	// UFCF is the unlevered free cash flow.
	UFCF int64 `json:"ufcf"`
	// SumPvUFCF is the sum of present value of UFCF.
	SumPvUFCF int64 `json:"sumPvUfcf"`
	// LongTermGrowthRate is the terminal growth rate.
	LongTermGrowthRate float64 `json:"longTermGrowthRate"`
	// TerminalValue is the terminal value.
	TerminalValue int64 `json:"terminalValue"`
	// PresentTerminalValue is the present value of terminal value.
	PresentTerminalValue int64 `json:"presentTerminalValue"`
	// EnterpriseValue is the enterprise value.
	EnterpriseValue int64 `json:"enterpriseValue"`
	// NetDebt is the net debt.
	NetDebt int64 `json:"netDebt"`
	// EquityValue is the equity value.
	EquityValue int64 `json:"equityValue"`
	// EquityValuePerShare is the value per share.
	EquityValuePerShare float64 `json:"equityValuePerShare"`
	// FreeCashFlowT1 is the free cash flow for next period.
	FreeCashFlowT1 int64 `json:"freeCashFlowT1"`
}

// CustomDCFLeveredResponse represents levered DCF calculation results.
type CustomDCFLeveredResponse struct {
	// Year is the projection year.
	Year string `json:"year"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Revenue is the projected revenue.
	Revenue int64 `json:"revenue"`
	// RevenuePercentage is the revenue growth percentage.
	RevenuePercentage float64 `json:"revenuePercentage"`
	// CapitalExpenditure is the capex amount.
	CapitalExpenditure int64 `json:"capitalExpenditure"`
	// CapitalExpenditurePercentage is the capex percentage.
	CapitalExpenditurePercentage float64 `json:"capitalExpenditurePercentage"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// Beta is the beta coefficient.
	Beta float64 `json:"beta"`
	// DilutedSharesOutstanding is the share count.
	DilutedSharesOutstanding int64 `json:"dilutedSharesOutstanding"`
	// CostOfDebt is the cost of debt.
	CostOfDebt float64 `json:"costofDebt"`
	// TaxRate is the tax rate.
	TaxRate float64 `json:"taxRate"`
	// AfterTaxCostOfDebt is the after-tax cost of debt.
	AfterTaxCostOfDebt float64 `json:"afterTaxCostOfDebt"`
	// RiskFreeRate is the risk-free rate.
	RiskFreeRate float64 `json:"riskFreeRate"`
	// MarketRiskPremium is the market risk premium.
	MarketRiskPremium float64 `json:"marketRiskPremium"`
	// CostOfEquity is the cost of equity.
	CostOfEquity float64 `json:"costOfEquity"`
	// TotalDebt is the total debt.
	TotalDebt int64 `json:"totalDebt"`
	// TotalEquity is the total equity.
	TotalEquity int64 `json:"totalEquity"`
	// TotalCapital is the total capital.
	TotalCapital int64 `json:"totalCapital"`
	// DebtWeighting is the debt weight percentage.
	DebtWeighting float64 `json:"debtWeighting"`
	// EquityWeighting is the equity weight percentage.
	EquityWeighting float64 `json:"equityWeighting"`
	// WACC is the weighted average cost of capital.
	WACC float64 `json:"wacc"`
	// OperatingCashFlow is the operating cash flow.
	OperatingCashFlow int64 `json:"operatingCashFlow"`
	// PvLFCF is the present value of levered FCF.
	PvLFCF int64 `json:"pvLfcf"`
	// SumPvLFCF is the sum of present value of LFCF.
	SumPvLFCF int64 `json:"sumPvLfcf"`
	// LongTermGrowthRate is the terminal growth rate.
	LongTermGrowthRate float64 `json:"longTermGrowthRate"`
	// FreeCashFlow is the free cash flow.
	FreeCashFlow int64 `json:"freeCashFlow"`
	// TerminalValue is the terminal value.
	TerminalValue int64 `json:"terminalValue"`
	// PresentTerminalValue is the present value of terminal value.
	PresentTerminalValue int64 `json:"presentTerminalValue"`
	// EnterpriseValue is the enterprise value.
	EnterpriseValue int64 `json:"enterpriseValue"`
	// NetDebt is the net debt.
	NetDebt int64 `json:"netDebt"`
	// EquityValue is the equity value.
	EquityValue int64 `json:"equityValue"`
	// EquityValuePerShare is the value per share.
	EquityValuePerShare float64 `json:"equityValuePerShare"`
	// FreeCashFlowT1 is the free cash flow for next period.
	FreeCashFlowT1 int64 `json:"freeCashFlowT1"`
	// OperatingCashFlowPercentage is the operating cash flow margin.
	OperatingCashFlowPercentage float64 `json:"operatingCashFlowPercentage"`
}

// DiscountedCashFlow retrieves basic DCF valuation for a company.
// Returns the intrinsic value based on discounted future cash flows.
//
// Endpoint: /discounted-cash-flow
func (c *Client) DiscountedCashFlow(params DCFParams) ([]DCFResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/discounted-cash-flow"

	var result []DCFResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// LeveredDiscountedCashFlow retrieves levered DCF valuation.
// Returns valuation incorporating the impact of debt.
//
// Endpoint: /levered-discounted-cash-flow
func (c *Client) LeveredDiscountedCashFlow(params DCFParams) ([]DCFResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/levered-discounted-cash-flow"

	var result []DCFResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CustomDiscountedCashFlow performs custom DCF analysis with user inputs.
// Returns detailed valuation with customizable assumptions.
//
// Endpoint: /custom-discounted-cash-flow
func (c *Client) CustomDiscountedCashFlow(params CustomDCFParams) ([]CustomDCFResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/custom-discounted-cash-flow"

	var result []CustomDCFResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CustomLeveredDiscountedCashFlow performs custom levered DCF analysis.
// Returns detailed levered valuation with customizable assumptions.
//
// Endpoint: /custom-levered-discounted-cash-flow
func (c *Client) CustomLeveredDiscountedCashFlow(params CustomDCFParams) ([]CustomDCFLeveredResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/custom-levered-discounted-cash-flow"

	var result []CustomDCFLeveredResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}