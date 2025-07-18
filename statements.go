package fmp_client

// StatementParams represents parameters for financial statement endpoints.
type StatementParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
	// Period is the time period: "Q1", "Q2", "Q3", "Q4", "FY", "annual", "quarter" (optional).
	Period *string `json:"period,omitempty"`
}

// AdvancedStatementParams represents parameters for advanced statement endpoints.
type AdvancedStatementParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Period is the reporting period: "annual" or "quarter" (required).
	Period string `json:"period"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// BulkStatementParams represents parameters for bulk statement requests.
type BulkStatementParams struct {
	// Year is the year for the statements (required).
	Year int `json:"year"`
	// Period is the reporting period: "annual" or specific quarter (required).
	Period string `json:"period"`
}

// IncomeStatementResponse represents income statement data.
type IncomeStatementResponse struct {
	// Date is the reporting date.
	Date string `json:"date"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// ReportedCurrency is the currency of the report.
	ReportedCurrency string `json:"reportedCurrency"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// FilingDate is the filing date.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is the accepted date.
	AcceptedDate string `json:"acceptedDate"`
	// FiscalYear is the fiscal year.
	FiscalYear string `json:"fiscalYear"`
	// Period is the reporting period.
	Period string `json:"period"`
	// Revenue is total revenue.
	Revenue int64 `json:"revenue"`
	// CostOfRevenue is cost of revenue.
	CostOfRevenue int64 `json:"costOfRevenue"`
	// GrossProfit is gross profit.
	GrossProfit int64 `json:"grossProfit"`
	// ResearchAndDevelopmentExpenses is R&D expenses.
	ResearchAndDevelopmentExpenses int64 `json:"researchAndDevelopmentExpenses"`
	// GeneralAndAdministrativeExpenses is G&A expenses.
	GeneralAndAdministrativeExpenses int64 `json:"generalAndAdministrativeExpenses"`
	// SellingAndMarketingExpenses is S&M expenses.
	SellingAndMarketingExpenses int64 `json:"sellingAndMarketingExpenses"`
	// SellingGeneralAndAdministrativeExpenses is SG&A expenses.
	SellingGeneralAndAdministrativeExpenses int64 `json:"sellingGeneralAndAdministrativeExpenses"`
	// OtherExpenses is other expenses.
	OtherExpenses int64 `json:"otherExpenses"`
	// OperatingExpenses is total operating expenses.
	OperatingExpenses int64 `json:"operatingExpenses"`
	// CostAndExpenses is total costs and expenses.
	CostAndExpenses int64 `json:"costAndExpenses"`
	// NetInterestIncome is net interest income.
	NetInterestIncome int64 `json:"netInterestIncome"`
	// InterestIncome is interest income.
	InterestIncome int64 `json:"interestIncome"`
	// InterestExpense is interest expense.
	InterestExpense int64 `json:"interestExpense"`
	// DepreciationAndAmortization is D&A.
	DepreciationAndAmortization int64 `json:"depreciationAndAmortization"`
	// EBITDA is earnings before interest, taxes, D&A.
	EBITDA int64 `json:"ebitda"`
	// EBIT is earnings before interest and taxes.
	EBIT int64 `json:"ebit"`
	// NonOperatingIncomeExcludingInterest is non-operating income.
	NonOperatingIncomeExcludingInterest int64 `json:"nonOperatingIncomeExcludingInterest"`
	// OperatingIncome is operating income.
	OperatingIncome int64 `json:"operatingIncome"`
	// TotalOtherIncomeExpensesNet is other income/expenses.
	TotalOtherIncomeExpensesNet int64 `json:"totalOtherIncomeExpensesNet"`
	// IncomeBeforeTax is pre-tax income.
	IncomeBeforeTax int64 `json:"incomeBeforeTax"`
	// IncomeTaxExpense is income tax expense.
	IncomeTaxExpense int64 `json:"incomeTaxExpense"`
	// NetIncomeFromContinuingOperations is income from continuing ops.
	NetIncomeFromContinuingOperations int64 `json:"netIncomeFromContinuingOperations"`
	// NetIncomeFromDiscontinuedOperations is income from discontinued ops.
	NetIncomeFromDiscontinuedOperations int64 `json:"netIncomeFromDiscontinuedOperations"`
	// OtherAdjustmentsToNetIncome is other adjustments.
	OtherAdjustmentsToNetIncome int64 `json:"otherAdjustmentsToNetIncome"`
	// NetIncome is net income.
	NetIncome int64 `json:"netIncome"`
	// NetIncomeDeductions is net income deductions.
	NetIncomeDeductions int64 `json:"netIncomeDeductions"`
	// BottomLineNetIncome is bottom line net income.
	BottomLineNetIncome int64 `json:"bottomLineNetIncome"`
	// EPS is earnings per share.
	EPS float64 `json:"eps"`
	// EPSDiluted is diluted EPS.
	EPSDiluted float64 `json:"epsDiluted"`
	// WeightedAverageShsOut is weighted average shares outstanding.
	WeightedAverageShsOut int64 `json:"weightedAverageShsOut"`
	// WeightedAverageShsOutDil is diluted weighted average shares.
	WeightedAverageShsOutDil int64 `json:"weightedAverageShsOutDil"`
}

// BalanceSheetResponse represents balance sheet data.
type BalanceSheetResponse struct {
	// Date is the reporting date.
	Date string `json:"date"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// ReportedCurrency is the currency of the report.
	ReportedCurrency string `json:"reportedCurrency"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// FilingDate is the filing date.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is the accepted date.
	AcceptedDate string `json:"acceptedDate"`
	// FiscalYear is the fiscal year.
	FiscalYear string `json:"fiscalYear"`
	// Period is the reporting period.
	Period string `json:"period"`
	// CashAndCashEquivalents is cash and cash equivalents.
	CashAndCashEquivalents int64 `json:"cashAndCashEquivalents"`
	// ShortTermInvestments is short-term investments.
	ShortTermInvestments int64 `json:"shortTermInvestments"`
	// CashAndShortTermInvestments is total cash and short-term investments.
	CashAndShortTermInvestments int64 `json:"cashAndShortTermInvestments"`
	// NetReceivables is net receivables.
	NetReceivables int64 `json:"netReceivables"`
	// Inventory is inventory.
	Inventory int64 `json:"inventory"`
	// OtherCurrentAssets is other current assets.
	OtherCurrentAssets int64 `json:"otherCurrentAssets"`
	// TotalCurrentAssets is total current assets.
	TotalCurrentAssets int64 `json:"totalCurrentAssets"`
	// PropertyPlantEquipmentNet is net PP&E.
	PropertyPlantEquipmentNet int64 `json:"propertyPlantEquipmentNet"`
	// Goodwill is goodwill.
	Goodwill int64 `json:"goodwill"`
	// IntangibleAssets is intangible assets.
	IntangibleAssets int64 `json:"intangibleAssets"`
	// GoodwillAndIntangibleAssets is total goodwill and intangibles.
	GoodwillAndIntangibleAssets int64 `json:"goodwillAndIntangibleAssets"`
	// LongTermInvestments is long-term investments.
	LongTermInvestments int64 `json:"longTermInvestments"`
	// TaxAssets is tax assets.
	TaxAssets int64 `json:"taxAssets"`
	// OtherNonCurrentAssets is other non-current assets.
	OtherNonCurrentAssets int64 `json:"otherNonCurrentAssets"`
	// TotalNonCurrentAssets is total non-current assets.
	TotalNonCurrentAssets int64 `json:"totalNonCurrentAssets"`
	// OtherAssets is other assets.
	OtherAssets int64 `json:"otherAssets"`
	// TotalAssets is total assets.
	TotalAssets int64 `json:"totalAssets"`
	// AccountPayables is accounts payable.
	AccountPayables int64 `json:"accountPayables"`
	// ShortTermDebt is short-term debt.
	ShortTermDebt int64 `json:"shortTermDebt"`
	// TaxPayables is tax payables.
	TaxPayables int64 `json:"taxPayables"`
	// DeferredRevenue is deferred revenue.
	DeferredRevenue int64 `json:"deferredRevenue"`
	// OtherCurrentLiabilities is other current liabilities.
	OtherCurrentLiabilities int64 `json:"otherCurrentLiabilities"`
	// TotalCurrentLiabilities is total current liabilities.
	TotalCurrentLiabilities int64 `json:"totalCurrentLiabilities"`
	// LongTermDebt is long-term debt.
	LongTermDebt int64 `json:"longTermDebt"`
	// DeferredRevenueNonCurrent is non-current deferred revenue.
	DeferredRevenueNonCurrent int64 `json:"deferredRevenueNonCurrent"`
	// DeferredTaxLiabilitiesNonCurrent is deferred tax liabilities.
	DeferredTaxLiabilitiesNonCurrent int64 `json:"deferredTaxLiabilitiesNonCurrent"`
	// OtherNonCurrentLiabilities is other non-current liabilities.
	OtherNonCurrentLiabilities int64 `json:"otherNonCurrentLiabilities"`
	// TotalNonCurrentLiabilities is total non-current liabilities.
	TotalNonCurrentLiabilities int64 `json:"totalNonCurrentLiabilities"`
	// OtherLiabilities is other liabilities.
	OtherLiabilities int64 `json:"otherLiabilities"`
	// CapitalLeaseObligations is capital lease obligations.
	CapitalLeaseObligations int64 `json:"capitalLeaseObligations"`
	// TotalLiabilities is total liabilities.
	TotalLiabilities int64 `json:"totalLiabilities"`
	// PreferredStock is preferred stock.
	PreferredStock int64 `json:"preferredStock"`
	// CommonStock is common stock.
	CommonStock int64 `json:"commonStock"`
	// RetainedEarnings is retained earnings.
	RetainedEarnings int64 `json:"retainedEarnings"`
	// AccumulatedOtherComprehensiveIncomeLoss is AOCI.
	AccumulatedOtherComprehensiveIncomeLoss int64 `json:"accumulatedOtherComprehensiveIncomeLoss"`
	// OtherTotalStockholdersEquity is other equity.
	OtherTotalStockholdersEquity int64 `json:"otherTotalStockholdersEquity"`
	// TotalStockholdersEquity is total stockholders' equity.
	TotalStockholdersEquity int64 `json:"totalStockholdersEquity"`
	// TotalEquity is total equity.
	TotalEquity int64 `json:"totalEquity"`
	// TotalLiabilitiesAndStockholdersEquity is total liabilities and equity.
	TotalLiabilitiesAndStockholdersEquity int64 `json:"totalLiabilitiesAndStockholdersEquity"`
	// MinorityInterest is minority interest.
	MinorityInterest int64 `json:"minorityInterest"`
	// TotalLiabilitiesAndTotalEquity is total liabilities and total equity.
	TotalLiabilitiesAndTotalEquity int64 `json:"totalLiabilitiesAndTotalEquity"`
	// TotalInvestments is total investments.
	TotalInvestments int64 `json:"totalInvestments"`
	// TotalDebt is total debt.
	TotalDebt int64 `json:"totalDebt"`
	// NetDebt is net debt.
	NetDebt int64 `json:"netDebt"`
}

// CashFlowResponse represents cash flow statement data.
type CashFlowResponse struct {
	// Date is the reporting date.
	Date string `json:"date"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// ReportedCurrency is the currency of the report.
	ReportedCurrency string `json:"reportedCurrency"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// FilingDate is the filing date.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is the accepted date.
	AcceptedDate string `json:"acceptedDate"`
	// FiscalYear is the fiscal year.
	FiscalYear string `json:"fiscalYear"`
	// Period is the reporting period.
	Period string `json:"period"`
	// NetIncome is net income.
	NetIncome int64 `json:"netIncome"`
	// DepreciationAndAmortization is D&A.
	DepreciationAndAmortization int64 `json:"depreciationAndAmortization"`
	// StockBasedCompensation is stock-based compensation.
	StockBasedCompensation int64 `json:"stockBasedCompensation"`
	// DeferredIncomeTax is deferred income tax.
	DeferredIncomeTax int64 `json:"deferredIncomeTax"`
	// ChangeInWorkingCapital is change in working capital.
	ChangeInWorkingCapital int64 `json:"changeInWorkingCapital"`
	// AccountsReceivables is change in accounts receivable.
	AccountsReceivables int64 `json:"accountsReceivables"`
	// Inventory is change in inventory.
	Inventory int64 `json:"inventory"`
	// AccountsPayables is change in accounts payable.
	AccountsPayables int64 `json:"accountsPayables"`
	// OtherWorkingCapital is other working capital changes.
	OtherWorkingCapital int64 `json:"otherWorkingCapital"`
	// OtherOperatingActivities is other operating activities.
	OtherOperatingActivities int64 `json:"otherOperatingActivities"`
	// NetCashProvidedByOperatingActivities is operating cash flow.
	NetCashProvidedByOperatingActivities int64 `json:"netCashProvidedByOperatingActivities"`
	// InvestmentsInPropertyPlantAndEquipment is CapEx.
	InvestmentsInPropertyPlantAndEquipment int64 `json:"investmentsInPropertyPlantAndEquipment"`
	// AcquisitionsNet is net acquisitions.
	AcquisitionsNet int64 `json:"acquisitionsNet"`
	// PurchasesOfInvestments is investment purchases.
	PurchasesOfInvestments int64 `json:"purchasesOfInvestments"`
	// SalesMaturitiesOfInvestments is investment sales/maturities.
	SalesMaturitiesOfInvestments int64 `json:"salesMaturitiesOfInvestments"`
	// OtherInvestingActivities is other investing activities.
	OtherInvestingActivities int64 `json:"otherInvestingActivities"`
	// NetCashUsedForInvestingActivities is investing cash flow.
	NetCashUsedForInvestingActivities int64 `json:"netCashUsedForInvestingActivities"`
	// DebtRepayment is debt repayment.
	DebtRepayment int64 `json:"debtRepayment"`
	// CommonStockIssued is common stock issued.
	CommonStockIssued int64 `json:"commonStockIssued"`
	// CommonStockRepurchased is common stock repurchased.
	CommonStockRepurchased int64 `json:"commonStockRepurchased"`
	// DividendsPaid is dividends paid.
	DividendsPaid int64 `json:"dividendsPaid"`
	// OtherFinancingActivities is other financing activities.
	OtherFinancingActivities int64 `json:"otherFinancingActivities"`
	// NetCashUsedProvidedByFinancingActivities is financing cash flow.
	NetCashUsedProvidedByFinancingActivities int64 `json:"netCashUsedProvidedByFinancingActivities"`
	// EffectOfForexChangesOnCash is FX effect on cash.
	EffectOfForexChangesOnCash int64 `json:"effectOfForexChangesOnCash"`
	// NetChangeInCash is net change in cash.
	NetChangeInCash int64 `json:"netChangeInCash"`
	// CashAtEndOfPeriod is ending cash balance.
	CashAtEndOfPeriod int64 `json:"cashAtEndOfPeriod"`
	// CashAtBeginningOfPeriod is beginning cash balance.
	CashAtBeginningOfPeriod int64 `json:"cashAtBeginningOfPeriod"`
	// OperatingCashFlow is operating cash flow.
	OperatingCashFlow int64 `json:"operatingCashFlow"`
	// CapitalExpenditure is capital expenditures.
	CapitalExpenditure int64 `json:"capitalExpenditure"`
	// FreeCashFlow is free cash flow.
	FreeCashFlow int64 `json:"freeCashFlow"`
}

// StatementSymbolsResponse represents available symbols.
type StatementSymbolsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
}

// IncomeStatement retrieves income statement data.
// Returns revenue, expenses, and profitability metrics.
//
// Endpoint: /income-statement
func (c *Client) IncomeStatement(params StatementParams) ([]IncomeStatementResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/income-statement"

	var result []IncomeStatementResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BalanceSheet retrieves balance sheet data.
// Returns assets, liabilities, and equity information.
//
// Endpoint: /balance-sheet-statement
func (c *Client) BalanceSheet(params StatementParams) ([]BalanceSheetResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/balance-sheet-statement"

	var result []BalanceSheetResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CashFlowStatement retrieves cash flow statement data.
// Returns operating, investing, and financing cash flows.
//
// Endpoint: /cash-flow-statement
func (c *Client) CashFlowStatement(params StatementParams) ([]CashFlowResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/cash-flow-statement"

	var result []CashFlowResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IncomeStatementAsReported retrieves as-reported income statement.
// Returns income statement in original reported format.
//
// Endpoint: /income-statement-as-reported
func (c *Client) IncomeStatementAsReported(params AdvancedStatementParams) ([]map[string]interface{}, error) {
	queryParams := StructToMap(params)
	pathName := "/income-statement-as-reported"

	var result []map[string]interface{}
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BalanceSheetAsReported retrieves as-reported balance sheet.
// Returns balance sheet in original reported format.
//
// Endpoint: /balance-sheet-statement-as-reported
func (c *Client) BalanceSheetAsReported(params AdvancedStatementParams) ([]map[string]interface{}, error) {
	queryParams := StructToMap(params)
	pathName := "/balance-sheet-statement-as-reported"

	var result []map[string]interface{}
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CashFlowStatementAsReported retrieves as-reported cash flow.
// Returns cash flow statement in original reported format.
//
// Endpoint: /cash-flow-statement-as-reported
func (c *Client) CashFlowStatementAsReported(params AdvancedStatementParams) ([]map[string]interface{}, error) {
	queryParams := StructToMap(params)
	pathName := "/cash-flow-statement-as-reported"

	var result []map[string]interface{}
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// FullFinancialStatementAsReported retrieves all as-reported statements.
// Returns complete financial statements in original format.
//
// Endpoint: /financial-statement-full-as-reported
func (c *Client) FullFinancialStatementAsReported(params AdvancedStatementParams) ([]map[string]interface{}, error) {
	queryParams := StructToMap(params)
	pathName := "/financial-statement-full-as-reported"

	var result []map[string]interface{}
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// FinancialStatementSymbolsList retrieves available symbols.
// Returns list of symbols with financial statements.
//
// Endpoint: /financial-statement-symbols-list
func (c *Client) FinancialStatementSymbolsList() ([]StatementSymbolsResponse, error) {
	pathName := "/financial-statement-symbols-list"

	var result []StatementSymbolsResponse
	err := c.doRequest(pathName, nil, &result)
	return result, err
}

// BulkIncomeStatements retrieves bulk income statements.
// Returns income statements for multiple companies.
//
// Endpoint: /income-statement-bulk
func (c *Client) BulkIncomeStatements(params BulkStatementParams) ([]IncomeStatementResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/income-statement-bulk"

	var result []IncomeStatementResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BulkBalanceSheets retrieves bulk balance sheets.
// Returns balance sheets for multiple companies.
//
// Endpoint: /balance-sheet-statement-bulk
func (c *Client) BulkBalanceSheets(params BulkStatementParams) ([]BalanceSheetResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/balance-sheet-statement-bulk"

	var result []BalanceSheetResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BulkCashFlowStatements retrieves bulk cash flow statements.
// Returns cash flow statements for multiple companies.
//
// Endpoint: /cash-flow-statement-bulk
func (c *Client) BulkCashFlowStatements(params BulkStatementParams) ([]CashFlowResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/cash-flow-statement-bulk"

	var result []CashFlowResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}