package fmp_client

import "time"

// CompanyProfileParams represents parameters for company profile endpoint.
type CompanyProfileParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// CompanyProfileByCIKParams represents parameters for profile by CIK endpoint.
type CompanyProfileByCIKParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
}

// DelistedCompaniesParams represents parameters for delisted companies endpoint.
type DelistedCompaniesParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// EmployeeCountParams represents parameters for employee count endpoints.
type EmployeeCountParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// MarketCapDateParams represents parameters for historical market cap.
type MarketCapDateParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// BatchMarketCapParams represents parameters for batch market cap.
type BatchMarketCapParams struct {
	// Symbols is a comma-separated list of ticker symbols (required).
	Symbols string `json:"symbols"`
}

// AllSharesFloatParams represents parameters for all shares float endpoint.
type AllSharesFloatParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// MergersAcquisitionsParams represents parameters for M&A endpoints.
type MergersAcquisitionsParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// MergersAcquisitionsSearchParams represents parameters for M&A search.
type MergersAcquisitionsSearchParams struct {
	// Name is the company name to search for (required).
	Name string `json:"name"`
}

// KeyExecutivesParams represents parameters for key executives endpoint.
type KeyExecutivesParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Active filters for active executives only (optional).
	Active *string `json:"active,omitempty"`
}

// ExecutiveCompensationParams represents parameters for executive compensation.
type ExecutiveCompensationParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// ExecutiveCompensationBenchmarkParams represents benchmark parameters.
type ExecutiveCompensationBenchmarkParams struct {
	// Year is the year for compensation data (optional).
	Year *string `json:"year,omitempty"`
}

// CompanyProfileResponse represents detailed company profile data.
type CompanyProfileResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// MarketCap is the market capitalization.
	MarketCap int64 `json:"marketCap"`
	// Beta is the stock's beta coefficient.
	Beta float64 `json:"beta"`
	// LastDividend is the most recent dividend amount.
	LastDividend float64 `json:"lastDividend"`
	// Range is the 52-week price range.
	Range string `json:"range"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// ChangePercentage is the price change percentage.
	ChangePercentage float64 `json:"changePercentage"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
	// AverageVolume is the average trading volume.
	AverageVolume int64 `json:"averageVolume"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// Currency is the trading currency.
	Currency string `json:"currency"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// ISIN is the International Securities Identification Number.
	ISIN string `json:"isin"`
	// CUSIP is the CUSIP identifier.
	CUSIP string `json:"cusip"`
	// ExchangeFullName is the full name of the exchange.
	ExchangeFullName string `json:"exchangeFullName"`
	// Exchange is the short name of the exchange.
	Exchange string `json:"exchange"`
	// Industry is the company's industry.
	Industry string `json:"industry"`
	// Website is the company's website URL.
	Website string `json:"website"`
	// Description is a detailed company description.
	Description string `json:"description"`
	// CEO is the name of the CEO.
	CEO string `json:"ceo"`
	// Sector is the business sector.
	Sector string `json:"sector"`
	// Country is the country code.
	Country string `json:"country"`
	// FullTimeEmployees is the employee count.
	FullTimeEmployees string `json:"fullTimeEmployees"`
	// Phone is the company phone number.
	Phone string `json:"phone"`
	// Address is the street address.
	Address string `json:"address"`
	// City is the city name.
	City string `json:"city"`
	// State is the state/province code.
	State string `json:"state"`
	// ZIP is the postal code.
	ZIP string `json:"zip"`
	// Image is the company logo URL.
	Image string `json:"image"`
	// IPODate is the initial public offering date.
	IPODate string `json:"ipoDate"`
	// DefaultImage indicates if using default image.
	DefaultImage bool `json:"defaultImage"`
	// IsETF indicates if this is an ETF.
	IsETF bool `json:"isEtf"`
	// IsActivelyTrading indicates if actively trading.
	IsActivelyTrading bool `json:"isActivelyTrading"`
	// IsADR indicates if this is an ADR.
	IsADR bool `json:"isAdr"`
	// IsFund indicates if this is a fund.
	IsFund bool `json:"isFund"`
}

// CompanyNotesResponse represents company-issued notes information.
type CompanyNotesResponse struct {
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Title is the note title and terms.
	Title string `json:"title"`
	// Exchange is the exchange where notes are listed.
	Exchange string `json:"exchange"`
}

