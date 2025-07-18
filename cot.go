package fmp_client

import "time"

// COTReportParams represents parameters for COT report endpoints.
type COTReportParams struct {
	// Symbol is the commodity or futures symbol (optional).
	Symbol *string `json:"symbol,omitempty"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// COTAnalysisParams represents parameters for COT analysis endpoints.
type COTAnalysisParams struct {
	// Symbol is the commodity or futures symbol (optional).
	Symbol *string `json:"symbol,omitempty"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// COTReportResponse represents a Commitments of Traders report.
type COTReportResponse struct {
	// Symbol is the commodity or futures symbol.
	Symbol string `json:"symbol"`
	// Date is the report date.
	Date string `json:"date"`
	// Name is the commodity name.
	Name string `json:"name"`
	// Sector is the market sector.
	Sector string `json:"sector"`
	// MarketAndExchangeNames describes the market and exchange.
	MarketAndExchangeNames string `json:"marketAndExchangeNames"`
	// CFTCContractMarketCode is the CFTC contract code.
	CFTCContractMarketCode string `json:"cftcContractMarketCode"`
	// CFTCMarketCode is the CFTC market code.
	CFTCMarketCode string `json:"cftcMarketCode"`
	// CFTCRegionCode is the CFTC region code.
	CFTCRegionCode string `json:"cftcRegionCode"`
	// CFTCCommodityCode is the CFTC commodity code.
	CFTCCommodityCode string `json:"cftcCommodityCode"`
	// OpenInterestAll is total open interest.
	OpenInterestAll int64 `json:"openInterestAll"`
	// NoncommPositionsLongAll is non-commercial long positions.
	NoncommPositionsLongAll int64 `json:"noncommPositionsLongAll"`
	// NoncommPositionsShortAll is non-commercial short positions.
	NoncommPositionsShortAll int64 `json:"noncommPositionsShortAll"`
	// NoncommPositionsSpreadAll is non-commercial spread positions.
	NoncommPositionsSpreadAll int64 `json:"noncommPositionsSpreadAll"`
	// CommPositionsLongAll is commercial long positions.
	CommPositionsLongAll int64 `json:"commPositionsLongAll"`
	// CommPositionsShortAll is commercial short positions.
	CommPositionsShortAll int64 `json:"commPositionsShortAll"`
	// TotReptPositionsLongAll is total reportable long positions.
	TotReptPositionsLongAll int64 `json:"totReptPositionsLongAll"`
	// TotReptPositionsShortAll is total reportable short positions.
	TotReptPositionsShortAll int64 `json:"totReptPositionsShortAll"`
	// NonreptPositionsLongAll is non-reportable long positions.
	NonreptPositionsLongAll int64 `json:"nonreptPositionsLongAll"`
	// NonreptPositionsShortAll is non-reportable short positions.
	NonreptPositionsShortAll int64 `json:"nonreptPositionsShortAll"`
	// OpenInterestOld is old open interest.
	OpenInterestOld int64 `json:"openInterestOld"`
	// NoncommPositionsLongOld is old non-commercial long positions.
	NoncommPositionsLongOld int64 `json:"noncommPositionsLongOld"`
	// NoncommPositionsShortOld is old non-commercial short positions.
	NoncommPositionsShortOld int64 `json:"noncommPositionsShortOld"`
	// NoncommPositionsSpreadOld is old non-commercial spread positions.
	NoncommPositionsSpreadOld int64 `json:"noncommPositionsSpreadOld"`
	// CommPositionsLongOld is old commercial long positions.
	CommPositionsLongOld int64 `json:"commPositionsLongOld"`
	// CommPositionsShortOld is old commercial short positions.
	CommPositionsShortOld int64 `json:"commPositionsShortOld"`
	// TotReptPositionsLongOld is old total reportable long positions.
	TotReptPositionsLongOld int64 `json:"totReptPositionsLongOld"`
	// TotReptPositionsShortOld is old total reportable short positions.
	TotReptPositionsShortOld int64 `json:"totReptPositionsShortOld"`
	// NonreptPositionsLongOld is old non-reportable long positions.
	NonreptPositionsLongOld int64 `json:"nonreptPositionsLongOld"`
	// NonreptPositionsShortOld is old non-reportable short positions.
	NonreptPositionsShortOld int64 `json:"nonreptPositionsShortOld"`
	// OpenInterestOther is other open interest.
	OpenInterestOther int64 `json:"openInterestOther"`
	// NoncommPositionsLongOther is other non-commercial long positions.
	NoncommPositionsLongOther int64 `json:"noncommPositionsLongOther"`
	// NoncommPositionsShortOther is other non-commercial short positions.
	NoncommPositionsShortOther int64 `json:"noncommPositionsShortOther"`
	// NoncommPositionsSpreadOther is other non-commercial spread positions.
	NoncommPositionsSpreadOther int64 `json:"noncommPositionsSpreadOther"`
	// CommPositionsLongOther is other commercial long positions.
	CommPositionsLongOther int64 `json:"commPositionsLongOther"`
	// CommPositionsShortOther is other commercial short positions.
	CommPositionsShortOther int64 `json:"commPositionsShortOther"`
	// TotReptPositionsLongOther is other total reportable long positions.
	TotReptPositionsLongOther int64 `json:"totReptPositionsLongOther"`
	// TotReptPositionsShortOther is other total reportable short positions.
	TotReptPositionsShortOther int64 `json:"totReptPositionsShortOther"`
	// NonreptPositionsLongOther is other non-reportable long positions.
	NonreptPositionsLongOther int64 `json:"nonreptPositionsLongOther"`
	// NonreptPositionsShortOther is other non-reportable short positions.
	NonreptPositionsShortOther int64 `json:"nonreptPositionsShortOther"`
	// ChangeInOpenInterestAll is change in total open interest.
	ChangeInOpenInterestAll int64 `json:"changeInOpenInterestAll"`
	// ChangeInNoncommLongAll is change in non-commercial long positions.
	ChangeInNoncommLongAll int64 `json:"changeInNoncommLongAll"`
	// ChangeInNoncommShortAll is change in non-commercial short positions.
	ChangeInNoncommShortAll int64 `json:"changeInNoncommShortAll"`
	// ChangeInNoncommSpeadAll is change in non-commercial spread positions.
	ChangeInNoncommSpeadAll int64 `json:"changeInNoncommSpeadAll"`
	// ChangeInCommLongAll is change in commercial long positions.
	ChangeInCommLongAll int64 `json:"changeInCommLongAll"`
	// ChangeInCommShortAll is change in commercial short positions.
	ChangeInCommShortAll int64 `json:"changeInCommShortAll"`
	// ChangeInTotReptLongAll is change in total reportable long positions.
	ChangeInTotReptLongAll int64 `json:"changeInTotReptLongAll"`
	// ChangeInTotReptShortAll is change in total reportable short positions.
	ChangeInTotReptShortAll int64 `json:"changeInTotReptShortAll"`
	// ChangeInNonreptLongAll is change in non-reportable long positions.
	ChangeInNonreptLongAll int64 `json:"changeInNonreptLongAll"`
	// ChangeInNonreptShortAll is change in non-reportable short positions.
	ChangeInNonreptShortAll int64 `json:"changeInNonreptShortAll"`
	// PctOfOpenInterestAll is percentage of open interest.
	PctOfOpenInterestAll float64 `json:"pctOfOpenInterestAll"`
	// PctOfOiNoncommLongAll is percentage of non-commercial long positions.
	PctOfOiNoncommLongAll float64 `json:"pctOfOiNoncommLongAll"`
	// PctOfOiNoncommShortAll is percentage of non-commercial short positions.
	PctOfOiNoncommShortAll float64 `json:"pctOfOiNoncommShortAll"`
	// PctOfOiNoncommSpreadAll is percentage of non-commercial spread positions.
	PctOfOiNoncommSpreadAll float64 `json:"pctOfOiNoncommSpreadAll"`
	// PctOfOiCommLongAll is percentage of commercial long positions.
	PctOfOiCommLongAll float64 `json:"pctOfOiCommLongAll"`
	// PctOfOiCommShortAll is percentage of commercial short positions.
	PctOfOiCommShortAll float64 `json:"pctOfOiCommShortAll"`
	// PctOfOiTotReptLongAll is percentage of total reportable long positions.
	PctOfOiTotReptLongAll float64 `json:"pctOfOiTotReptLongAll"`
	// PctOfOiTotReptShortAll is percentage of total reportable short positions.
	PctOfOiTotReptShortAll float64 `json:"pctOfOiTotReptShortAll"`
	// PctOfOiNonreptLongAll is percentage of non-reportable long positions.
	PctOfOiNonreptLongAll float64 `json:"pctOfOiNonreptLongAll"`
	// PctOfOiNonreptShortAll is percentage of non-reportable short positions.
	PctOfOiNonreptShortAll float64 `json:"pctOfOiNonreptShortAll"`
	// PctOfOpenInterestOl is percentage of old open interest.
	PctOfOpenInterestOl float64 `json:"pctOfOpenInterestOl"`
	// PctOfOiNoncommLongOl is percentage of old non-commercial long positions.
	PctOfOiNoncommLongOl float64 `json:"pctOfOiNoncommLongOl"`
	// PctOfOiNoncommShortOl is percentage of old non-commercial short positions.
	PctOfOiNoncommShortOl float64 `json:"pctOfOiNoncommShortOl"`
	// PctOfOiNoncommSpreadOl is percentage of old non-commercial spread positions.
	PctOfOiNoncommSpreadOl float64 `json:"pctOfOiNoncommSpreadOl"`
	// PctOfOiCommLongOl is percentage of old commercial long positions.
	PctOfOiCommLongOl float64 `json:"pctOfOiCommLongOl"`
	// PctOfOiCommShortOl is percentage of old commercial short positions.
	PctOfOiCommShortOl float64 `json:"pctOfOiCommShortOl"`
	// PctOfOiTotReptLongOl is percentage of old total reportable long positions.
	PctOfOiTotReptLongOl float64 `json:"pctOfOiTotReptLongOl"`
	// PctOfOiTotReptShortOl is percentage of old total reportable short positions.
	PctOfOiTotReptShortOl float64 `json:"pctOfOiTotReptShortOl"`
	// PctOfOiNonreptLongOl is percentage of old non-reportable long positions.
	PctOfOiNonreptLongOl float64 `json:"pctOfOiNonreptLongOl"`
	// PctOfOiNonreptShortOl is percentage of old non-reportable short positions.
	PctOfOiNonreptShortOl float64 `json:"pctOfOiNonreptShortOl"`
	// PctOfOpenInterestOther is percentage of other open interest.
	PctOfOpenInterestOther float64 `json:"pctOfOpenInterestOther"`
	// PctOfOiNoncommLongOther is percentage of other non-commercial long positions.
	PctOfOiNoncommLongOther float64 `json:"pctOfOiNoncommLongOther"`
	// PctOfOiNoncommShortOther is percentage of other non-commercial short positions.
	PctOfOiNoncommShortOther float64 `json:"pctOfOiNoncommShortOther"`
	// PctOfOiNoncommSpreadOther is percentage of other non-commercial spread positions.
	PctOfOiNoncommSpreadOther float64 `json:"pctOfOiNoncommSpreadOther"`
	// PctOfOiCommLongOther is percentage of other commercial long positions.
	PctOfOiCommLongOther float64 `json:"pctOfOiCommLongOther"`
	// PctOfOiCommShortOther is percentage of other commercial short positions.
	PctOfOiCommShortOther float64 `json:"pctOfOiCommShortOther"`
	// PctOfOiTotReptLongOther is percentage of other total reportable long positions.
	PctOfOiTotReptLongOther float64 `json:"pctOfOiTotReptLongOther"`
	// PctOfOiTotReptShortOther is percentage of other total reportable short positions.
	PctOfOiTotReptShortOther float64 `json:"pctOfOiTotReptShortOther"`
	// PctOfOiNonreptLongOther is percentage of other non-reportable long positions.
	PctOfOiNonreptLongOther float64 `json:"pctOfOiNonreptLongOther"`
	// PctOfOiNonreptShortOther is percentage of other non-reportable short positions.
	PctOfOiNonreptShortOther float64 `json:"pctOfOiNonreptShortOther"`
	// TradersTotAll is total number of traders.
	TradersTotAll int `json:"tradersTotAll"`
	// TradersNoncommLongAll is number of non-commercial long traders.
	TradersNoncommLongAll int `json:"tradersNoncommLongAll"`
	// TradersNoncommShortAll is number of non-commercial short traders.
	TradersNoncommShortAll int `json:"tradersNoncommShortAll"`
	// TradersNoncommSpreadAll is number of non-commercial spread traders.
	TradersNoncommSpreadAll int `json:"tradersNoncommSpreadAll"`
	// TradersCommLongAll is number of commercial long traders.
	TradersCommLongAll int `json:"tradersCommLongAll"`
	// TradersCommShortAll is number of commercial short traders.
	TradersCommShortAll int `json:"tradersCommShortAll"`
	// TradersTotReptLongAll is number of total reportable long traders.
	TradersTotReptLongAll int `json:"tradersTotReptLongAll"`
	// TradersTotReptShortAll is number of total reportable short traders.
	TradersTotReptShortAll int `json:"tradersTotReptShortAll"`
	// TradersTotOl is total number of old traders.
	TradersTotOl int `json:"tradersTotOl"`
	// TradersNoncommLongOl is number of old non-commercial long traders.
	TradersNoncommLongOl int `json:"tradersNoncommLongOl"`
	// TradersNoncommShortOl is number of old non-commercial short traders.
	TradersNoncommShortOl int `json:"tradersNoncommShortOl"`
	// TradersNoncommSpeadOl is number of old non-commercial spread traders.
	TradersNoncommSpeadOl int `json:"tradersNoncommSpeadOl"`
	// TradersCommLongOl is number of old commercial long traders.
	TradersCommLongOl int `json:"tradersCommLongOl"`
	// TradersCommShortOl is number of old commercial short traders.
	TradersCommShortOl int `json:"tradersCommShortOl"`
	// TradersTotReptLongOl is number of old total reportable long traders.
	TradersTotReptLongOl int `json:"tradersTotReptLongOl"`
	// TradersTotReptShortOl is number of old total reportable short traders.
	TradersTotReptShortOl int `json:"tradersTotReptShortOl"`
	// TradersTotOther is total number of other traders.
	TradersTotOther int `json:"tradersTotOther"`
	// TradersNoncommLongOther is number of other non-commercial long traders.
	TradersNoncommLongOther int `json:"tradersNoncommLongOther"`
	// TradersNoncommShortOther is number of other non-commercial short traders.
	TradersNoncommShortOther int `json:"tradersNoncommShortOther"`
	// TradersNoncommSpreadOther is number of other non-commercial spread traders.
	TradersNoncommSpreadOther int `json:"tradersNoncommSpreadOther"`
	// TradersCommLongOther is number of other commercial long traders.
	TradersCommLongOther int `json:"tradersCommLongOther"`
	// TradersCommShortOther is number of other commercial short traders.
	TradersCommShortOther int `json:"tradersCommShortOther"`
	// TradersTotReptLongOther is number of other total reportable long traders.
	TradersTotReptLongOther int `json:"tradersTotReptLongOther"`
	// TradersTotReptShortOther is number of other total reportable short traders.
	TradersTotReptShortOther int `json:"tradersTotReptShortOther"`
	// ConcGrossLe4TdrLongAll is concentration gross <= 4 traders long.
	ConcGrossLe4TdrLongAll float64 `json:"concGrossLe4TdrLongAll"`
	// ConcGrossLe4TdrShortAll is concentration gross <= 4 traders short.
	ConcGrossLe4TdrShortAll float64 `json:"concGrossLe4TdrShortAll"`
	// ConcGrossLe8TdrLongAll is concentration gross <= 8 traders long.
	ConcGrossLe8TdrLongAll float64 `json:"concGrossLe8TdrLongAll"`
	// ConcGrossLe8TdrShortAll is concentration gross <= 8 traders short.
	ConcGrossLe8TdrShortAll float64 `json:"concGrossLe8TdrShortAll"`
	// ConcNetLe4TdrLongAll is concentration net <= 4 traders long.
	ConcNetLe4TdrLongAll float64 `json:"concNetLe4TdrLongAll"`
	// ConcNetLe4TdrShortAll is concentration net <= 4 traders short.
	ConcNetLe4TdrShortAll float64 `json:"concNetLe4TdrShortAll"`
	// ConcNetLe8TdrLongAll is concentration net <= 8 traders long.
	ConcNetLe8TdrLongAll float64 `json:"concNetLe8TdrLongAll"`
	// ConcNetLe8TdrShortAll is concentration net <= 8 traders short.
	ConcNetLe8TdrShortAll float64 `json:"concNetLe8TdrShortAll"`
	// ConcGrossLe4TdrLongOl is old concentration gross <= 4 traders long.
	ConcGrossLe4TdrLongOl float64 `json:"concGrossLe4TdrLongOl"`
	// ConcGrossLe4TdrShortOl is old concentration gross <= 4 traders short.
	ConcGrossLe4TdrShortOl float64 `json:"concGrossLe4TdrShortOl"`
	// ConcGrossLe8TdrLongOl is old concentration gross <= 8 traders long.
	ConcGrossLe8TdrLongOl float64 `json:"concGrossLe8TdrLongOl"`
	// ConcGrossLe8TdrShortOl is old concentration gross <= 8 traders short.
	ConcGrossLe8TdrShortOl float64 `json:"concGrossLe8TdrShortOl"`
	// ConcNetLe4TdrLongOl is old concentration net <= 4 traders long.
	ConcNetLe4TdrLongOl float64 `json:"concNetLe4TdrLongOl"`
	// ConcNetLe4TdrShortOl is old concentration net <= 4 traders short.
	ConcNetLe4TdrShortOl float64 `json:"concNetLe4TdrShortOl"`
	// ConcNetLe8TdrLongOl is old concentration net <= 8 traders long.
	ConcNetLe8TdrLongOl float64 `json:"concNetLe8TdrLongOl"`
	// ConcNetLe8TdrShortOl is old concentration net <= 8 traders short.
	ConcNetLe8TdrShortOl float64 `json:"concNetLe8TdrShortOl"`
	// ConcGrossLe4TdrLongOther is other concentration gross <= 4 traders long.
	ConcGrossLe4TdrLongOther float64 `json:"concGrossLe4TdrLongOther"`
	// ConcGrossLe4TdrShortOther is other concentration gross <= 4 traders short.
	ConcGrossLe4TdrShortOther float64 `json:"concGrossLe4TdrShortOther"`
	// ConcGrossLe8TdrLongOther is other concentration gross <= 8 traders long.
	ConcGrossLe8TdrLongOther float64 `json:"concGrossLe8TdrLongOther"`
	// ConcGrossLe8TdrShortOther is other concentration gross <= 8 traders short.
	ConcGrossLe8TdrShortOther float64 `json:"concGrossLe8TdrShortOther"`
	// ConcNetLe4TdrLongOther is other concentration net <= 4 traders long.
	ConcNetLe4TdrLongOther float64 `json:"concNetLe4TdrLongOther"`
	// ConcNetLe4TdrShortOther is other concentration net <= 4 traders short.
	ConcNetLe4TdrShortOther float64 `json:"concNetLe4TdrShortOther"`
	// ConcNetLe8TdrLongOther is other concentration net <= 8 traders long.
	ConcNetLe8TdrLongOther float64 `json:"concNetLe8TdrLongOther"`
	// ConcNetLe8TdrShortOther is other concentration net <= 8 traders short.
	ConcNetLe8TdrShortOther float64 `json:"concNetLe8TdrShortOther"`
	// ContractUnits describes the contract unit size.
	ContractUnits string `json:"contractUnits"`
}

// COTAnalysisResponse represents COT market sentiment analysis.
type COTAnalysisResponse struct {
	// Symbol is the commodity or futures symbol.
	Symbol string `json:"symbol"`
	// Date is the analysis date.
	Date string `json:"date"`
	// Name is the commodity name.
	Name string `json:"name"`
	// Sector is the market sector.
	Sector string `json:"sector"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
	// CurrentLongMarketSituation is current long position percentage.
	CurrentLongMarketSituation float64 `json:"currentLongMarketSituation"`
	// CurrentShortMarketSituation is current short position percentage.
	CurrentShortMarketSituation float64 `json:"currentShortMarketSituation"`
	// MarketSituation describes current market sentiment.
	MarketSituation string `json:"marketSituation"`
	// PreviousLongMarketSituation is previous long position percentage.
	PreviousLongMarketSituation float64 `json:"previousLongMarketSituation"`
	// PreviousShortMarketSituation is previous short position percentage.
	PreviousShortMarketSituation float64 `json:"previousShortMarketSituation"`
	// PreviousMarketSituation describes previous market sentiment.
	PreviousMarketSituation string `json:"previousMarketSituation"`
	// NetPosition is the net position value.
	NetPosition int64 `json:"netPostion"`
	// PreviousNetPosition is the previous net position value.
	PreviousNetPosition int64 `json:"previousNetPosition"`
	// ChangeInNetPosition is the percentage change in net position.
	ChangeInNetPosition float64 `json:"changeInNetPosition"`
	// MarketSentiment describes the trend in market sentiment.
	MarketSentiment string `json:"marketSentiment"`
	// ReversalTrend indicates if there's a trend reversal.
	ReversalTrend bool `json:"reversalTrend"`
}

// COTListResponse represents a COT report commodity listing.
type COTListResponse struct {
	// Symbol is the commodity or futures symbol.
	Symbol string `json:"symbol"`
	// Name is the commodity name.
	Name string `json:"name"`
}

// COTReport retrieves Commitments of Traders reports.
// Returns detailed position data for commodities and futures.
//
// Endpoint: /commitment-of-traders-report
func (c *Client) COTReport(params COTReportParams) ([]COTReportResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/commitment-of-traders-report"

	var result []COTReportResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// COTAnalysis retrieves COT market sentiment analysis.
// Returns bullish/bearish sentiment and trend analysis.
//
// Endpoint: /commitment-of-traders-analysis
func (c *Client) COTAnalysis(params COTAnalysisParams) ([]COTAnalysisResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/commitment-of-traders-analysis"

	var result []COTAnalysisResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// COTList retrieves available COT report symbols.
// Returns list of commodities and futures with COT data.
//
// Endpoint: /commitment-of-traders-list
func (c *Client) COTList() ([]COTListResponse, error) {
	pathName := "/commitment-of-traders-list"

	var result []COTListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}