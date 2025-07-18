package fmp_client

import "time"

// DividendsCompanyParams represents parameters for company dividends endpoint.
type DividendsCompanyParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// CalendarDateRangeParams represents parameters for calendar endpoints with date ranges.
type CalendarDateRangeParams struct {
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// EarningsReportParams represents parameters for earnings report endpoint.
type EarningsReportParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// StockSplitsParams represents parameters for stock splits endpoint.
type StockSplitsParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// DividendResponse represents dividend information.
type DividendResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the ex-dividend date.
	Date string `json:"date"`
	// RecordDate is when shareholders must own shares to receive dividend.
	RecordDate string `json:"recordDate"`
	// PaymentDate is when the dividend is paid.
	PaymentDate string `json:"paymentDate"`
	// DeclarationDate is when the dividend was announced.
	DeclarationDate string `json:"declarationDate"`
	// AdjDividend is the adjusted dividend amount.
	AdjDividend float64 `json:"adjDividend"`
	// Dividend is the dividend amount per share.
	Dividend float64 `json:"dividend"`
	// Yield is the dividend yield percentage.
	Yield float64 `json:"yield"`
	// Frequency describes how often dividends are paid.
	Frequency string `json:"frequency"`
}

// EarningsResponse represents earnings information.
type EarningsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the earnings report date.
	Date string `json:"date"`
	// EPSActual is the actual earnings per share.
	EPSActual *float64 `json:"epsActual"`
	// EPSEstimated is the estimated earnings per share.
	EPSEstimated *float64 `json:"epsEstimated"`
	// RevenueActual is the actual revenue.
	RevenueActual *int64 `json:"revenueActual"`
	// RevenueEstimated is the estimated revenue.
	RevenueEstimated *int64 `json:"revenueEstimated"`
	// LastUpdated is when the data was last updated.
	LastUpdated string `json:"lastUpdated"`
}

// IPOCalendarResponse represents IPO calendar information.
type IPOCalendarResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the expected IPO date.
	Date string `json:"date"`
	// DAA is the date of allocation announcement.
	DAA string `json:"daa"`
	// Company is the company name.
	Company string `json:"company"`
	// Exchange is where the stock will be listed.
	Exchange string `json:"exchange"`
	// Actions describes the IPO status.
	Actions string `json:"actions"`
	// Shares is the number of shares offered.
	Shares *int64 `json:"shares"`
	// PriceRange is the expected price range.
	PriceRange *string `json:"priceRange"`
	// MarketCap is the expected market capitalization.
	MarketCap *int64 `json:"marketCap"`
}

// IPODisclosureResponse represents IPO disclosure filing information.
type IPODisclosureResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// FilingDate is when the form was filed.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is when the filing was accepted.
	AcceptedDate string `json:"acceptedDate"`
	// EffectivenessDate is when the registration becomes effective.
	EffectivenessDate string `json:"effectivenessDate"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// Form is the SEC form type.
	Form string `json:"form"`
	// URL is the link to the SEC filing.
	URL string `json:"url"`
}

// IPOProspectusResponse represents IPO prospectus information.
type IPOProspectusResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// AcceptedDate is when the filing was accepted.
	AcceptedDate string `json:"acceptedDate"`
	// FilingDate is when the form was filed.
	FilingDate string `json:"filingDate"`
	// IPODate is the date of the IPO.
	IPODate string `json:"ipoDate"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// PricePublicPerShare is the public offering price per share.
	PricePublicPerShare float64 `json:"pricePublicPerShare"`
	// PricePublicTotal is the total public offering amount.
	PricePublicTotal float64 `json:"pricePublicTotal"`
	// DiscountsAndCommissionsPerShare is underwriting discounts per share.
	DiscountsAndCommissionsPerShare float64 `json:"discountsAndCommissionsPerShare"`
	// DiscountsAndCommissionsTotal is total underwriting discounts.
	DiscountsAndCommissionsTotal float64 `json:"discountsAndCommissionsTotal"`
	// ProceedsBeforeExpensesPerShare is net proceeds per share before expenses.
	ProceedsBeforeExpensesPerShare float64 `json:"proceedsBeforeExpensesPerShare"`
	// ProceedsBeforeExpensesTotal is total net proceeds before expenses.
	ProceedsBeforeExpensesTotal float64 `json:"proceedsBeforeExpensesTotal"`
	// Form is the SEC form type.
	Form string `json:"form"`
	// URL is the link to the SEC filing.
	URL string `json:"url"`
}

// StockSplitResponse represents stock split information.
type StockSplitResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the split date.
	Date string `json:"date"`
	// Numerator is the split numerator (shares after split).
	Numerator int `json:"numerator"`
	// Denominator is the split denominator (shares before split).
	Denominator int `json:"denominator"`
}

// DividendsCompany retrieves dividend history for a specific company.
// Returns past and upcoming dividend payments with related dates.
//
// Endpoint: /dividends
func (c *Client) DividendsCompany(params DividendsCompanyParams) ([]DividendResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/dividends"

	var result []DividendResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// DividendsCalendar retrieves dividend events for all companies within a date range.
// Returns upcoming dividend payments across the market.
//
// Endpoint: /dividends-calendar
func (c *Client) DividendsCalendar(params CalendarDateRangeParams) ([]DividendResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/dividends-calendar"

	var result []DividendResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// EarningsReport retrieves earnings report data for a specific company.
// Returns past and upcoming earnings with estimates and actuals.
//
// Endpoint: /earnings
func (c *Client) EarningsReport(params EarningsReportParams) ([]EarningsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/earnings"

	var result []EarningsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// EarningsCalendar retrieves earnings announcements within a date range.
// Returns earnings events across all companies.
//
// Endpoint: /earnings-calendar
func (c *Client) EarningsCalendar(params CalendarDateRangeParams) ([]EarningsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/earnings-calendar"

	var result []EarningsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IPOsCalendar retrieves upcoming initial public offerings.
// Returns companies expected to go public within the date range.
//
// Endpoint: /ipos-calendar
func (c *Client) IPOsCalendar(params CalendarDateRangeParams) ([]IPOCalendarResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/ipos-calendar"

	var result []IPOCalendarResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IPOsDisclosure retrieves IPO disclosure filings.
// Returns regulatory filings for upcoming IPOs.
//
// Endpoint: /ipos-disclosure
func (c *Client) IPOsDisclosure(params CalendarDateRangeParams) ([]IPODisclosureResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/ipos-disclosure"

	var result []IPODisclosureResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IPOsProspectus retrieves IPO prospectus information.
// Returns detailed financial data from IPO prospectuses.
//
// Endpoint: /ipos-prospectus
func (c *Client) IPOsProspectus(params CalendarDateRangeParams) ([]IPOProspectusResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/ipos-prospectus"

	var result []IPOProspectusResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockSplits retrieves stock split history for a specific company.
// Returns past stock splits with split ratios.
//
// Endpoint: /splits
func (c *Client) StockSplits(params StockSplitsParams) ([]StockSplitResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/splits"

	var result []StockSplitResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockSplitsCalendar retrieves upcoming stock splits across the market.
// Returns companies with announced stock splits within the date range.
//
// Endpoint: /splits-calendar
func (c *Client) StockSplitsCalendar(params CalendarDateRangeParams) ([]StockSplitResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/splits-calendar"

	var result []StockSplitResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}