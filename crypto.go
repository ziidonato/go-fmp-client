package fmp_client

import "time"

// CryptoQuoteParams represents parameters for crypto quote endpoints.
type CryptoQuoteParams struct {
	// Symbol is the cryptocurrency symbol (required).
	Symbol string `json:"symbol"`
}

// CryptoHistoricalParams represents parameters for historical crypto data.
type CryptoHistoricalParams struct {
	// Symbol is the cryptocurrency symbol (required).
	Symbol string `json:"symbol"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// CryptoBatchQuotesParams represents parameters for batch crypto quotes.
type CryptoBatchQuotesParams struct {
	// Short indicates whether to return short form data (optional).
	Short *bool `json:"short,omitempty"`
}

// CryptoListResponse represents a cryptocurrency listing.
type CryptoListResponse struct {
	// Symbol is the cryptocurrency symbol.
	Symbol string `json:"symbol"`
	// Name is the cryptocurrency name.
	Name string `json:"name"`
	// Exchange is the exchange where traded.
	Exchange string `json:"exchange"`
	// ICODate is the initial coin offering date.
	ICODate string `json:"icoDate"`
	// CirculatingSupply is the circulating supply.
	CirculatingSupply float64 `json:"circulatingSupply"`
	// TotalSupply is the total supply.
	TotalSupply *float64 `json:"totalSupply"`
}

// CryptoQuoteResponse represents a cryptocurrency quote.
type CryptoQuoteResponse struct {
	// Symbol is the cryptocurrency symbol.
	Symbol string `json:"symbol"`
	// Name is the cryptocurrency name.
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
	// MarketCap is the market capitalization.
	MarketCap int64 `json:"marketCap"`
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

// CryptoQuoteShortResponse represents a short crypto quote.
type CryptoQuoteShortResponse struct {
	// Symbol is the cryptocurrency symbol.
	Symbol string `json:"symbol"`
	// Price is the current price.
	Price float64 `json:"price"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// CryptoHistoricalLightResponse represents simplified historical data.
type CryptoHistoricalLightResponse struct {
	// Symbol is the cryptocurrency symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Price is the closing price.
	Price float64 `json:"price"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// CryptoHistoricalFullResponse represents full historical data.
type CryptoHistoricalFullResponse struct {
	// Symbol is the cryptocurrency symbol.
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

// CryptoIntradayResponse represents intraday crypto data.
type CryptoIntradayResponse struct {
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
	Volume float64 `json:"volume"`
}

// CryptocurrencyList retrieves all available cryptocurrency symbols.
// Returns list of digital assets traded on exchanges worldwide.
//
// Endpoint: /cryptocurrency-list
func (c *Client) CryptocurrencyList() ([]CryptoListResponse, error) {
	pathName := "/cryptocurrency-list"

	var result []CryptoListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// CryptoQuote retrieves real-time quote for a specific cryptocurrency.
// Returns current price, changes, and market data.
//
// Endpoint: /quote
func (c *Client) CryptoQuote(params CryptoQuoteParams) ([]CryptoQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote"

	var result []CryptoQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoQuoteShort retrieves short form quote for a cryptocurrency.
// Returns essential price and volume data.
//
// Endpoint: /quote-short
func (c *Client) CryptoQuoteShort(params CryptoQuoteParams) ([]CryptoQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote-short"

	var result []CryptoQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AllCryptocurrenciesQuotes retrieves quotes for all cryptocurrencies.
// Returns batch quotes for digital assets in a single request.
//
// Endpoint: /batch-crypto-quotes
func (c *Client) AllCryptocurrenciesQuotes(params CryptoBatchQuotesParams) ([]CryptoQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-crypto-quotes"

	var result []CryptoQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoHistoricalLight retrieves simplified historical cryptocurrency prices.
// Returns end-of-day prices and volume.
//
// Endpoint: /historical-price-eod/light
func (c *Client) CryptoHistoricalLight(params CryptoHistoricalParams) ([]CryptoHistoricalLightResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/light"

	var result []CryptoHistoricalLightResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoHistoricalFull retrieves full historical cryptocurrency data.
// Returns OHLC, volume, and price changes.
//
// Endpoint: /historical-price-eod/full
func (c *Client) CryptoHistoricalFull(params CryptoHistoricalParams) ([]CryptoHistoricalFullResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/full"

	var result []CryptoHistoricalFullResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoChart1Min retrieves 1-minute interval cryptocurrency data.
// Returns intraday price movements for detailed analysis.
//
// Endpoint: /historical-chart/1min
func (c *Client) CryptoChart1Min(params CryptoHistoricalParams) ([]CryptoIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1min"

	var result []CryptoIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoChart5Min retrieves 5-minute interval cryptocurrency data.
// Returns short-term price trends for trading analysis.
//
// Endpoint: /historical-chart/5min
func (c *Client) CryptoChart5Min(params CryptoHistoricalParams) ([]CryptoIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/5min"

	var result []CryptoIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoChart1Hour retrieves hourly cryptocurrency data.
// Returns hourly price movements for mid-term analysis.
//
// Endpoint: /historical-chart/1hour
func (c *Client) CryptoChart1Hour(params CryptoHistoricalParams) ([]CryptoIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1hour"

	var result []CryptoIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}