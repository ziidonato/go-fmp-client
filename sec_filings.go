package fmp_client

import "time"

// SECFilingsParams represents parameters for SEC filings endpoints.
type SECFilingsParams struct {
	// From is the start date for filtering (required).
	From time.Time `json:"from"`
	// To is the end date for filtering (required).
	To time.Time `json:"to"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// FilingsByFormTypeParams represents parameters for filings by form type.
type FilingsByFormTypeParams struct {
	// FormType is the form type to filter by (required).
	FormType string `json:"formType"`
	// From is the start date for filtering (required).
	From time.Time `json:"from"`
	// To is the end date for filtering (required).
	To time.Time `json:"to"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// FilingsBySymbolParams represents parameters for filings by symbol.
type FilingsBySymbolParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// From is the start date for filtering (required).
	From time.Time `json:"from"`
	// To is the end date for filtering (required).
	To time.Time `json:"to"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// FilingsByCIKParams represents parameters for filings by CIK.
type FilingsByCIKParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
	// From is the start date for filtering (required).
	From time.Time `json:"from"`
	// To is the end date for filtering (required).
	To time.Time `json:"to"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// CompanySearchByNameParams represents parameters for company search by name.
type CompanySearchByNameParams struct {
	// Company is the company name to search for (required).
	Company string `json:"company"`
}

// CompanySearchBySymbolParams represents parameters for company search by symbol.
type CompanySearchBySymbolParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// CompanySearchByCIKParams represents parameters for company search by CIK.
type CompanySearchByCIKParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
}

// SECProfileParams represents parameters for SEC profile.
type SECProfileParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// CIK is the Central Index Key (optional).
	CIK *string `json:"cik-A,omitempty"`
}

// IndustryClassificationParams represents parameters for industry classification.
type IndustryClassificationParams struct {
	// IndustryTitle is the industry title to filter by (optional).
	IndustryTitle *string `json:"industryTitle,omitempty"`
	// SICCode is the SIC code to filter by (optional).
	SICCode *string `json:"sicCode,omitempty"`
}

// IndustrySearchParams represents parameters for industry search.
type IndustrySearchParams struct {
	// Symbol is the stock ticker symbol (optional).
	Symbol *string `json:"symbol,omitempty"`
	// CIK is the Central Index Key (optional).
	CIK *string `json:"cik,omitempty"`
	// SICCode is the SIC code (optional).
	SICCode *string `json:"sicCode,omitempty"`
}

// PaginationParams represents basic pagination parameters.
type PaginationParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// SECFilingResponse represents a SEC filing.
type SECFilingResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// FilingDate is the date of filing.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is the date the filing was accepted.
	AcceptedDate string `json:"acceptedDate"`
	// FormType is the type of form filed.
	FormType string `json:"formType"`
	// HasFinancials indicates if financials are included.
	HasFinancials bool `json:"hasFinancials,omitempty"`
	// Link is the URL to the filing.
	Link string `json:"link"`
	// FinalLink is the URL to the final document.
	FinalLink string `json:"finalLink"`
}

