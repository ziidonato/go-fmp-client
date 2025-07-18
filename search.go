package fmp_client

// StockSymbolSearchParams represents the parameters for searching stock symbols.
type StockSymbolSearchParams struct {
	// Query is the search query for finding stock symbols (required).
	Query string `json:"query"`
	// Limit is the maximum number of results to return (optional).
	Limit *int `json:"limit,omitempty"`
	// Number is an alternative parameter for limiting results (optional).
	Number *int `json:"number,omitempty"`
}

// StockSymbolSearchResponse represents a single stock symbol search result.
type StockSymbolSearchResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the company name.
	Name string `json:"name"`
	// Currency is the trading currency for the stock.
	Currency string `json:"currency"`
	// ExchangeFullName is the full name of the exchange.
	ExchangeFullName string `json:"exchangeFullName"`
	// Exchange is the short name of the exchange.
	Exchange string `json:"exchange"`
}

// StockSymbolSearch searches for stock symbols based on a query string.
// It returns a list of matching stock symbols with their details.
//
// Endpoint: /search-symbol
func (c *Client) StockSymbolSearch(params StockSymbolSearchParams) ([]StockSymbolSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-symbol"

	var result []StockSymbolSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CompanyNameSearchParams represents the parameters for searching companies by name.
type CompanyNameSearchParams struct {
	// Query is the search query for finding companies (required).
	Query string `json:"query"`
	// Limit is the maximum number of results to return (optional).
	Limit *int `json:"limit,omitempty"`
	// Number is an alternative parameter for limiting results (optional).
	Number *int `json:"number,omitempty"`
	// Exchange filters results by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
}

// CompanyNameSearchResponse represents a single company name search result.
type CompanyNameSearchResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the company name.
	Name string `json:"name"`
	// Currency is the trading currency for the stock.
	Currency string `json:"currency"`
	// ExchangeFullName is the full name of the exchange.
	ExchangeFullName string `json:"exchangeFullName"`
	// Exchange is the short name of the exchange.
	Exchange string `json:"exchange"`
}

// CompanyNameSearch searches for companies by name or partial name.
// It returns a list of matching companies with their ticker symbols and details.
//
// Endpoint: /search-name
func (c *Client) CompanyNameSearch(params CompanyNameSearchParams) ([]CompanyNameSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-name"

	var result []CompanyNameSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CIKParams represents the parameters for searching by CIK.
type CIKParams struct {
	// CIK is the Central Index Key to search for (required).
	CIK string `json:"cik"`
}

// CIKResponse represents a single CIK search result.
type CIKResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full name of the company.
	CompanyName string `json:"companyName"`
	// CIK is the Central Index Key with leading zeros.
	CIK string `json:"cik"`
	// ExchangeFullName is the full name of the exchange.
	ExchangeFullName string `json:"exchangeFullName"`
	// Exchange is the short name of the exchange.
	Exchange string `json:"exchange"`
	// Currency is the trading currency for the stock.
	Currency string `json:"currency"`
}

// CIK retrieves company information by Central Index Key (CIK).
// CIK is a unique identifier assigned by the SEC to companies for regulatory filings.
//
// Endpoint: /search-cik
func (c *Client) CIK(params CIKParams) ([]CIKResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-cik"

	var result []CIKResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CUSIPParams represents the parameters for searching by CUSIP.
type CUSIPParams struct {
	// CUSIP is the CUSIP number to search for (required).
	CUSIP string `json:"cusip"`
}

// CUSIPResponse represents a single CUSIP search result.
type CUSIPResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full name of the company.
	CompanyName string `json:"companyName"`
	// CUSIP is the CUSIP identifier.
	CUSIP string `json:"cusip"`
	// MarketCap is the market capitalization in dollars.
	MarketCap int64 `json:"marketCap"`
}

