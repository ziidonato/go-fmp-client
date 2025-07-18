package fmp_client

import "time"

// NewsPageParams represents pagination parameters for news endpoints.
type NewsPageParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// NewsDateParams represents parameters with date range and pagination.
type NewsDateParams struct {
	// From is the start date for filtering news (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering news (optional).
	To *time.Time `json:"to,omitempty"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// NewsSearchParams represents parameters for searching news by symbols.
type NewsSearchParams struct {
	// Symbols is a comma-separated list of symbols to search for (required).
	Symbols string `json:"symbols"`
	// From is the start date for filtering news (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering news (optional).
	To *time.Time `json:"to,omitempty"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// FMPArticleResponse represents an FMP article.
type FMPArticleResponse struct {
	// Title is the article headline.
	Title string `json:"title"`
	// Date is the publication date and time.
	Date string `json:"date"`
	// Content is the article content (may include HTML).
	Content string `json:"content"`
	// Tickers is the stock tickers mentioned in the article.
	Tickers string `json:"tickers"`
	// Image is the URL to the article image.
	Image string `json:"image"`
	// Link is the URL to the full article.
	Link string `json:"link"`
	// Author is the name of the article author.
	Author string `json:"author"`
	// Site is the source website.
	Site string `json:"site"`
}

// GeneralNewsResponse represents a general news article.
type GeneralNewsResponse struct {
	// Symbol is the related symbol (may be null).
	Symbol *string `json:"symbol"`
	// PublishedDate is the publication date and time.
	PublishedDate string `json:"publishedDate"`
	// Publisher is the news source.
	Publisher string `json:"publisher"`
	// Title is the article headline.
	Title string `json:"title"`
	// Image is the URL to the article image.
	Image string `json:"image"`
	// Site is the publisher's website.
	Site string `json:"site"`
	// Text is the article snippet or summary.
	Text string `json:"text"`
	// URL is the link to the full article.
	URL string `json:"url"`
}

// PressReleaseResponse represents a company press release.
type PressReleaseResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// PublishedDate is the publication date and time.
	PublishedDate string `json:"publishedDate"`
	// Publisher is the press release distributor.
	Publisher string `json:"publisher"`
	// Title is the press release headline.
	Title string `json:"title"`
	// Image is the URL to the press release image.
	Image string `json:"image"`
	// Site is the publisher's website.
	Site string `json:"site"`
	// Text is the press release content or summary.
	Text string `json:"text"`
	// URL is the link to the full press release.
	URL string `json:"url"`
}

// StockNewsResponse represents a stock-specific news article.
type StockNewsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// PublishedDate is the publication date and time.
	PublishedDate string `json:"publishedDate"`
	// Publisher is the news source.
	Publisher string `json:"publisher"`
	// Title is the article headline.
	Title string `json:"title"`
	// Image is the URL to the article image.
	Image string `json:"image"`
	// Site is the publisher's website.
	Site string `json:"site"`
	// Text is the article snippet or summary.
	Text string `json:"text"`
	// URL is the link to the full article.
	URL string `json:"url"`
}

// CryptoNewsResponse represents a cryptocurrency news article.
type CryptoNewsResponse struct {
	// Symbol is the cryptocurrency symbol.
	Symbol string `json:"symbol"`
	// PublishedDate is the publication date and time.
	PublishedDate string `json:"publishedDate"`
	// Publisher is the news source.
	Publisher string `json:"publisher"`
	// Title is the article headline.
	Title string `json:"title"`
	// Image is the URL to the article image.
	Image string `json:"image"`
	// Site is the publisher's website.
	Site string `json:"site"`
	// Text is the article snippet or summary.
	Text string `json:"text"`
	// URL is the link to the full article.
	URL string `json:"url"`
}

// ForexNewsResponse represents a forex news article.
type ForexNewsResponse struct {
	// Symbol is the currency pair symbol.
	Symbol string `json:"symbol"`
	// PublishedDate is the publication date and time.
	PublishedDate string `json:"publishedDate"`
	// Publisher is the news source.
	Publisher string `json:"publisher"`
	// Title is the article headline.
	Title string `json:"title"`
	// Image is the URL to the article image.
	Image string `json:"image"`
	// Site is the publisher's website.
	Site string `json:"site"`
	// Text is the article snippet or summary.
	Text string `json:"text"`
	// URL is the link to the full article.
	URL string `json:"url"`
}

// FMPArticles retrieves the latest articles from Financial Modeling Prep.
// Returns FMP's proprietary news content with analysis and insights.
//
// Endpoint: /fmp-articles
func (c *Client) FMPArticles(params NewsPageParams) ([]FMPArticleResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/fmp-articles"

	var result []FMPArticleResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// GeneralNews retrieves the latest general news from various sources.
// Returns broad market and business news coverage.
//
// Endpoint: /news/general-latest
func (c *Client) GeneralNews(params NewsDateParams) ([]GeneralNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/general-latest"

	var result []GeneralNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// PressReleases retrieves the latest company press releases.
// Returns official corporate announcements and updates.
//
// Endpoint: /news/press-releases-latest
func (c *Client) PressReleases(params NewsDateParams) ([]PressReleaseResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/press-releases-latest"

	var result []PressReleaseResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockNews retrieves the latest stock market news.
// Returns news articles related to stock market activity.
//
// Endpoint: /news/stock-latest
func (c *Client) StockNews(params NewsDateParams) ([]StockNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/stock-latest"

	var result []StockNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CryptoNews retrieves the latest cryptocurrency news.
// Returns news articles about digital assets and crypto markets.
//
// Endpoint: /news/crypto-latest
func (c *Client) CryptoNews(params NewsDateParams) ([]CryptoNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/crypto-latest"

	var result []CryptoNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ForexNews retrieves the latest forex market news.
// Returns news articles about currency markets and forex trading.
//
// Endpoint: /news/forex-latest
func (c *Client) ForexNews(params NewsDateParams) ([]ForexNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/forex-latest"

	var result []ForexNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchPressReleases searches for press releases by company symbols.
// Returns press releases for specific companies.
//
// Endpoint: /news/press-releases
func (c *Client) SearchPressReleases(params NewsSearchParams) ([]PressReleaseResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/press-releases"

	var result []PressReleaseResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchStockNews searches for stock news by ticker symbols.
// Returns news articles for specific stocks.
//
// Endpoint: /news/stock
func (c *Client) SearchStockNews(params NewsSearchParams) ([]StockNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/stock"

	var result []StockNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchCryptoNews searches for cryptocurrency news by symbols.
// Returns news articles for specific cryptocurrencies.
//
// Endpoint: /news/crypto
func (c *Client) SearchCryptoNews(params NewsSearchParams) ([]CryptoNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/crypto"

	var result []CryptoNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchForexNews searches for forex news by currency pair symbols.
// Returns news articles for specific currency pairs.
//
// Endpoint: /news/forex
func (c *Client) SearchForexNews(params NewsSearchParams) ([]ForexNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/news/forex"

	var result []ForexNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}