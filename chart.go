package fmp_client

import "time"

// ChartPeriodParams represents parameters common to chart endpoints that accept date ranges.
type ChartPeriodParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// From is the start date for the data range (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for the data range (optional).
	To *time.Time `json:"to,omitempty"`
}

// IntradayChartParams represents parameters for intraday chart endpoints.
type IntradayChartParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// From is the start date for the data range (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for the data range (optional).
	To *time.Time `json:"to,omitempty"`
	// NonAdjusted indicates whether to return non-adjusted prices (optional).
	NonAdjusted *bool `json:"nonadjusted,omitempty"`
}

// StockChartLightResponse represents simplified chart data.
type StockChartLightResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Price is the closing price.
	Price float64 `json:"price"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// StockChartFullResponse represents full price and volume data.
type StockChartFullResponse struct {
	// Symbol is the stock ticker symbol.
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

// StockChartAdjustedResponse represents adjusted price data.
type StockChartAdjustedResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// AdjOpen is the adjusted opening price.
	AdjOpen float64 `json:"adjOpen"`
	// AdjHigh is the adjusted highest price.
	AdjHigh float64 `json:"adjHigh"`
	// AdjLow is the adjusted lowest price.
	AdjLow float64 `json:"adjLow"`
	// AdjClose is the adjusted closing price.
	AdjClose float64 `json:"adjClose"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// IntradayChartResponse represents intraday chart data.
type IntradayChartResponse struct {
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

// StockChartLight retrieves simplified stock chart data including date, price, and volume.
// This is ideal for tracking basic stock performance and creating simple charts.
//
// Endpoint: /historical-price-eod/light
func (c *Client) StockChartLight(params ChartPeriodParams) ([]StockChartLightResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/light"

	var result []StockChartLightResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChartFull retrieves comprehensive price and volume data including OHLC, changes, and VWAP.
// This provides detailed insights for technical analysis and advanced charting.
//
// Endpoint: /historical-price-eod/full
func (c *Client) StockChartFull(params ChartPeriodParams) ([]StockChartFullResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/full"

	var result []StockChartFullResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChartNonSplitAdjusted retrieves stock price data without adjustments for stock splits.
// This shows actual historical prices as they were traded.
//
// Endpoint: /historical-price-eod/non-split-adjusted
func (c *Client) StockChartNonSplitAdjusted(params ChartPeriodParams) ([]StockChartAdjustedResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/non-split-adjusted"

	var result []StockChartAdjustedResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChartDividendAdjusted retrieves stock price data adjusted for dividend payouts.
// This provides a more accurate view of total returns including dividends.
//
// Endpoint: /historical-price-eod/dividend-adjusted
func (c *Client) StockChartDividendAdjusted(params ChartPeriodParams) ([]StockChartAdjustedResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-price-eod/dividend-adjusted"

	var result []StockChartAdjustedResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChart1Min retrieves intraday stock data in 1-minute intervals.
// This provides the most granular data for high-frequency trading and analysis.
//
// Endpoint: /historical-chart/1min
func (c *Client) StockChart1Min(params IntradayChartParams) ([]IntradayChartResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1min"

	var result []IntradayChartResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChart5Min retrieves intraday stock data in 5-minute intervals.
// This is ideal for short-term trading analysis and intraday pattern recognition.
//
// Endpoint: /historical-chart/5min
func (c *Client) StockChart5Min(params IntradayChartParams) ([]IntradayChartResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/5min"

	var result []IntradayChartResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChart15Min retrieves intraday stock data in 15-minute intervals.
// This provides a balance between detail and broader intraday trends.
//
// Endpoint: /historical-chart/15min
func (c *Client) StockChart15Min(params IntradayChartParams) ([]IntradayChartResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/15min"

	var result []IntradayChartResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChart30Min retrieves intraday stock data in 30-minute intervals.
// This is suitable for medium-term intraday analysis and position trading.
//
// Endpoint: /historical-chart/30min
func (c *Client) StockChart30Min(params IntradayChartParams) ([]IntradayChartResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/30min"

	var result []IntradayChartResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChart1Hour retrieves intraday stock data in 1-hour intervals.
// This captures broader intraday trends while maintaining precision.
//
// Endpoint: /historical-chart/1hour
func (c *Client) StockChart1Hour(params IntradayChartParams) ([]IntradayChartResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/1hour"

	var result []IntradayChartResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockChart4Hour retrieves intraday stock data in 4-hour intervals.
// This is ideal for tracking extended intraday movements and swing trading.
//
// Endpoint: /historical-chart/4hour
func (c *Client) StockChart4Hour(params IntradayChartParams) ([]IntradayChartResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-chart/4hour"

	var result []IntradayChartResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}