// CompanyInfoResponse represents company information.
type CompanyInfoResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the company name.
	Name string `json:"name"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// SICCode is the Standard Industrial Classification code.
	SICCode string `json:"sicCode"`
	// IndustryTitle is the industry description.
	IndustryTitle string `json:"industryTitle"`
	// BusinessAddress is the business address.
	BusinessAddress string `json:"businessAddress"`
	// PhoneNumber is the phone number.
	PhoneNumber string `json:"phoneNumber"`
}

// SECProfileResponse represents a comprehensive SEC company profile.
type SECProfileResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// RegistrantName is the official company name.
	RegistrantName string `json:"registrantName"`
	// SICCode is the Standard Industrial Classification code.
	SICCode string `json:"sicCode"`
	// SICDescription is the SIC description.
	SICDescription string `json:"sicDescription"`
	// SICGroup is the SIC group classification.
	SICGroup string `json:"sicGroup"`
	// ISIN is the International Securities Identification Number.
	ISIN string `json:"isin"`
	// BusinessAddress is the business address.
	BusinessAddress string `json:"businessAddress"`
	// MailingAddress is the mailing address.
	MailingAddress string `json:"mailingAddress"`
	// PhoneNumber is the phone number.
	PhoneNumber string `json:"phoneNumber"`
	// PostalCode is the postal code.
	PostalCode string `json:"postalCode"`
	// City is the city.
	City string `json:"city"`
	// State is the state.
	State string `json:"state"`
	// Country is the country code.
	Country string `json:"country"`
	// Description is the company description.
	Description string `json:"description"`
	// CEO is the CEO name.
	CEO string `json:"ceo"`
	// Website is the company website.
	Website string `json:"website"`
	// Exchange is the stock exchange.
	Exchange string `json:"exchange"`
	// StateLocation is the state location.
	StateLocation string `json:"stateLocation"`
	// StateOfIncorporation is the state of incorporation.
	StateOfIncorporation string `json:"stateOfIncorporation"`
	// FiscalYearEnd is the fiscal year end date.
	FiscalYearEnd string `json:"fiscalYearEnd"`
	// IPODate is the IPO date.
	IPODate string `json:"ipoDate"`
	// Employees is the number of employees.
	Employees string `json:"employees"`
	// SECFilingsURL is the URL to SEC filings.
	SECFilingsURL string `json:"secFilingsUrl"`
	// TaxIdentificationNumber is the tax ID.
	TaxIdentificationNumber string `json:"taxIdentificationNumber"`
	// FiftyTwoWeekRange is the 52-week price range.
	FiftyTwoWeekRange string `json:"fiftyTwoWeekRange"`
	// IsActive indicates if the company is active.
	IsActive bool `json:"isActive"`
	// AssetType is the asset type.
	AssetType string `json:"assetType"`
	// OpenFIGIComposite is the OpenFIGI identifier.
	OpenFIGIComposite string `json:"openFigiComposite"`
	// PriceCurrency is the price currency.
	PriceCurrency string `json:"priceCurrency"`
	// MarketSector is the market sector.
	MarketSector string `json:"marketSector"`
	// SecurityType is the security type.
	SecurityType *string `json:"securityType"`
	// IsETF indicates if this is an ETF.
	IsETF bool `json:"isEtf"`
	// IsADR indicates if this is an ADR.
	IsADR bool `json:"isAdr"`
	// IsFund indicates if this is a fund.
	IsFund bool `json:"isFund"`
}

// IndustryClassificationResponse represents industry classification.
type IndustryClassificationResponse struct {
	// Office is the SEC office classification.
	Office string `json:"office"`
	// SICCode is the Standard Industrial Classification code.
	SICCode string `json:"sicCode"`
	// IndustryTitle is the industry description.
	IndustryTitle string `json:"industryTitle"`
}

// Latest8KFilings retrieves latest 8-K SEC filings.
// Returns significant company events and material changes.
//
// Endpoint: /sec-filings-8k
func (c *Client) Latest8KFilings(params SECFilingsParams) ([]SECFilingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-8k"

	var result []SECFilingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// LatestSECFilings retrieves latest SEC filings with financials.
// Returns recent regulatory filings that include financial data.
//
// Endpoint: /sec-filings-financials
func (c *Client) LatestSECFilings(params SECFilingsParams) ([]SECFilingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-financials"

	var result []SECFilingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SECFilingsByFormType retrieves filings by form type.
// Returns filings filtered by specific form type (10-K, 10-Q, 8-K, etc.).
//
// Endpoint: /sec-filings-search/form-type
func (c *Client) SECFilingsByFormType(params FilingsByFormTypeParams) ([]SECFilingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-search/form-type"

	var result []SECFilingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SECFilingsBySymbol retrieves filings by stock symbol.
// Returns all SEC filings for a specific company.
//
// Endpoint: /sec-filings-search/symbol
func (c *Client) SECFilingsBySymbol(params FilingsBySymbolParams) ([]SECFilingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-search/symbol"

	var result []SECFilingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SECFilingsByCIK retrieves filings by CIK number.
// Returns all SEC filings for a specific Central Index Key.
//
// Endpoint: /sec-filings-search/cik
func (c *Client) SECFilingsByCIK(params FilingsByCIKParams) ([]SECFilingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-search/cik"

	var result []SECFilingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SECFilingsCompanySearchByName searches companies by name.
// Returns company information matching the search term.
//
// Endpoint: /sec-filings-company-search/name
func (c *Client) SECFilingsCompanySearchByName(params CompanySearchByNameParams) ([]CompanyInfoResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-company-search/name"

	var result []CompanyInfoResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SECFilingsCompanySearchBySymbol searches companies by symbol.
// Returns company information for the given ticker symbol.
//
// Endpoint: /sec-filings-company-search/symbol
func (c *Client) SECFilingsCompanySearchBySymbol(params CompanySearchBySymbolParams) ([]CompanyInfoResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-company-search/symbol"

	var result []CompanyInfoResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SECFilingsCompanySearchByCIK searches companies by CIK.
// Returns company information for the given CIK number.
//
// Endpoint: /sec-filings-company-search/cik
func (c *Client) SECFilingsCompanySearchByCIK(params CompanySearchByCIKParams) ([]CompanyInfoResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-filings-company-search/cik"

	var result []CompanyInfoResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SECCompanyProfile retrieves comprehensive company profile.
// Returns detailed SEC filing information and company data.
//
// Endpoint: /sec-profile
func (c *Client) SECCompanyProfile(params SECProfileParams) ([]SECProfileResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sec-profile"

	var result []SECProfileResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndustryClassificationList retrieves industry classifications.
// Returns list of SIC codes and industry titles.
//
// Endpoint: /standard-industrial-classification-list
func (c *Client) IndustryClassificationList(params IndustryClassificationParams) ([]IndustryClassificationResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/standard-industrial-classification-list"

	var result []IndustryClassificationResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndustryClassificationSearch searches industry classifications.
// Returns company industry classification details.
//
// Endpoint: /industry-classification-search
func (c *Client) IndustryClassificationSearch(params IndustrySearchParams) ([]CompanyInfoResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/industry-classification-search"

	var result []CompanyInfoResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AllIndustryClassification retrieves all industry classifications.
// Returns comprehensive list of companies with their industry codes.
//
// Endpoint: /all-industry-classification
func (c *Client) AllIndustryClassification(params PaginationParams) ([]CompanyInfoResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/all-industry-classification"

	var result []CompanyInfoResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}