// StockPeersResponse represents peer comparison data.
type StockPeersResponse struct {
	// Symbol is the peer company's ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the peer company's name.
	CompanyName string `json:"companyName"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// MktCap is the market capitalization.
	MktCap int64 `json:"mktCap"`
}

// DelistedCompanyResponse represents a delisted company.
type DelistedCompanyResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the company name.
	CompanyName string `json:"companyName"`
	// Exchange is the exchange it was listed on.
	Exchange string `json:"exchange"`
	// IPODate is the initial public offering date.
	IPODate string `json:"ipoDate"`
	// DelistedDate is the date of delisting.
	DelistedDate string `json:"delistedDate"`
}

// EmployeeCountResponse represents employee count data.
type EmployeeCountResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// AcceptanceTime is when the filing was accepted.
	AcceptanceTime string `json:"acceptanceTime"`
	// PeriodOfReport is the reporting period date.
	PeriodOfReport string `json:"periodOfReport"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// FormType is the SEC form type.
	FormType string `json:"formType"`
	// FilingDate is when the form was filed.
	FilingDate string `json:"filingDate"`
	// EmployeeCount is the number of employees.
	EmployeeCount int `json:"employeeCount"`
	// Source is the URL to the SEC filing.
	Source string `json:"source"`
}

// MarketCapResponse represents market capitalization data.
type MarketCapResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the market cap value.
	Date string `json:"date"`
	// MarketCap is the market capitalization value.
	MarketCap int64 `json:"marketCap"`
}

// SharesFloatResponse represents shares float data.
type SharesFloatResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the date and time of the data.
	Date string `json:"date"`
	// FreeFloat is the percentage of freely traded shares.
	FreeFloat float64 `json:"freeFloat"`
	// FloatShares is the number of shares available for trading.
	FloatShares int64 `json:"floatShares"`
	// OutstandingShares is the total shares outstanding.
	OutstandingShares int64 `json:"outstandingShares"`
}

// MergersAcquisitionsResponse represents M&A transaction data.
type MergersAcquisitionsResponse struct {
	// Symbol is the acquiring company's ticker.
	Symbol string `json:"symbol"`
	// CompanyName is the acquiring company's name.
	CompanyName string `json:"companyName"`
	// CIK is the acquiring company's CIK.
	CIK string `json:"cik"`
	// TargetedCompanyName is the target company's name.
	TargetedCompanyName string `json:"targetedCompanyName"`
	// TargetedCIK is the target company's CIK.
	TargetedCIK string `json:"targetedCik"`
	// TargetedSymbol is the target company's ticker.
	TargetedSymbol string `json:"targetedSymbol"`
	// TransactionDate is the date of the transaction.
	TransactionDate string `json:"transactionDate"`
	// AcceptedDate is when the filing was accepted.
	AcceptedDate string `json:"acceptedDate"`
	// Link is the URL to the SEC filing.
	Link string `json:"link"`
}

// KeyExecutiveResponse represents executive information.
type KeyExecutiveResponse struct {
	// Title is the executive's job title.
	Title string `json:"title"`
	// Name is the executive's full name.
	Name string `json:"name"`
	// Pay is the compensation amount.
	Pay *float64 `json:"pay"`
	// CurrencyPay is the currency of compensation.
	CurrencyPay string `json:"currencyPay"`
	// Gender is the executive's gender.
	Gender string `json:"gender"`
	// YearBorn is the birth year.
	YearBorn *int `json:"yearBorn"`
	// Active indicates if currently active.
	Active *bool `json:"active"`
}

