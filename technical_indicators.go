package fmp_client

import "time"

// TechnicalIndicatorParams represents parameters for technical indicator endpoints.
type TechnicalIndicatorParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// PeriodLength is the period for the indicator calculation (required).
	PeriodLength int `json:"periodLength"`
	// Timeframe is the timeframe: "1min", "5min", "15min", "30min", "1hour", "4hour", "1day" (required).
	Timeframe string `json:"timeframe"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// MACDParams represents parameters for MACD indicator.
type MACDParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// FastPeriod is the fast EMA period (required).
	FastPeriod int `json:"fastPeriod"`
	// SlowPeriod is the slow EMA period (required).
	SlowPeriod int `json:"slowPeriod"`
	// SignalPeriod is the signal line period (required).
	SignalPeriod int `json:"signalPeriod"`
	// Timeframe is the timeframe (required).
	Timeframe string `json:"timeframe"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// BollingerBandsParams represents parameters for Bollinger Bands.
type BollingerBandsParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// PeriodLength is the moving average period (required).
	PeriodLength int `json:"periodLength"`
	// StandardDeviations is the number of standard deviations (required).
	StandardDeviations int `json:"standardDeviations"`
	// Timeframe is the timeframe (required).
	Timeframe string `json:"timeframe"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// TechnicalIndicatorResponse represents a base response with price data.
type TechnicalIndicatorResponse struct {
	// Date is the date and time of the data point.
	Date string `json:"date"`
	// Open is the opening price.
	Open float64 `json:"open"`
	// High is the highest price.
	High float64 `json:"high"`
	// Low is the lowest price.
	Low float64 `json:"low"`
	// Close is the closing price.
	Close float64 `json:"close"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// SMAResponse represents Simple Moving Average response.
type SMAResponse struct {
	TechnicalIndicatorResponse
	// SMA is the simple moving average value.
	SMA float64 `json:"sma"`
}

// EMAResponse represents Exponential Moving Average response.
type EMAResponse struct {
	TechnicalIndicatorResponse
	// EMA is the exponential moving average value.
	EMA float64 `json:"ema"`
}

// WMAResponse represents Weighted Moving Average response.
type WMAResponse struct {
	TechnicalIndicatorResponse
	// WMA is the weighted moving average value.
	WMA float64 `json:"wma"`
}

// DEMAResponse represents Double Exponential Moving Average response.
type DEMAResponse struct {
	TechnicalIndicatorResponse
	// DEMA is the double exponential moving average value.
	DEMA float64 `json:"dema"`
}

// TEMAResponse represents Triple Exponential Moving Average response.
type TEMAResponse struct {
	TechnicalIndicatorResponse
	// TEMA is the triple exponential moving average value.
	TEMA float64 `json:"tema"`
}

// RSIResponse represents Relative Strength Index response.
type RSIResponse struct {
	TechnicalIndicatorResponse
	// RSI is the relative strength index value.
	RSI float64 `json:"rsi"`
}

// MACDResponse represents MACD response.
type MACDResponse struct {
	TechnicalIndicatorResponse
	// MACD is the MACD line value.
	MACD float64 `json:"macd"`
	// Signal is the signal line value.
	Signal float64 `json:"signal"`
	// Histogram is the MACD histogram value.
	Histogram float64 `json:"histogram"`
}

// BollingerBandsResponse represents Bollinger Bands response.
type BollingerBandsResponse struct {
	TechnicalIndicatorResponse
	// UpperBand is the upper Bollinger Band.
	UpperBand float64 `json:"upperBand"`
	// MiddleBand is the middle band (SMA).
	MiddleBand float64 `json:"middleBand"`
	// LowerBand is the lower Bollinger Band.
	LowerBand float64 `json:"lowerBand"`
}

// OBVResponse represents On-Balance Volume response.
type OBVResponse struct {
	TechnicalIndicatorResponse
	// OBV is the on-balance volume value.
	OBV float64 `json:"obv"`
}

// VolumeIndicatorResponse represents volume-based indicator response.
type VolumeIndicatorResponse struct {
	TechnicalIndicatorResponse
	// VolumeValue is the volume indicator value.
	VolumeValue float64 `json:"volumeValue"`
}

// SimpleMovingAverage calculates the Simple Moving Average.
// Returns average price over a specified period.
//
// Endpoint: /technical-indicators/sma
func (c *Client) SimpleMovingAverage(params TechnicalIndicatorParams) ([]SMAResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/sma"

	var result []SMAResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ExponentialMovingAverage calculates the Exponential Moving Average.
// Returns weighted average giving more weight to recent prices.
//
// Endpoint: /technical-indicators/ema
func (c *Client) ExponentialMovingAverage(params TechnicalIndicatorParams) ([]EMAResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/ema"

	var result []EMAResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// WeightedMovingAverage calculates the Weighted Moving Average.
// Returns linearly weighted average of prices.
//
// Endpoint: /technical-indicators/wma
func (c *Client) WeightedMovingAverage(params TechnicalIndicatorParams) ([]WMAResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/wma"

	var result []WMAResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// DoubleExponentialMovingAverage calculates the Double EMA.
// Returns smoothed EMA with reduced lag.
//
// Endpoint: /technical-indicators/dema
func (c *Client) DoubleExponentialMovingAverage(params TechnicalIndicatorParams) ([]DEMAResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/dema"

	var result []DEMAResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// TripleExponentialMovingAverage calculates the Triple EMA.
// Returns further smoothed EMA with minimal lag.
//
// Endpoint: /technical-indicators/tema
func (c *Client) TripleExponentialMovingAverage(params TechnicalIndicatorParams) ([]TEMAResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/tema"

	var result []TEMAResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// RelativeStrengthIndex calculates the RSI.
// Returns momentum oscillator measuring price changes.
//
// Endpoint: /technical-indicators/rsi
func (c *Client) RelativeStrengthIndex(params TechnicalIndicatorParams) ([]RSIResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/rsi"

	var result []RSIResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// MACD calculates Moving Average Convergence Divergence.
// Returns MACD line, signal line, and histogram.
//
// Endpoint: /technical-indicators/macd
func (c *Client) MACD(params MACDParams) ([]MACDResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/macd"

	var result []MACDResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BollingerBands calculates Bollinger Bands.
// Returns upper, middle, and lower bands.
//
// Endpoint: /technical-indicators/bb
func (c *Client) BollingerBands(params BollingerBandsParams) ([]BollingerBandsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/bb"

	var result []BollingerBandsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// OnBalanceVolume calculates OBV indicator.
// Returns cumulative volume based on price direction.
//
// Endpoint: /technical-indicators/obv
func (c *Client) OnBalanceVolume(params TechnicalIndicatorParams) ([]OBVResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/technical-indicators/obv"

	var result []OBVResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}