// CUSIP retrieves company information by CUSIP number.
// CUSIP is a 9-character alphanumeric code that identifies securities.
//
// Endpoint: /search-cusip
func (c *Client) CUSIP(params CUSIPParams) ([]CUSIPResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-cusip"

	var result []CUSIPResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ISINParams represents the parameters for searching by ISIN.
type ISINParams struct {
	// ISIN is the International Securities Identification Number to search for (required).
	ISIN string `json:"isin"`
}

// ISINResponse represents a single ISIN search result.
type ISINResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the company name.
	Name string `json:"name"`
	// ISIN is the International Securities Identification Number.
	ISIN string `json:"isin"`
	// MarketCap is the market capitalization in dollars.
	MarketCap int64 `json:"marketCap"`
}

// ISIN retrieves company information by International Securities Identification Number.
// ISIN is a 12-character alphanumeric code that uniquely identifies securities internationally.
//
// Endpoint: /search-isin
func (c *Client) ISIN(params ISINParams) ([]ISINResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-isin"

	var result []ISINResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockScreenerParams represents the parameters for screening stocks.
type StockScreenerParams struct {
	// MarketCapMoreThan filters for stocks with market cap greater than this value (optional).
	MarketCapMoreThan *int64 `json:"marketCapMoreThan,omitempty"`
	// MarketCapLowerThan filters for stocks with market cap less than this value (optional).
	MarketCapLowerThan *int64 `json:"marketCapLowerThan,omitempty"`
	// Sector filters by specific sector (optional).
	Sector *string `json:"sector,omitempty"`
	// Industry filters by specific industry (optional).
	Industry *string `json:"industry,omitempty"`
	// BetaMoreThan filters for stocks with beta greater than this value (optional).
	BetaMoreThan *float64 `json:"betaMoreThan,omitempty"`
	// BetaLowerThan filters for stocks with beta less than this value (optional).
	BetaLowerThan *float64 `json:"betaLowerThan,omitempty"`
	// PriceMoreThan filters for stocks with price greater than this value (optional).
	PriceMoreThan *float64 `json:"priceMoreThan,omitempty"`
	// PriceLowerThan filters for stocks with price less than this value (optional).
	PriceLowerThan *float64 `json:"priceLowerThan,omitempty"`
	// DividendMoreThan filters for stocks with dividend greater than this value (optional).
	DividendMoreThan *float64 `json:"dividendMoreThan,omitempty"`
	// DividendLowerThan filters for stocks with dividend less than this value (optional).
	DividendLowerThan *float64 `json:"dividendLowerThan,omitempty"`
	// VolumeMoreThan filters for stocks with volume greater than this value (optional).
	VolumeMoreThan *int64 `json:"volumeMoreThan,omitempty"`
	// VolumeLowerThan filters for stocks with volume less than this value (optional).
	VolumeLowerThan *int64 `json:"volumeLowerThan,omitempty"`
	// Exchange filters by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
	// Country filters by specific country (optional).
	Country *string `json:"country,omitempty"`
	// IsETF filters for ETFs when true (optional).
	IsETF *bool `json:"isEtf,omitempty"`
	// IsFund filters for funds when true (optional).
	IsFund *bool `json:"isFund,omitempty"`
	// IsActivelyTrading filters for actively trading stocks when true (optional).
	IsActivelyTrading *bool `json:"isActivelyTrading,omitempty"`
	// Limit is the maximum number of results to return (optional).
	Limit *int `json:"limit,omitempty"`
	// IncludeAllShareClasses includes all share classes when true (optional).
	IncludeAllShareClasses *bool `json:"includeAllShareClasses,omitempty"`
}

// StockScreenerResponse represents a single stock screening result.
type StockScreenerResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full name of the company.
	CompanyName string `json:"companyName"`
	// MarketCap is the market capitalization in dollars.
	MarketCap int64 `json:"marketCap"`
	// Sector is the business sector of the company.
	Sector string `json:"sector"`
	// Industry is the specific industry within the sector.
	Industry string `json:"industry"`
	// Beta is the stock's beta coefficient.
	Beta float64 `json:"beta"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// LastAnnualDividend is the most recent annual dividend per share.
	LastAnnualDividend float64 `json:"lastAnnualDividend"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
	// Exchange is the full name of the exchange.
	Exchange string `json:"exchange"`
	// ExchangeShortName is the short name of the exchange.
	ExchangeShortName string `json:"exchangeShortName"`
	// Country is the country where the company is based.
	Country string `json:"country"`
	// IsETF indicates if this is an ETF.
	IsETF bool `json:"isEtf"`
	// IsFund indicates if this is a fund.
	IsFund bool `json:"isFund"`
	// IsActivelyTrading indicates if the stock is actively trading.
	IsActivelyTrading bool `json:"isActivelyTrading"`
}

