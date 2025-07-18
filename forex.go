package fmp_client

import "time"

// ForexQuoteParams represents parameters for forex quote endpoints.
type ForexQuoteParams struct {
	// Symbol is the currency pair symbol (required).
	Symbol string `json:"symbol"`
}

// ForexHistoricalParams represents parameters for historical forex data.
type ForexHistoricalParams struct {
	// Symbol is the currency pair symbol (required).
	Symbol string `json:"symbol"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// ForexBatchQuotesParams represents parameters for batch forex quotes.
type ForexBatchQuotesParams struct {
	// Short indicates whether to return short form data (optional).
	Short *bool `json:"short,omitempty"`
}

// ForexListResponse represents a forex currency pair.
type ForexListResponse struct {
	// Symbol is the currency pair symbol.
	Symbol string `json:"symbol"`
	// FromCurrency is the base currency code.
	FromCurrency string `json:"fromCurrency"`
	// ToCurrency is the quote currency code.
	ToCurrency string `json:"toCurrency"`
	// FromName is the base currency name.
	FromName string `json:"fromName"`
	// ToName is the quote currency name.
	ToName string `json:"toName"`
}

// ForexQuoteResponse represents a forex quote.
type ForexQuoteResponse struct {
	// Symbol is the currency pair symbol.
	Symbol string `json:"symbol"`
	// Name is the currency pair name.
	Name string `json:"name"`
	// Price is the current exchange rate.
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
	// MarketCap is typically null for forex.
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

// ForexQuoteShortResponse represents a short forex quote.
type ForexQuoteShortResponse struct {
	// Symbol is the currency pair symbol.
	Symbol string `json:"symbol"`
	// Price is the current exchange rate.
	Price float64 `json:"price"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// ForexHistoricalLightResponse represents simplified historical data.
type ForexHistoricalLightResponse struct {
	// Symbol is the currency pair symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Price is the closing exchange rate.
	Price float64 `json:"price"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// ForexHistoricalFullResponse represents full historical data.
type ForexHistoricalFullResponse struct {
	// Symbol is the currency pair symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Open is the opening exchange rate.
	Open float64 `json:"open"`
	// High is the highest rate during the period.
	High float64 `json:"high"`
	// Low is the lowest rate during the period.
	Low float64 `json:"low"`
	// Close is the closing exchange rate.
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

// ForexIntradayResponse represents intraday forex data.
type ForexIntradayResponse struct {
	// Date is the date and time of the data point.
	Date string `json:"date"`
	// Open is the opening rate for the interval.
	Open float64 `json:"open"`
	// Low is the lowest rate during the interval.
	Low float64 `json:"low"`
	// High is the highest rate during the interval.
	High float64 `json:"high"`
	// Close is the closing rate for the interval.
	Close float64 `json:"close"`
	// Volume is the trading volume for the interval.
	Volume int64 `json:"volume"`
}

// ForexList retrieves all available forex currency pairs.
// Returns list of currency pairs traded on the forex market.
//
// Endpoint: /forex-list
func (c *Client) ForexList() ([]ForexListResponse, error) {
	pathName := "/forex-list"

	var result []ForexListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// ForexQuote retrieves real-time quote for a specific currency pair.
// Returns current exchange rate and market data.
//
// Endpoint: /quote
func (c *Client) ForexQuote(params ForexQuoteParams) ([]ForexQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote"

	var result []ForexQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexQuoteShort retrieves short form quote for a currency pair.
// Returns essential exchange rate and volume data.
//
// Endpoint: /quote-short
func (c *Client) ForexQuoteShort(params ForexQuoteParams) ([]ForexQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote-short"

	var result []ForexQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BatchForexQuotes retrieves quotes for all forex pairs.
// Returns batch quotes for currency pairs in a single request.
//
// Endpoint: /batch-forex-quotes
func (c *Client) BatchForexQuotes(params ForexBatchQuotesParams) ([]ForexQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-forex-quotes"

	var result []ForexQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexHistoricalLight retrieves simplified historical forex data.
// Returns end-of-day exchange rates and volume.
//
// Endpoint: /historical-price-eod/light
func (c *Client) ForexHistoricalLight(params ForexHistoricalParams) ([]ForexHistoricalLightResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/light"

	var result []ForexHistoricalLightResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexHistoricalFull retrieves full historical forex data.
// Returns OHLC exchange rates, volume, and changes.
//
// Endpoint: /historical-price-eod/full
func (c *Client) ForexHistoricalFull(params ForexHistoricalParams) ([]ForexHistoricalFullResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/full"

	var result []ForexHistoricalFullResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexChart1Min retrieves 1-minute interval forex data.
// Returns intraday exchange rate movements for detailed analysis.
//
// Endpoint: /historical-chart/1min
func (c *Client) ForexChart1Min(params ForexHistoricalParams) ([]ForexIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1min"

	var result []ForexIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexChart5Min retrieves 5-minute interval forex data.
// Returns short-term exchange rate trends for trading analysis.
//
// Endpoint: /historical-chart/5min
func (c *Client) ForexChart5Min(params ForexHistoricalParams) ([]ForexIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/5min"

	var result []ForexIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexChart1Hour retrieves hourly forex data.
// Returns hourly exchange rate movements for mid-term analysis.
//
// Endpoint: /historical-chart/1hour
func (c *Client) ForexChart1Hour(params ForexHistoricalParams) ([]ForexIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1hour"

	var result []ForexIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}