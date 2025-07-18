package fmp_client

import "time"

// IndexQuoteParams represents parameters for index quote endpoints.
type IndexQuoteParams struct {
	// Symbol is the index symbol (required).
	Symbol string `json:"symbol"`
}

// IndexHistoricalParams represents parameters for historical index data.
type IndexHistoricalParams struct {
	// Symbol is the index symbol (required).
	Symbol string `json:"symbol"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// IndexBatchQuotesParams represents parameters for batch index quotes.
type IndexBatchQuotesParams struct {
	// Short indicates whether to return short form data (optional).
	Short *bool `json:"short,omitempty"`
}

// IndexListResponse represents a stock market index listing.
type IndexListResponse struct {
	// Symbol is the index symbol.
	Symbol string `json:"symbol"`
	// Name is the index name.
	Name string `json:"name"`
	// Exchange is the exchange where the index is listed.
	Exchange string `json:"exchange"`
	// Currency is the trading currency.
	Currency string `json:"currency"`
}

// IndexQuoteResponse represents an index quote.
type IndexQuoteResponse struct {
	// Symbol is the index symbol.
	Symbol string `json:"symbol"`
	// Name is the index name.
	Name string `json:"name"`
	// Price is the current index value.
	Price float64 `json:"price"`
	// ChangePercentage is the percentage change.
	ChangePercentage float64 `json:"changePercentage"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
	// DayLow is the lowest value of the day.
	DayLow float64 `json:"dayLow"`
	// DayHigh is the highest value of the day.
	DayHigh float64 `json:"dayHigh"`
	// YearHigh is the 52-week high.
	YearHigh float64 `json:"yearHigh"`
	// YearLow is the 52-week low.
	YearLow float64 `json:"yearLow"`
	// MarketCap is typically null for indexes.
	MarketCap *int64 `json:"marketCap"`
	// PriceAvg50 is the 50-day moving average.
	PriceAvg50 float64 `json:"priceAvg50"`
	// PriceAvg200 is the 200-day moving average.
	PriceAvg200 float64 `json:"priceAvg200"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
	// Open is the opening value.
	Open float64 `json:"open"`
	// PreviousClose is the previous closing value.
	PreviousClose float64 `json:"previousClose"`
	// Timestamp is the Unix timestamp.
	Timestamp int64 `json:"timestamp"`
}

// IndexQuoteShortResponse represents a short index quote.
type IndexQuoteShortResponse struct {
	// Symbol is the index symbol.
	Symbol string `json:"symbol"`
	// Price is the current index value.
	Price float64 `json:"price"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// IndexHistoricalLightResponse represents simplified historical data.
type IndexHistoricalLightResponse struct {
	// Symbol is the index symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Price is the closing index value.
	Price float64 `json:"price"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// IndexHistoricalFullResponse represents full historical data.
type IndexHistoricalFullResponse struct {
	// Symbol is the index symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Open is the opening value.
	Open float64 `json:"open"`
	// High is the highest value during the period.
	High float64 `json:"high"`
	// Low is the lowest value during the period.
	Low float64 `json:"low"`
	// Close is the closing value.
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

// IndexIntradayResponse represents intraday index data.
type IndexIntradayResponse struct {
	// Date is the date and time of the data point.
	Date string `json:"date"`
	// Open is the opening value for the interval.
	Open float64 `json:"open"`
	// Low is the lowest value during the interval.
	Low float64 `json:"low"`
	// High is the highest value during the interval.
	High float64 `json:"high"`
	// Close is the closing value for the interval.
	Close float64 `json:"close"`
	// Volume is the trading volume for the interval.
	Volume int64 `json:"volume"`
}

// IndexConstituentResponse represents an index constituent company.
type IndexConstituentResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the company name.
	Name string `json:"name"`
	// Sector is the business sector.
	Sector string `json:"sector"`
	// SubSector is the sub-sector classification.
	SubSector string `json:"subSector"`
	// HeadQuarter is the company headquarters location.
	HeadQuarter string `json:"headQuarter"`
	// DateFirstAdded is when the company was added to the index.
	DateFirstAdded *string `json:"dateFirstAdded"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// Founded is the company founding date.
	Founded string `json:"founded"`
}

// HistoricalIndexChangeResponse represents historical index changes.
type HistoricalIndexChangeResponse struct {
	// DateAdded is when the change occurred.
	DateAdded string `json:"dateAdded"`
	// AddedSecurity is the company that was added.
	AddedSecurity string `json:"addedSecurity"`
	// RemovedTicker is the ticker of the removed company.
	RemovedTicker string `json:"removedTicker"`
	// RemovedSecurity is the company that was removed.
	RemovedSecurity string `json:"removedSecurity"`
	// Date is the effective date of the change.
	Date string `json:"date"`
	// Symbol is the added company's ticker.
	Symbol string `json:"symbol"`
	// Reason is the reason for the change.
	Reason string `json:"reason"`
}

// IndexList retrieves all available stock market indexes.
// Returns indexes across global exchanges with symbols and names.
//
// Endpoint: /index-list
func (c *Client) IndexList() ([]IndexListResponse, error) {
	pathName := "/index-list"

	var result []IndexListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// IndexQuote retrieves real-time quote for a specific index.
// Returns current value, changes, and market data.
//
// Endpoint: /quote
func (c *Client) IndexQuote(params IndexQuoteParams) ([]IndexQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote"

	var result []IndexQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndexQuoteShort retrieves short form quote for an index.
// Returns essential value and volume data.
//
// Endpoint: /quote-short
func (c *Client) IndexQuoteShort(params IndexQuoteParams) ([]IndexQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote-short"

	var result []IndexQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BatchIndexQuotes retrieves quotes for all indexes.
// Returns batch quotes for market indexes in a single request.
//
// Endpoint: /batch-index-quotes
func (c *Client) BatchIndexQuotes(params IndexBatchQuotesParams) ([]IndexQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-index-quotes"

	var result []IndexQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndexHistoricalLight retrieves simplified historical index data.
// Returns end-of-day values and volume.
//
// Endpoint: /historical-price-eod/light
func (c *Client) IndexHistoricalLight(params IndexHistoricalParams) ([]IndexHistoricalLightResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/light"

	var result []IndexHistoricalLightResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndexHistoricalFull retrieves full historical index data.
// Returns OHLC values, volume, and changes.
//
// Endpoint: /historical-price-eod/full
func (c *Client) IndexHistoricalFull(params IndexHistoricalParams) ([]IndexHistoricalFullResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/full"

	var result []IndexHistoricalFullResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndexChart1Min retrieves 1-minute interval index data.
// Returns intraday value movements for detailed analysis.
//
// Endpoint: /historical-chart/1min
func (c *Client) IndexChart1Min(params IndexHistoricalParams) ([]IndexIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1min"

	var result []IndexIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndexChart5Min retrieves 5-minute interval index data.
// Returns short-term value trends for trading analysis.
//
// Endpoint: /historical-chart/5min
func (c *Client) IndexChart5Min(params IndexHistoricalParams) ([]IndexIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/5min"

	var result []IndexIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndexChart1Hour retrieves hourly index data.
// Returns hourly value movements for mid-term analysis.
//
// Endpoint: /historical-chart/1hour
func (c *Client) IndexChart1Hour(params IndexHistoricalParams) ([]IndexIntradayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1hour"

	var result []IndexIntradayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SP500Constituents retrieves S&P 500 index constituents.
// Returns all companies in the S&P 500 with details.
//
// Endpoint: /sp500-constituent
func (c *Client) SP500Constituents() ([]IndexConstituentResponse, error) {
	pathName := "/sp500-constituent"

	var result []IndexConstituentResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// NasdaqConstituents retrieves Nasdaq index constituents.
// Returns all companies in the Nasdaq index with details.
//
// Endpoint: /nasdaq-constituent
func (c *Client) NasdaqConstituents() ([]IndexConstituentResponse, error) {
	pathName := "/nasdaq-constituent"

	var result []IndexConstituentResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// DowJonesConstituents retrieves Dow Jones index constituents.
// Returns all 30 companies in the Dow Jones Industrial Average.
//
// Endpoint: /dowjones-constituent
func (c *Client) DowJonesConstituents() ([]IndexConstituentResponse, error) {
	pathName := "/dowjones-constituent"

	var result []IndexConstituentResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// HistoricalSP500 retrieves historical S&P 500 index changes.
// Returns additions and removals from the S&P 500 over time.
//
// Endpoint: /historical-sp500-constituent
func (c *Client) HistoricalSP500() ([]HistoricalIndexChangeResponse, error) {
	pathName := "/historical-sp500-constituent"

	var result []HistoricalIndexChangeResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// HistoricalNasdaq retrieves historical Nasdaq index changes.
// Returns additions and removals from the Nasdaq index over time.
//
// Endpoint: /historical-nasdaq-constituent
func (c *Client) HistoricalNasdaq() ([]HistoricalIndexChangeResponse, error) {
	pathName := "/historical-nasdaq-constituent"

	var result []HistoricalIndexChangeResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// HistoricalDowJones retrieves historical Dow Jones index changes.
// Returns additions and removals from the Dow Jones over time.
//
// Endpoint: /historical-dowjones-constituent
func (c *Client) HistoricalDowJones() ([]HistoricalIndexChangeResponse, error) {
	pathName := "/historical-dowjones-constituent"

	var result []HistoricalIndexChangeResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}