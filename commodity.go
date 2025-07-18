package fmp_client

import "time"

// CommodityQuoteParams represents parameters for commodity quote endpoints.
type CommodityQuoteParams struct {
	// Symbol is the commodity ticker symbol (required).
	Symbol string `json:"symbol"`
}

// CommodityHistoricalParams represents parameters for historical commodity data.
type CommodityHistoricalParams struct {
	// Symbol is the commodity ticker symbol (required).
	Symbol string `json:"symbol"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// CommodityBatchQuotesParams represents parameters for batch commodity quotes.
type CommodityBatchQuotesParams struct {
	// Short indicates whether to return short form data (optional).
	Short *bool `json:"short,omitempty"`
}

// CommodityListResponse represents a commodity listing.
type CommodityListResponse struct {
	// Symbol is the commodity ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the commodity name.
	Name string `json:"name"`
	// Exchange is the exchange where traded.
	Exchange *string `json:"exchange"`
	// TradeMonth is the contract month.
	TradeMonth string `json:"tradeMonth"`
	// Currency is the trading currency.
	Currency string `json:"currency"`
}

// CommodityQuoteResponse represents a commodity quote.
type CommodityQuoteResponse struct {
	// Symbol is the commodity ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the commodity name.
	Name string `json:"name"`
	// Price is the current price.
	Price float64 `json:"price"`
	// ChangePercentage is the percentage change.
	ChangePercentage float64 `json:"changePercentage"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
	// DayLow is the lowest price of the day.
	DayLow float64 `json:"dayLow"`
	// DayHigh is the highest price of the day.
	DayHigh float64 `json:"dayHigh"`
	// YearHigh is the 52-week high.
	YearHigh float64 `json:"yearHigh"`
	// YearLow is the 52-week low.
	YearLow float64 `json:"yearLow"`
	// MarketCap is typically null for commodities.
	MarketCap *int64 `json:"marketCap"`
	// PriceAvg50 is the 50-day moving average.
	PriceAvg50 float64 `json:"priceAvg50"`
	// PriceAvg200 is the 200-day moving average.
	PriceAvg200 float64 `json:"priceAvg200"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
	// Open is the opening price.
	Open float64 `json:"open"`
	// PreviousClose is the previous closing price.
	PreviousClose float64 `json:"previousClose"`
	// Timestamp is the Unix timestamp.
	Timestamp int64 `json:"timestamp"`
}

// CommodityQuoteShortResponse represents a short commodity quote.
type CommodityQuoteShortResponse struct {
	// Symbol is the commodity ticker symbol.
	Symbol string `json:"symbol"`
	// Price is the current price.
	Price float64 `json:"price"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// CommodityHistoricalLightResponse represents simplified historical data.
type CommodityHistoricalLightResponse struct {
	// Symbol is the commodity ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Price is the closing price.
	Price float64 `json:"price"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// CommodityHistoricalFullResponse represents full historical data.
type CommodityHistoricalFullResponse struct {
	// Symbol is the commodity ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Open is the opening price.
	Open float64 `json:"open"`
	// High is the highest price during the period.
	High float64 `json:"high"`
	// Low is the lowest price during the period.
	Low float64 `json:"low"`
	// Close is the closing price.
	Close float64 `json:"close"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// ChangePercent is the price change percentage.
	ChangePercent float64 `json:"changePercent"`
	// VWAP is the volume-weighted average price.
	VWAP float64 `json:"vwap"`
}

// CommodityIntradayResponse represents intraday commodity data.
type CommodityIntradayResponse struct {
	// Date is the date and time of the data point.
	Date string `json:"date"`
	// Open is the opening price for the interval.
	Open float64 `json:"open"`
	// Low is the lowest price during the interval.
	Low float64 `json:"low"`
	// High is the highest price during the interval.
	High float64 `json:"high"`
	// Close is the closing price for the interval.
	Close float64 `json:"close"`
	// Volume is the trading volume for the interval.
	Volume int64 `json:"volume"`
}

// CommoditiesList retrieves all available commodity symbols.
// Returns list of tradable commodities including energy, metals, and agriculture.
//
// Endpoint: /commodities-list
func (c *Client) CommoditiesList() ([]CommodityListResponse, error) {
	pathName := "/commodities-list"

	var result []CommodityListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// CommodityQuote retrieves real-time quote for a specific commodity.
// Returns current price, changes, and trading information.
//
// Endpoint: /quote
func (c *Client) CommodityQuote(params CommodityQuoteParams) ([]CommodityQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote"

	var result []CommodityQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CommodityQuoteShort retrieves short form quote for a commodity.
// Returns essential price and volume data.
//
// Endpoint: /quote-short
func (c *Client) CommodityQuoteShort(params CommodityQuoteParams) ([]CommodityQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote-short"

	var result []CommodityQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AllCommoditiesQuotes retrieves quotes for all commodities.
// Returns batch quotes for energy, metals, and agricultural products.
//
// Endpoint: /batch-commodity-quotes
func (c *Client) AllCommoditiesQuotes(params CommodityBatchQuotesParams) ([]CommodityQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-commodity-quotes"

	var result []CommodityQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CommodityHistoricalLight retrieves simplified historical commodity prices.
// Returns end-of-day prices and volume.
//
// Endpoint: /historical-price-eod/light
func (c *Client) CommodityHistoricalLight(params CommodityHistoricalParams) ([]CommodityHistoricalLightResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/light"

	var result []CommodityHistoricalLightResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CommodityHistoricalFull retrieves full historical commodity data.
// Returns OHLC, volume, and price changes.
//
// Endpoint: /historical-price-eod/full
func (c *Client) CommodityHistoricalFull(params CommodityHistoricalParams) ([]CommodityHistoricalFullResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/full"

	var result []CommodityHistoricalFullResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CommodityChart1Min retrieves 1-minute interval commodity data.
// Returns intraday price movements for detailed analysis.
//
// Endpoint: /historical-chart/1min
func (c *Client) CommodityChart1Min(params CommodityHistoricalParams) ([]CommodityIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1min"

	var result []CommodityIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CommodityChart5Min retrieves 5-minute interval commodity data.
// Returns short-term price trends for trading analysis.
//
// Endpoint: /historical-chart/5min
func (c *Client) CommodityChart5Min(params CommodityHistoricalParams) ([]CommodityIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/5min"

	var result []CommodityIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CommodityChart1Hour retrieves hourly commodity data.
// Returns hourly price movements for mid-term analysis.
//
// Endpoint: /historical-chart/1hour
func (c *Client) CommodityChart1Hour(params CommodityHistoricalParams) ([]CommodityIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1hour"

	var result []CommodityIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}