// StockScreener filters stocks based on various criteria including market cap, price, volume, beta, sector, and more.
// This is useful for finding stocks that match specific investment criteria.
//
// Endpoint: /company-screener
func (c *Client) StockScreener(params StockScreenerParams) ([]StockScreenerResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/company-screener"

	var result []StockScreenerResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ExchangeVariantsParams represents the parameters for searching exchange variants.
type ExchangeVariantsParams struct {
	// Symbol is the stock ticker symbol to search for across exchanges (required).
	Symbol string `json:"symbol"`
}

// ExchangeVariantsResponse represents a single exchange variant result.
type ExchangeVariantsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// Beta is the stock's beta coefficient.
	Beta float64 `json:"beta"`
	// VolAvg is the average trading volume.
	VolAvg int64 `json:"volAvg"`
	// MktCap is the market capitalization in dollars.
	MktCap int64 `json:"mktCap"`
	// LastDiv is the last dividend amount.
	LastDiv float64 `json:"lastDiv"`
	// Range is the 52-week price range.
	Range string `json:"range"`
	// Changes is the price change.
	Changes float64 `json:"changes"`
	// CompanyName is the full name of the company.
	CompanyName string `json:"companyName"`
	// Currency is the trading currency.
	Currency string `json:"currency"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// ISIN is the International Securities Identification Number.
	ISIN string `json:"isin"`
	// CUSIP is the CUSIP identifier.
	CUSIP string `json:"cusip"`
	// Exchange is the full name of the exchange.
	Exchange string `json:"exchange"`
	// ExchangeShortName is the short name of the exchange.
	ExchangeShortName string `json:"exchangeShortName"`
	// Industry is the company's industry.
	Industry string `json:"industry"`
	// Website is the company's website URL.
	Website string `json:"website"`
	// Description is a brief description of the company.
	Description string `json:"description"`
	// CEO is the name of the CEO.
	CEO string `json:"ceo"`
	// Sector is the business sector.
	Sector string `json:"sector"`
	// Country is the country where the company is based.
	Country string `json:"country"`
	// FullTimeEmployees is the number of full-time employees.
	FullTimeEmployees string `json:"fullTimeEmployees"`
	// Phone is the company's phone number.
	Phone string `json:"phone"`
	// Address is the company's street address.
	Address string `json:"address"`
	// City is the city where the company is located.
	City string `json:"city"`
	// State is the state/province where the company is located.
	State string `json:"state"`
	// ZIP is the postal code.
	ZIP string `json:"zip"`
	// DCFDiff is the DCF difference.
	DCFDiff float64 `json:"dcfDiff"`
	// DCF is the discounted cash flow value.
	DCF float64 `json:"dcf"`
	// Image is the URL to the company's logo image.
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

// ExchangeVariants finds all exchanges where a given stock symbol is listed.
// This allows users to identify all trading venues for a security.
//
// Endpoint: /search-exchange-variants
func (c *Client) ExchangeVariants(params ExchangeVariantsParams) ([]ExchangeVariantsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/search-exchange-variants"

	var result []ExchangeVariantsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}
