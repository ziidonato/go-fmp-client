package fmp_client

// BulkProfileParams represents parameters for bulk profile request.
type BulkProfileParams struct {
	// Part is the part number for pagination (required).
	Part string `json:"part"`
}

// BulkProfileResponse represents bulk company profile data.
// Note: This endpoint returns CSV data, but the structure represents a single record.
type BulkProfileResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Price is the current stock price.
	Price string `json:"price"`
	// MarketCap is the market capitalization.
	MarketCap string `json:"marketCap"`
	// Beta is the stock's beta value.
	Beta string `json:"beta"`
	// LastDividend is the last dividend amount.
	LastDividend string `json:"lastDividend"`
	// Range is the price range.
	Range string `json:"range"`
	// Change is the price change.
	Change string `json:"change"`
	// ChangePercentage is the percentage change.
	ChangePercentage string `json:"changePercentage"`
	// Volume is the trading volume.
	Volume string `json:"volume"`
	// AverageVolume is the average trading volume.
	AverageVolume string `json:"averageVolume"`
	// CompanyName is the company name.
	CompanyName string `json:"companyName"`
	// Currency is the currency code.
	Currency string `json:"currency"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// ISIN is the International Securities Identification Number.
	ISIN string `json:"isin"`
	// CUSIP is the CUSIP number.
	CUSIP string `json:"cusip"`
	// ExchangeFullName is the full exchange name.
	ExchangeFullName string `json:"exchangeFullName"`
	// Exchange is the exchange code.
	Exchange string `json:"exchange"`
	// Industry is the industry classification.
	Industry string `json:"industry"`
	// Website is the company website.
	Website string `json:"website"`
	// Description is the company description.
	Description string `json:"description"`
	// CEO is the CEO name.
	CEO string `json:"ceo"`
	// Sector is the sector classification.
	Sector string `json:"sector"`
	// Country is the country code.
	Country string `json:"country"`
	// FullTimeEmployees is the number of employees.
	FullTimeEmployees string `json:"fullTimeEmployees"`
	// Phone is the phone number.
	Phone string `json:"phone"`
	// Address is the street address.
	Address string `json:"address"`
	// City is the city.
	City string `json:"city"`
	// State is the state.
	State string `json:"state"`
	// Zip is the zip code.
	Zip string `json:"zip"`
	// Image is the company logo URL.
	Image string `json:"image"`
	// IPODate is the IPO date.
	IPODate string `json:"ipoDate"`
	// DefaultImage is the default image URL.
	DefaultImage string `json:"defaultImage"`
	// IsETF indicates if this is an ETF.
	IsETF string `json:"isEtf"`
	// IsActivelyTrading indicates if actively trading.
	IsActivelyTrading string `json:"isActivelyTrading"`
	// IsADR indicates if this is an ADR.
	IsADR string `json:"isAdr"`
	// IsFund indicates if this is a fund.
	IsFund string `json:"isFund"`
}

// BulkRatingResponse represents bulk rating data.
// Note: This endpoint returns CSV data, but the structure represents a single record.
type BulkRatingResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the rating date.
	Date string `json:"date"`
	// Rating is the overall rating.
	Rating string `json:"rating"`
	// DiscountedCashFlowScore is the DCF score.
	DiscountedCashFlowScore string `json:"discountedCashFlowScore"`
	// ReturnOnEquityScore is the ROE score.
	ReturnOnEquityScore string `json:"returnOnEquityScore"`
	// ReturnOnAssetsScore is the ROA score.
	ReturnOnAssetsScore string `json:"returnOnAssetsScore"`
	// DebtToEquityScore is the D/E score.
	DebtToEquityScore string `json:"debtToEquityScore"`
	// PriceToEarningsScore is the P/E score.
	PriceToEarningsScore string `json:"priceToEarningsScore"`
	// PriceToBookScore is the P/B score.
	PriceToBookScore string `json:"priceToBookScore"`
}

// BulkDCFResponse represents bulk DCF valuation data.
// Note: This endpoint returns CSV data, but the structure represents a single record.
type BulkDCFResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the valuation date.
	Date string `json:"date"`
	// DCF is the discounted cash flow value.
	DCF string `json:"dcf"`
	// Price is the stock price.
	Price string `json:"price"`
}

// BulkRatiosResponse represents bulk financial ratios data.
// Note: This endpoint returns CSV data, but the structure represents a single record.
type BulkRatiosResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Period is the reporting period.
	Period string `json:"period"`
	// CurrentRatio is the current ratio.
	CurrentRatio string `json:"currentRatio"`
	// QuickRatio is the quick ratio.
	QuickRatio string `json:"quickRatio"`
	// CashRatio is the cash ratio.
	CashRatio string `json:"cashRatio"`
	// DebtToEquity is the debt-to-equity ratio.
	DebtToEquity string `json:"debtToEquity"`
	// GrossMargin is the gross margin.
	GrossMargin string `json:"grossMargin"`
	// OperatingMargin is the operating margin.
	OperatingMargin string `json:"operatingMargin"`
	// NetMargin is the net margin.
	NetMargin string `json:"netMargin"`
	// ReturnOnAssets is the ROA.
	ReturnOnAssets string `json:"returnOnAssets"`
	// ReturnOnEquity is the ROE.
	ReturnOnEquity string `json:"returnOnEquity"`
	// PriceToEarnings is the P/E ratio.
	PriceToEarnings string `json:"priceToEarnings"`
	// PriceToBook is the P/B ratio.
	PriceToBook string `json:"priceToBook"`
	// PriceToSales is the P/S ratio.
	PriceToSales string `json:"priceToSales"`
}

// BulkMetricsResponse represents bulk key metrics data.
// Note: This endpoint returns CSV data, but the structure represents a single record.
type BulkMetricsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Period is the reporting period.
	Period string `json:"period"`
	// RevenuePerShare is revenue per share.
	RevenuePerShare string `json:"revenuePerShare"`
	// NetIncomePerShare is net income per share.
	NetIncomePerShare string `json:"netIncomePerShare"`
	// OperatingCashFlowPerShare is operating cash flow per share.
	OperatingCashFlowPerShare string `json:"operatingCashFlowPerShare"`
	// FreeCashFlowPerShare is free cash flow per share.
	FreeCashFlowPerShare string `json:"freeCashFlowPerShare"`
	// CashPerShare is cash per share.
	CashPerShare string `json:"cashPerShare"`
	// BookValuePerShare is book value per share.
	BookValuePerShare string `json:"bookValuePerShare"`
	// TangibleBookValuePerShare is tangible book value per share.
	TangibleBookValuePerShare string `json:"tangibleBookValuePerShare"`
	// ShareholdersEquityPerShare is shareholders equity per share.
	ShareholdersEquityPerShare string `json:"shareholdersEquityPerShare"`
	// InterestDebtPerShare is interest debt per share.
	InterestDebtPerShare string `json:"interestDebtPerShare"`
	// MarketCap is market capitalization.
	MarketCap string `json:"marketCap"`
	// EnterpriseValue is enterprise value.
	EnterpriseValue string `json:"enterpriseValue"`
	// PERatio is the P/E ratio.
	PERatio string `json:"peRatio"`
	// PSRatio is the P/S ratio.
	PSRatio string `json:"psRatio"`
	// PBRatio is the P/B ratio.
	PBRatio string `json:"pbRatio"`
}

// CompanyProfileBulk retrieves company profiles in bulk.
// Returns comprehensive company data as CSV format.
//
// Endpoint: /profile-bulk
func (c *Client) CompanyProfileBulk(params BulkProfileParams) ([]BulkProfileResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/profile-bulk"

	var result []BulkProfileResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockRatingBulk retrieves stock ratings in bulk.
// Returns financial ratings and scores as CSV format.
//
// Endpoint: /rating-bulk
func (c *Client) StockRatingBulk() ([]BulkRatingResponse, error) {
	pathName := "/rating-bulk"

	var result []BulkRatingResponse
	err := c.doRequest(pathName, nil, &result)
	return result, err
}

// DCFValuationsBulk retrieves DCF valuations in bulk.
// Returns discounted cash flow values as CSV format.
//
// Endpoint: /dcf-bulk
func (c *Client) DCFValuationsBulk() ([]BulkDCFResponse, error) {
	pathName := "/dcf-bulk"

	var result []BulkDCFResponse
	err := c.doRequest(pathName, nil, &result)
	return result, err
}

// FinancialRatiosBulk retrieves financial ratios in bulk.
// Returns comprehensive ratio analysis as CSV format.
//
// Endpoint: /ratio-bulk
func (c *Client) FinancialRatiosBulk() ([]BulkRatiosResponse, error) {
	pathName := "/ratio-bulk"

	var result []BulkRatiosResponse
	err := c.doRequest(pathName, nil, &result)
	return result, err
}

// KeyMetricsBulk retrieves key metrics in bulk.
// Returns per-share metrics and valuations as CSV format.
//
// Endpoint: /metrics-bulk
func (c *Client) KeyMetricsBulk() ([]BulkMetricsResponse, error) {
	pathName := "/metrics-bulk"

	var result []BulkMetricsResponse
	err := c.doRequest(pathName, nil, &result)
	return result, err
}