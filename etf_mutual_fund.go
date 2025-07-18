package fmp_client

import "time"

// ETFParams represents parameters for ETF endpoints.
type ETFParams struct {
	// Symbol is the ETF or mutual fund ticker symbol (required).
	Symbol string `json:"symbol"`
}

// ETFSearchParams represents parameters for ETF search endpoints.
type ETFSearchParams struct {
	// Name is the search query for ETF name (optional).
	Name *string `json:"name,omitempty"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// ETFHoldingResponse represents ETF holding data.
type ETFHoldingResponse struct {
	// Symbol is the ETF ticker symbol.
	Symbol string `json:"symbol"`
	// Asset is the held asset ticker.
	Asset string `json:"asset"`
	// Name is the asset name.
	Name string `json:"name"`
	// ISIN is the International Securities Identification Number.
	ISIN string `json:"isin"`
	// SecurityCUSIP is the CUSIP identifier.
	SecurityCUSIP string `json:"securityCusip"`
	// SharesNumber is the number of shares held.
	SharesNumber int64 `json:"sharesNumber"`
	// WeightPercentage is the portfolio weight percentage.
	WeightPercentage float64 `json:"weightPercentage"`
	// MarketValue is the market value of the holding.
	MarketValue float64 `json:"marketValue"`
	// UpdatedAt is when the data was last updated.
	UpdatedAt string `json:"updatedAt"`
	// Updated is the update timestamp.
	Updated string `json:"updated"`
}

// ETFInfoResponse represents comprehensive ETF information.
type ETFInfoResponse struct {
	// Symbol is the ETF ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the ETF name.
	Name string `json:"name"`
	// Description is the ETF description.
	Description string `json:"description"`
	// ISIN is the International Securities Identification Number.
	ISIN string `json:"isin"`
	// AssetClass is the asset class.
	AssetClass string `json:"assetClass"`
	// SecurityCUSIP is the CUSIP identifier.
	SecurityCUSIP string `json:"securityCusip"`
	// Domicile is the fund domicile country.
	Domicile string `json:"domicile"`
	// Website is the fund website.
	Website string `json:"website"`
	// ETFCompany is the ETF provider company.
	ETFCompany string `json:"etfCompany"`
	// ExpenseRatio is the expense ratio.
	ExpenseRatio float64 `json:"expenseRatio"`
	// AssetsUnderManagement is the AUM.
	AssetsUnderManagement int64 `json:"assetsUnderManagement"`
	// AvgVolume is the average trading volume.
	AvgVolume int64 `json:"avgVolume"`
	// InceptionDate is the fund inception date.
	InceptionDate string `json:"inceptionDate"`
	// NAV is the net asset value.
	NAV float64 `json:"nav"`
	// NAVCurrency is the NAV currency.
	NAVCurrency string `json:"navCurrency"`
	// HoldingsCount is the number of holdings.
	HoldingsCount int `json:"holdingsCount"`
	// UpdatedAt is when the data was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
	// SectorsList is the list of sector exposures.
	SectorsList []SectorExposure `json:"sectorsList"`
}

// SectorExposure represents sector allocation data.
type SectorExposure struct {
	// Industry is the industry/sector name.
	Industry string `json:"industry"`
	// Exposure is the percentage exposure.
	Exposure float64 `json:"exposure"`
}

// CountryWeightingResponse represents country allocation data.
type CountryWeightingResponse struct {
	// Country is the country name.
	Country string `json:"country"`
	// WeightPercentage is the allocation percentage.
	WeightPercentage string `json:"weightPercentage"`
}

// AssetExposureResponse represents ETF exposure to an asset.
type AssetExposureResponse struct {
	// ETFSymbol is the ETF ticker symbol.
	ETFSymbol string `json:"etfSymbol"`
	// SharesNumber is the number of shares held.
	SharesNumber int64 `json:"sharesNumber"`
	// WeightPercentage is the portfolio weight percentage.
	WeightPercentage float64 `json:"weightPercentage"`
	// MarketValue is the market value of the holding.
	MarketValue float64 `json:"marketValue"`
	// LastUpdated is when the data was last updated.
	LastUpdated string `json:"lastUpdated"`
}

// ETFSearchResponse represents ETF search results.
type ETFSearchResponse struct {
	// Symbol is the ETF ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the ETF name.
	Name string `json:"name"`
	// Price is the current price.
	Price float64 `json:"price"`
	// Exchange is the exchange code.
	Exchange string `json:"exchange"`
	// ExchangeShortName is the short exchange name.
	ExchangeShortName string `json:"exchangeShortName"`
	// Type is the security type.
	Type string `json:"type"`
}

// ETFSharesFloatResponse represents ETF shares float data.
type ETFSharesFloatResponse struct {
	// Symbol is the ETF ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the data date.
	Date string `json:"date"`
	// FloatShares is the number of float shares.
	FloatShares int64 `json:"floatShares"`
	// OutstandingShares is the number of outstanding shares.
	OutstandingShares int64 `json:"outstandingShares"`
}

// ETFHoldings retrieves holdings for an ETF or mutual fund.
// Returns detailed portfolio composition and weights.
//
// Endpoint: /etf/holdings
func (c *Client) ETFHoldings(params ETFParams) ([]ETFHoldingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/etf/holdings"

	var result []ETFHoldingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ETFInfo retrieves comprehensive ETF or mutual fund information.
// Returns fund details, expense ratio, AUM, and sector breakdown.
//
// Endpoint: /etf/info
func (c *Client) ETFInfo(params ETFParams) ([]ETFInfoResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/etf/info"

	var result []ETFInfoResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ETFCountryWeightings retrieves country allocation for an ETF.
// Returns percentage allocation across different countries.
//
// Endpoint: /etf/country-weightings
func (c *Client) ETFCountryWeightings(params ETFParams) ([]CountryWeightingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/etf/country-weightings"

	var result []CountryWeightingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ETFAssetExposure retrieves ETFs holding a specific asset.
// Returns list of ETFs with exposure to the given stock.
//
// Endpoint: /etf/asset-exposure
func (c *Client) ETFAssetExposure(params ETFParams) ([]AssetExposureResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/etf/asset-exposure"

	var result []AssetExposureResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchETFs searches for ETFs by name.
// Returns matching ETFs with basic information.
//
// Endpoint: /etf/search
func (c *Client) SearchETFs(params ETFSearchParams) ([]ETFSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/etf/search"

	var result []ETFSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ETFSharesFloat retrieves shares float data for an ETF.
// Returns float and outstanding share information.
//
// Endpoint: /etf/shares-float
func (c *Client) ETFSharesFloat(params ETFParams) ([]ETFSharesFloatResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/etf/shares-float"

	var result []ETFSharesFloatResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}