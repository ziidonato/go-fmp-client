package fmp_client

// QuoteParams represents parameters for single quote endpoints.
type QuoteParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// BatchQuoteParams represents parameters for batch quote endpoints.
type BatchQuoteParams struct {
	// Symbols is a comma-separated list of ticker symbols (required).
	Symbols string `json:"symbols"`
}

// ExchangeQuoteParams represents parameters for exchange-wide quotes.
type ExchangeQuoteParams struct {
	// Exchange is the exchange name (required).
	Exchange string `json:"exchange"`
	// Short indicates whether to return short form data (optional).
	Short *bool `json:"short,omitempty"`
}

// BulkQuoteParams represents parameters for bulk quote endpoints.
type BulkQuoteParams struct {
	// Short indicates whether to return short form data (optional).
	Short *bool `json:"short,omitempty"`
}

// StockQuoteResponse represents a full stock quote.
type StockQuoteResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the company name.
	Name string `json:"name"`
	// Price is the current stock price.
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

// StockQuoteShortResponse represents a short form stock quote.
type StockQuoteShortResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
}

// AftermarketTradeResponse represents aftermarket trade data.
type AftermarketTradeResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Price is the trade price.
	Price float64 `json:"price"`
	// TradeSize is the size of the trade.
	TradeSize int `json:"tradeSize"`
	// Timestamp is the Unix timestamp in milliseconds.
	Timestamp int64 `json:"timestamp"`
}

// AftermarketQuoteResponse represents aftermarket quote data.
type AftermarketQuoteResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// BidSize is the size of the bid.
	BidSize int `json:"bidSize"`
	// BidPrice is the bid price.
	BidPrice float64 `json:"bidPrice"`
	// AskSize is the size of the ask.
	AskSize int `json:"askSize"`
	// AskPrice is the ask price.
	AskPrice float64 `json:"askPrice"`
	// Volume is the trading volume.
	Volume int64 `json:"volume"`
	// Timestamp is the Unix timestamp in milliseconds.
	Timestamp int64 `json:"timestamp"`
}

// StockPriceChangeResponse represents price changes over various periods.
type StockPriceChangeResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// OneDay is the 1-day percentage change.
	OneDay float64 `json:"1D"`
	// FiveDay is the 5-day percentage change.
	FiveDay float64 `json:"5D"`
	// OneMonth is the 1-month percentage change.
	OneMonth float64 `json:"1M"`
	// ThreeMonth is the 3-month percentage change.
	ThreeMonth float64 `json:"3M"`
	// SixMonth is the 6-month percentage change.
	SixMonth float64 `json:"6M"`
	// YTD is the year-to-date percentage change.
	YTD float64 `json:"ytd"`
	// OneYear is the 1-year percentage change.
	OneYear float64 `json:"1Y"`
	// ThreeYear is the 3-year percentage change.
	ThreeYear float64 `json:"3Y"`
	// FiveYear is the 5-year percentage change.
	FiveYear float64 `json:"5Y"`
	// TenYear is the 10-year percentage change.
	TenYear float64 `json:"10Y"`
	// Max is the maximum historical percentage change.
	Max float64 `json:"max"`
}

// StockQuote retrieves real-time stock quote data for a single symbol.
// Returns comprehensive price, volume, and change information.
//
// Endpoint: /quote
func (c *Client) StockQuote(params QuoteParams) ([]StockQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote"

	var result []StockQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockQuoteShort retrieves a short form stock quote for quick snapshots.
// Returns only essential data: price, change, and volume.
//
// Endpoint: /quote-short
func (c *Client) StockQuoteShort(params QuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/quote-short"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AftermarketTrade retrieves aftermarket trade data for a symbol.
// Shows individual trades executed after regular market hours.
//
// Endpoint: /aftermarket-trade
func (c *Client) AftermarketTrade(params QuoteParams) ([]AftermarketTradeResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/aftermarket-trade"

	var result []AftermarketTradeResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AftermarketQuote retrieves aftermarket quote data for a symbol.
// Shows bid/ask spreads during post-market trading.
//
// Endpoint: /aftermarket-quote
func (c *Client) AftermarketQuote(params QuoteParams) ([]AftermarketQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/aftermarket-quote"

	var result []AftermarketQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockPriceChange retrieves price changes over various time periods.
// Shows percentage changes from 1 day to maximum historical range.
//
// Endpoint: /stock-price-change
func (c *Client) StockPriceChange(params QuoteParams) ([]StockPriceChangeResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/stock-price-change"

	var result []StockPriceChangeResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockBatchQuote retrieves full quotes for multiple symbols at once.
// Efficient way to get comprehensive data for multiple stocks.
//
// Endpoint: /batch-quote
func (c *Client) StockBatchQuote(params BatchQuoteParams) ([]StockQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-quote"

	var result []StockQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockBatchQuoteShort retrieves short form quotes for multiple symbols.
// Efficient way to get essential data for multiple stocks.
//
// Endpoint: /batch-quote-short
func (c *Client) StockBatchQuoteShort(params BatchQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-quote-short"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BatchAftermarketTrade retrieves aftermarket trades for multiple symbols.
// Shows post-market trading activity across multiple stocks.
//
// Endpoint: /batch-aftermarket-trade
func (c *Client) BatchAftermarketTrade(params BatchQuoteParams) ([]AftermarketTradeResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-aftermarket-trade"

	var result []AftermarketTradeResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BatchAftermarketQuote retrieves aftermarket quotes for multiple symbols.
// Shows post-market bid/ask data across multiple stocks.
//
// Endpoint: /batch-aftermarket-quote
func (c *Client) BatchAftermarketQuote(params BatchQuoteParams) ([]AftermarketQuoteResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-aftermarket-quote"

	var result []AftermarketQuoteResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ExchangeStockQuotes retrieves quotes for all stocks on a specific exchange.
// Can return either full or short form data based on the short parameter.
//
// Endpoint: /batch-exchange-quote
func (c *Client) ExchangeStockQuotes(params ExchangeQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-exchange-quote"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// MutualFundQuotes retrieves quotes for all mutual funds.
// Returns price and change data for mutual funds.
//
// Endpoint: /batch-mutualfund-quotes
func (c *Client) MutualFundQuotes(params BulkQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-mutualfund-quotes"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ETFQuotes retrieves quotes for all ETFs.
// Returns price and change data for exchange-traded funds.
//
// Endpoint: /batch-etf-quotes
func (c *Client) ETFQuotes(params BulkQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-etf-quotes"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CommodityQuotes retrieves quotes for all commodities.
// Returns price and change data for commodities like oil, gold, etc.
//
// Endpoint: /batch-commodity-quotes
func (c *Client) CommodityQuotes(params BulkQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-commodity-quotes"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoQuotes retrieves quotes for all cryptocurrencies.
// Returns price and change data for digital assets.
//
// Endpoint: /batch-crypto-quotes
func (c *Client) CryptoQuotes(params BulkQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-crypto-quotes"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexQuotes retrieves quotes for all forex currency pairs.
// Returns price and change data for currency pairs.
//
// Endpoint: /batch-forex-quotes
func (c *Client) ForexQuotes(params BulkQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-forex-quotes"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndexQuotes retrieves quotes for all stock market indexes.
// Returns price and change data for major indexes.
//
// Endpoint: /batch-index-quotes
func (c *Client) IndexQuotes(params BulkQuoteParams) ([]StockQuoteShortResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/batch-index-quotes"

	var result []StockQuoteShortResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}