// ExecutiveCompensationResponse represents detailed compensation data.
type ExecutiveCompensationResponse struct {
	// CIK is the company's Central Index Key.
	CIK string `json:"cik"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// FilingDate is when the form was filed.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is when the filing was accepted.
	AcceptedDate string `json:"acceptedDate"`
	// NameAndPosition is the executive's name and title.
	NameAndPosition string `json:"nameAndPosition"`
	// Year is the compensation year.
	Year int `json:"year"`
	// Salary is the base salary amount.
	Salary int64 `json:"salary"`
	// Bonus is the bonus amount.
	Bonus int64 `json:"bonus"`
	// StockAward is the value of stock awards.
	StockAward int64 `json:"stockAward"`
	// OptionAward is the value of option awards.
	OptionAward int64 `json:"optionAward"`
	// IncentivePlanCompensation is incentive compensation.
	IncentivePlanCompensation int64 `json:"incentivePlanCompensation"`
	// AllOtherCompensation is other compensation.
	AllOtherCompensation int64 `json:"allOtherCompensation"`
	// Total is the total compensation.
	Total int64 `json:"total"`
	// Link is the URL to the SEC filing.
	Link string `json:"link"`
}

// ExecutiveCompensationBenchmarkResponse represents industry compensation benchmarks.
type ExecutiveCompensationBenchmarkResponse struct {
	// IndustryTitle is the industry name.
	IndustryTitle string `json:"industryTitle"`
	// Year is the benchmark year.
	Year int `json:"year"`
	// AverageCompensation is the industry average compensation.
	AverageCompensation float64 `json:"averageCompensation"`
}

// CompanyProfile retrieves detailed company profile data by ticker symbol.
// Returns comprehensive company information including financials and corporate details.
//
// Endpoint: /profile
func (c *Client) CompanyProfile(params CompanyProfileParams) ([]CompanyProfileResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/profile"

	var result []CompanyProfileResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CompanyProfileByCIK retrieves company profile data using CIK identifier.
// Returns the same comprehensive data as CompanyProfile but searched by CIK.
//
// Endpoint: /profile-cik
func (c *Client) CompanyProfileByCIK(params CompanyProfileByCIKParams) ([]CompanyProfileResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/profile-cik"

	var result []CompanyProfileResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CompanyNotes retrieves information about company-issued notes.
// Returns details about bonds and notes issued by the company.
//
// Endpoint: /company-notes
func (c *Client) CompanyNotes(params CompanyProfileParams) ([]CompanyNotesResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/company-notes"

	var result []CompanyNotesResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockPeers retrieves peer companies for comparison.
// Returns companies in the same sector and market cap range.
//
// Endpoint: /stock-peers
func (c *Client) StockPeers(params CompanyProfileParams) ([]StockPeersResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/stock-peers"

	var result []StockPeersResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// DelistedCompanies retrieves companies delisted from US exchanges.
// Returns information about companies no longer trading.
//
// Endpoint: /delisted-companies
func (c *Client) DelistedCompanies(params DelistedCompaniesParams) ([]DelistedCompanyResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/delisted-companies"

	var result []DelistedCompanyResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// EmployeeCount retrieves current employee count data.
// Returns workforce information from SEC filings.
//
// Endpoint: /employee-count
func (c *Client) EmployeeCount(params EmployeeCountParams) ([]EmployeeCountResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/employee-count"

	var result []EmployeeCountResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HistoricalEmployeeCount retrieves historical employee count data.
// Returns workforce trends over time from SEC filings.
//
// Endpoint: /historical-employee-count
func (c *Client) HistoricalEmployeeCount(params EmployeeCountParams) ([]EmployeeCountResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-employee-count"

	var result []EmployeeCountResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// MarketCapitalization retrieves current market cap for a company.
// Returns the total market value of outstanding shares.
//
// Endpoint: /market-capitalization
func (c *Client) MarketCapitalization(params CompanyProfileParams) ([]MarketCapResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/market-capitalization"

	var result []MarketCapResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BatchMarketCapitalization retrieves market cap for multiple companies.
// Returns market values for several companies in one request.
//
// Endpoint: /market-capitalization-batch
func (c *Client) BatchMarketCapitalization(params BatchMarketCapParams) ([]MarketCapResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/market-capitalization-batch"

	var result []MarketCapResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HistoricalMarketCapitalization retrieves historical market cap data.
// Returns market value changes over time for analysis.
//
// Endpoint: /historical-market-capitalization
func (c *Client) HistoricalMarketCapitalization(params MarketCapDateParams) ([]MarketCapResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-market-capitalization"

	var result []MarketCapResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SharesFloat retrieves shares float and liquidity data.
// Returns information about publicly tradable shares.
//
// Endpoint: /shares-float
func (c *Client) SharesFloat(params CompanyProfileParams) ([]SharesFloatResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/shares-float"

	var result []SharesFloatResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AllSharesFloat retrieves shares float data for all companies.
// Returns comprehensive float data across the market.
//
// Endpoint: /shares-float-all
func (c *Client) AllSharesFloat(params AllSharesFloatParams) ([]SharesFloatResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/shares-float-all"

	var result []SharesFloatResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// LatestMergersAcquisitions retrieves recent M&A transactions.
// Returns the latest merger and acquisition activity.
//
// Endpoint: /mergers-acquisitions-latest
func (c *Client) LatestMergersAcquisitions(params MergersAcquisitionsParams) ([]MergersAcquisitionsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/mergers-acquisitions-latest"

	var result []MergersAcquisitionsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchMergersAcquisitions searches for specific M&A transactions.
// Returns M&A activity matching the search criteria.
//
// Endpoint: /mergers-acquisitions-search
func (c *Client) SearchMergersAcquisitions(params MergersAcquisitionsSearchParams) ([]MergersAcquisitionsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/mergers-acquisitions-search"

	var result []MergersAcquisitionsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// KeyExecutives retrieves information about company executives.
// Returns details about key management personnel.
//
// Endpoint: /key-executives
func (c *Client) KeyExecutives(params KeyExecutivesParams) ([]KeyExecutiveResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/key-executives"

	var result []KeyExecutiveResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ExecutiveCompensation retrieves detailed executive pay data.
// Returns comprehensive compensation information from proxy statements.
//
// Endpoint: /governance-executive-compensation
func (c *Client) ExecutiveCompensation(params ExecutiveCompensationParams) ([]ExecutiveCompensationResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/governance-executive-compensation"

	var result []ExecutiveCompensationResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ExecutiveCompensationBenchmark retrieves industry compensation averages.
// Returns benchmark data for executive pay by industry.
//
// Endpoint: /executive-compensation-benchmark
func (c *Client) ExecutiveCompensationBenchmark(params ExecutiveCompensationBenchmarkParams) ([]ExecutiveCompensationBenchmarkResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/executive-compensation-benchmark"

	var result []ExecutiveCompensationBenchmarkResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}