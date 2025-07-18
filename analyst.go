package fmp_client

// AnalystEstimatesParams represents parameters for analyst estimates.
type AnalystEstimatesParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Period is the time period: "annual" or "quarter" (required).
	Period string `json:"period"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// RatingsParams represents parameters for ratings endpoints.
type RatingsParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// PriceTargetParams represents parameters for price target endpoints.
type PriceTargetParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// PriceTargetNewsParams represents parameters for price target news.
type PriceTargetNewsParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// LatestNewsParams represents parameters for latest news endpoints.
type LatestNewsParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// GradesParams represents parameters for grades endpoints.
type GradesParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// GradesNewsParams represents parameters for grades news.
type GradesNewsParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// AnalystEstimatesResponse represents analyst financial estimates.
type AnalystEstimatesResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the estimate date.
	Date string `json:"date"`
	// RevenueLow is the low revenue estimate.
	RevenueLow int64 `json:"revenueLow"`
	// RevenueHigh is the high revenue estimate.
	RevenueHigh int64 `json:"revenueHigh"`
	// RevenueAvg is the average revenue estimate.
	RevenueAvg int64 `json:"revenueAvg"`
	// EbitdaLow is the low EBITDA estimate.
	EbitdaLow int64 `json:"ebitdaLow"`
	// EbitdaHigh is the high EBITDA estimate.
	EbitdaHigh int64 `json:"ebitdaHigh"`
	// EbitdaAvg is the average EBITDA estimate.
	EbitdaAvg int64 `json:"ebitdaAvg"`
	// EbitLow is the low EBIT estimate.
	EbitLow int64 `json:"ebitLow"`
	// EbitHigh is the high EBIT estimate.
	EbitHigh int64 `json:"ebitHigh"`
	// EbitAvg is the average EBIT estimate.
	EbitAvg int64 `json:"ebitAvg"`
	// NetIncomeLow is the low net income estimate.
	NetIncomeLow int64 `json:"netIncomeLow"`
	// NetIncomeHigh is the high net income estimate.
	NetIncomeHigh int64 `json:"netIncomeHigh"`
	// NetIncomeAvg is the average net income estimate.
	NetIncomeAvg int64 `json:"netIncomeAvg"`
	// SgaExpenseLow is the low SG&A expense estimate.
	SgaExpenseLow int64 `json:"sgaExpenseLow"`
	// SgaExpenseHigh is the high SG&A expense estimate.
	SgaExpenseHigh int64 `json:"sgaExpenseHigh"`
	// SgaExpenseAvg is the average SG&A expense estimate.
	SgaExpenseAvg int64 `json:"sgaExpenseAvg"`
	// EPSAvg is the average EPS estimate.
	EPSAvg float64 `json:"epsAvg"`
	// EPSHigh is the high EPS estimate.
	EPSHigh float64 `json:"epsHigh"`
	// EPSLow is the low EPS estimate.
	EPSLow float64 `json:"epsLow"`
	// NumAnalystsRevenue is the number of analysts providing revenue estimates.
	NumAnalystsRevenue int `json:"numAnalystsRevenue"`
	// NumAnalystsEPS is the number of analysts providing EPS estimates.
	NumAnalystsEPS int `json:"numAnalystsEps"`
}

// RatingsResponse represents financial ratings.
type RatingsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the rating date (for historical).
	Date string `json:"date,omitempty"`
	// Rating is the overall letter rating.
	Rating string `json:"rating"`
	// OverallScore is the overall numerical score.
	OverallScore int `json:"overallScore"`
	// DiscountedCashFlowScore is the DCF score.
	DiscountedCashFlowScore int `json:"discountedCashFlowScore"`
	// ReturnOnEquityScore is the ROE score.
	ReturnOnEquityScore int `json:"returnOnEquityScore"`
	// ReturnOnAssetsScore is the ROA score.
	ReturnOnAssetsScore int `json:"returnOnAssetsScore"`
	// DebtToEquityScore is the D/E ratio score.
	DebtToEquityScore int `json:"debtToEquityScore"`
	// PriceToEarningsScore is the P/E ratio score.
	PriceToEarningsScore int `json:"priceToEarningsScore"`
	// PriceToBookScore is the P/B ratio score.
	PriceToBookScore int `json:"priceToBookScore"`
}

// PriceTargetSummaryResponse represents price target summary.
type PriceTargetSummaryResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// LastMonthCount is the number of targets in last month.
	LastMonthCount int `json:"lastMonthCount"`
	// LastMonthAvgPriceTarget is the average target for last month.
	LastMonthAvgPriceTarget float64 `json:"lastMonthAvgPriceTarget"`
	// LastQuarterCount is the number of targets in last quarter.
	LastQuarterCount int `json:"lastQuarterCount"`
	// LastQuarterAvgPriceTarget is the average target for last quarter.
	LastQuarterAvgPriceTarget float64 `json:"lastQuarterAvgPriceTarget"`
	// LastYearCount is the number of targets in last year.
	LastYearCount int `json:"lastYearCount"`
	// LastYearAvgPriceTarget is the average target for last year.
	LastYearAvgPriceTarget float64 `json:"lastYearAvgPriceTarget"`
	// AllTimeCount is the total number of price targets.
	AllTimeCount int `json:"allTimeCount"`
	// AllTimeAvgPriceTarget is the all-time average target.
	AllTimeAvgPriceTarget float64 `json:"allTimeAvgPriceTarget"`
	// Publishers is the list of publishers.
	Publishers string `json:"publishers"`
}

// PriceTargetConsensusResponse represents consensus price targets.
type PriceTargetConsensusResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// TargetHigh is the highest price target.
	TargetHigh float64 `json:"targetHigh"`
	// TargetLow is the lowest price target.
	TargetLow float64 `json:"targetLow"`
	// TargetConsensus is the consensus price target.
	TargetConsensus float64 `json:"targetConsensus"`
	// TargetMedian is the median price target.
	TargetMedian float64 `json:"targetMedian"`
}

// PriceTargetNewsResponse represents price target news.
type PriceTargetNewsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// PublishedDate is when the news was published.
	PublishedDate string `json:"publishedDate"`
	// NewsURL is the link to the news article.
	NewsURL string `json:"newsURL"`
	// NewsTitle is the news headline.
	NewsTitle string `json:"newsTitle"`
	// AnalystName is the analyst's name.
	AnalystName string `json:"analystName"`
	// PriceTarget is the price target.
	PriceTarget float64 `json:"priceTarget"`
	// AdjPriceTarget is the adjusted price target.
	AdjPriceTarget float64 `json:"adjPriceTarget"`
	// PriceWhenPosted is the stock price when posted.
	PriceWhenPosted float64 `json:"priceWhenPosted"`
	// NewsPublisher is the news publisher.
	NewsPublisher string `json:"newsPublisher"`
	// NewsBaseURL is the base URL of the publisher.
	NewsBaseURL string `json:"newsBaseURL"`
	// AnalystCompany is the analyst's firm.
	AnalystCompany string `json:"analystCompany"`
}

// GradesResponse represents stock grades.
type GradesResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the grade date.
	Date string `json:"date"`
	// GradingCompany is the firm issuing the grade.
	GradingCompany string `json:"gradingCompany"`
	// PreviousGrade is the previous rating.
	PreviousGrade string `json:"previousGrade"`
	// NewGrade is the new rating.
	NewGrade string `json:"newGrade"`
	// Action is the action taken (upgrade/downgrade/maintain).
	Action string `json:"action"`
}

// GradesHistoricalResponse represents historical grade counts.
type GradesHistoricalResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Date is the date of the ratings.
	Date string `json:"date"`
	// AnalystRatingsBuy is the number of buy ratings.
	AnalystRatingsBuy int `json:"analystRatingsBuy"`
	// AnalystRatingsHold is the number of hold ratings.
	AnalystRatingsHold int `json:"analystRatingsHold"`
	// AnalystRatingsSell is the number of sell ratings.
	AnalystRatingsSell int `json:"analystRatingsSell"`
	// AnalystRatingsStrongSell is the number of strong sell ratings.
	AnalystRatingsStrongSell int `json:"analystRatingsStrongSell"`
}

// GradesConsensusResponse represents consensus grades.
type GradesConsensusResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// StrongBuy is the number of strong buy ratings.
	StrongBuy int `json:"strongBuy"`
	// Buy is the number of buy ratings.
	Buy int `json:"buy"`
	// Hold is the number of hold ratings.
	Hold int `json:"hold"`
	// Sell is the number of sell ratings.
	Sell int `json:"sell"`
	// StrongSell is the number of strong sell ratings.
	StrongSell int `json:"strongSell"`
	// Consensus is the overall consensus rating.
	Consensus string `json:"consensus"`
}

// GradesNewsResponse represents grade change news.
type GradesNewsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// PublishedDate is when the news was published.
	PublishedDate string `json:"publishedDate"`
	// NewsURL is the link to the news article.
	NewsURL string `json:"newsURL"`
	// NewsTitle is the news headline.
	NewsTitle string `json:"newsTitle"`
	// NewsBaseURL is the base URL of the publisher.
	NewsBaseURL string `json:"newsBaseURL"`
	// NewsPublisher is the news publisher.
	NewsPublisher string `json:"newsPublisher"`
	// NewGrade is the new rating.
	NewGrade string `json:"newGrade"`
	// PreviousGrade is the previous rating.
	PreviousGrade string `json:"previousGrade"`
	// GradingCompany is the firm issuing the grade.
	GradingCompany string `json:"gradingCompany"`
	// Action is the action taken.
	Action string `json:"action"`
	// PriceWhenPosted is the stock price when posted.
	PriceWhenPosted float64 `json:"priceWhenPosted"`
}

// AnalystEstimates retrieves analyst financial estimates.
// Returns projected revenue, earnings, and other metrics.
//
// Endpoint: /analyst-estimates
func (c *Client) AnalystEstimates(params AnalystEstimatesParams) ([]AnalystEstimatesResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/analyst-estimates"

	var result []AnalystEstimatesResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// RatingsSnapshot retrieves current financial ratings.
// Returns overall rating and component scores.
//
// Endpoint: /ratings-snapshot
func (c *Client) RatingsSnapshot(params RatingsParams) ([]RatingsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/ratings-snapshot"

	var result []RatingsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// RatingsHistorical retrieves historical financial ratings.
// Returns ratings and scores over time.
//
// Endpoint: /ratings-historical
func (c *Client) RatingsHistorical(params RatingsParams) ([]RatingsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/ratings-historical"

	var result []RatingsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// PriceTargetSummary retrieves price target summary statistics.
// Returns average targets across different timeframes.
//
// Endpoint: /price-target-summary
func (c *Client) PriceTargetSummary(params PriceTargetParams) ([]PriceTargetSummaryResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/price-target-summary"

	var result []PriceTargetSummaryResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// PriceTargetConsensus retrieves consensus price targets.
// Returns high, low, median, and consensus targets.
//
// Endpoint: /price-target-consensus
func (c *Client) PriceTargetConsensus(params PriceTargetParams) ([]PriceTargetConsensusResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/price-target-consensus"

	var result []PriceTargetConsensusResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// PriceTargetNews retrieves price target news for a symbol.
// Returns analyst price target updates with news links.
//
// Endpoint: /price-target-news
func (c *Client) PriceTargetNews(params PriceTargetNewsParams) ([]PriceTargetNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/price-target-news"

	var result []PriceTargetNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// PriceTargetLatestNews retrieves latest price target news.
// Returns recent price target updates across all symbols.
//
// Endpoint: /price-target-latest-news
func (c *Client) PriceTargetLatestNews(params LatestNewsParams) ([]PriceTargetNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/price-target-latest-news"

	var result []PriceTargetNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// StockGrades retrieves recent stock grade changes.
// Returns analyst rating changes for a symbol.
//
// Endpoint: /grades
func (c *Client) StockGrades(params GradesParams) ([]GradesResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/grades"

	var result []GradesResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// GradesHistorical retrieves historical analyst rating counts.
// Returns buy/hold/sell rating distributions over time.
//
// Endpoint: /grades-historical
func (c *Client) GradesHistorical(params RatingsParams) ([]GradesHistoricalResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/grades-historical"

	var result []GradesHistoricalResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// GradesConsensus retrieves consensus analyst ratings.
// Returns current rating distribution and consensus.
//
// Endpoint: /grades-consensus
func (c *Client) GradesConsensus(params GradesParams) ([]GradesConsensusResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/grades-consensus"

	var result []GradesConsensusResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// GradesNews retrieves stock grade news for a symbol.
// Returns analyst rating changes with news links.
//
// Endpoint: /grades-news
func (c *Client) GradesNews(params GradesNewsParams) ([]GradesNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/grades-news"

	var result []GradesNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// GradesLatestNews retrieves latest stock grade news.
// Returns recent rating changes across all symbols.
//
// Endpoint: /grades-latest-news
func (c *Client) GradesLatestNews(params LatestNewsParams) ([]GradesNewsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/grades-latest-news"

	var result []GradesNewsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}