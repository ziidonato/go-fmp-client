package go_fmp

import (
	"fmt"
)

// COTReportParams represents the parameters for the COT Report API
type COTReportParams struct {
	Symbol *string `json:"symbol"` // Optional: Symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Required: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Required: End date (e.g., "2024-03-01")
}

// COTReportResponse represents the response from the COT Report API
type COTReportResponse struct {
	Symbol                      string  `json:"symbol"`
	Date                        string  `json:"date"`
	Name                        string  `json:"name"`
	Sector                      string  `json:"sector"`
	MarketAndExchangeNames      string  `json:"marketAndExchangeNames"`
	CFTCContractMarketCode      string  `json:"cftcContractMarketCode"`
	CFTCMarketCode              string  `json:"cftcMarketCode"`
	CFTCRegionCode              string  `json:"cftcRegionCode"`
	CFTCCommodityCode           string  `json:"cftcCommodityCode"`
	OpenInterestAll             int     `json:"openInterestAll"`
	NoncommPositionsLongAll     int     `json:"noncommPositionsLongAll"`
	NoncommPositionsShortAll    int     `json:"noncommPositionsShortAll"`
	NoncommPositionsSpreadAll   int     `json:"noncommPositionsSpreadAll"`
	CommPositionsLongAll        int     `json:"commPositionsLongAll"`
	CommPositionsShortAll       int     `json:"commPositionsShortAll"`
	TotReptPositionsLongAll     int     `json:"totReptPositionsLongAll"`
	TotReptPositionsShortAll    int     `json:"totReptPositionsShortAll"`
	NonreptPositionsLongAll     int     `json:"nonreptPositionsLongAll"`
	NonreptPositionsShortAll    int     `json:"nonreptPositionsShortAll"`
	OpenInterestOld             int     `json:"openInterestOld"`
	NoncommPositionsLongOld     int     `json:"noncommPositionsLongOld"`
	NoncommPositionsShortOld    int     `json:"noncommPositionsShortOld"`
	NoncommPositionsSpreadOld   int     `json:"noncommPositionsSpreadOld"`
	CommPositionsLongOld        int     `json:"commPositionsLongOld"`
	CommPositionsShortOld       int     `json:"commPositionsShortOld"`
	TotReptPositionsLongOld     int     `json:"totReptPositionsLongOld"`
	TotReptPositionsShortOld    int     `json:"totReptPositionsShortOld"`
	NonreptPositionsLongOld     int     `json:"nonreptPositionsLongOld"`
	NonreptPositionsShortOld    int     `json:"nonreptPositionsShortOld"`
	OpenInterestOther           int     `json:"openInterestOther"`
	NoncommPositionsLongOther   int     `json:"noncommPositionsLongOther"`
	NoncommPositionsShortOther  int     `json:"noncommPositionsShortOther"`
	NoncommPositionsSpreadOther int     `json:"noncommPositionsSpreadOther"`
	CommPositionsLongOther      int     `json:"commPositionsLongOther"`
	CommPositionsShortOther     int     `json:"commPositionsShortOther"`
	TotReptPositionsLongOther   int     `json:"totReptPositionsLongOther"`
	TotReptPositionsShortOther  int     `json:"totReptPositionsShortOther"`
	NonreptPositionsLongOther   int     `json:"nonreptPositionsLongOther"`
	NonreptPositionsShortOther  int     `json:"nonreptPositionsShortOther"`
	ChangeInOpenInterestAll     int     `json:"changeInOpenInterestAll"`
	ChangeInNoncommLongAll      int     `json:"changeInNoncommLongAll"`
	ChangeInNoncommShortAll     int     `json:"changeInNoncommShortAll"`
	ChangeInNoncommSpeadAll     int     `json:"changeInNoncommSpeadAll"`
	ChangeInCommLongAll         int     `json:"changeInCommLongAll"`
	ChangeInCommShortAll        int     `json:"changeInCommShortAll"`
	ChangeInTotReptLongAll      int     `json:"changeInTotReptLongAll"`
	ChangeInTotReptShortAll     int     `json:"changeInTotReptShortAll"`
	ChangeInNonreptLongAll      int     `json:"changeInNonreptLongAll"`
	ChangeInNonreptShortAll     int     `json:"changeInNonreptShortAll"`
	PctOfOpenInterestAll        float64 `json:"pctOfOpenInterestAll"`
	PctOfOiNoncommLongAll       float64 `json:"pctOfOiNoncommLongAll"`
	PctOfOiNoncommShortAll      float64 `json:"pctOfOiNoncommShortAll"`
	PctOfOiNoncommSpreadAll     float64 `json:"pctOfOiNoncommSpreadAll"`
	PctOfOiCommLongAll          float64 `json:"pctOfOiCommLongAll"`
	PctOfOiCommShortAll         float64 `json:"pctOfOiCommShortAll"`
	PctOfOiTotReptLongAll       float64 `json:"pctOfOiTotReptLongAll"`
	PctOfOiTotReptShortAll      float64 `json:"pctOfOiTotReptShortAll"`
	PctOfOiNonreptLongAll       float64 `json:"pctOfOiNonreptLongAll"`
	PctOfOiNonreptShortAll      float64 `json:"pctOfOiNonreptShortAll"`
	PctOfOpenInterestOl         float64 `json:"pctOfOpenInterestOl"`
	PctOfOiNoncommLongOl        float64 `json:"pctOfOiNoncommLongOl"`
	PctOfOiNoncommShortOl       float64 `json:"pctOfOiNoncommShortOl"`
	PctOfOiNoncommSpreadOl      float64 `json:"pctOfOiNoncommSpreadOl"`
	PctOfOiCommLongOl           float64 `json:"pctOfOiCommLongOl"`
	PctOfOiCommShortOl          float64 `json:"pctOfOiCommShortOl"`
	PctOfOiTotReptLongOl        float64 `json:"pctOfOiTotReptLongOl"`
	PctOfOiTotReptShortOl       float64 `json:"pctOfOiTotReptShortOl"`
	PctOfOiNonreptLongOl        float64 `json:"pctOfOiNonreptLongOl"`
	PctOfOiNonreptShortOl       float64 `json:"pctOfOiNonreptShortOl"`
	PctOfOpenInterestOther      float64 `json:"pctOfOpenInterestOther"`
	PctOfOiNoncommLongOther     float64 `json:"pctOfOiNoncommLongOther"`
	PctOfOiNoncommShortOther    float64 `json:"pctOfOiNoncommShortOther"`
	PctOfOiNoncommSpreadOther   float64 `json:"pctOfOiNoncommSpreadOther"`
	PctOfOiCommLongOther        float64 `json:"pctOfOiCommLongOther"`
	PctOfOiCommShortOther       float64 `json:"pctOfOiCommShortOther"`
	PctOfOiTotReptLongOther     float64 `json:"pctOfOiTotReptLongOther"`
	PctOfOiTotReptShortOther    float64 `json:"pctOfOiTotReptShortOther"`
	PctOfOiNonreptLongOther     float64 `json:"pctOfOiNonreptLongOther"`
	PctOfOiNonreptShortOther    float64 `json:"pctOfOiNonreptShortOther"`
	TradersTotAll               int     `json:"tradersTotAll"`
	TradersNoncommLongAll       int     `json:"tradersNoncommLongAll"`
	TradersNoncommShortAll      int     `json:"tradersNoncommShortAll"`
	TradersNoncommSpreadAll     int     `json:"tradersNoncommSpreadAll"`
	TradersCommLongAll          int     `json:"tradersCommLongAll"`
	TradersCommShortAll         int     `json:"tradersCommShortAll"`
	TradersTotReptLongAll       int     `json:"tradersTotReptLongAll"`
	TradersTotReptShortAll      int     `json:"tradersTotReptShortAll"`
	TradersTotOl                int     `json:"tradersTotOl"`
	TradersNoncommLongOl        int     `json:"tradersNoncommLongOl"`
	TradersNoncommShortOl       int     `json:"tradersNoncommShortOl"`
	TradersNoncommSpeadOl       int     `json:"tradersNoncommSpeadOl"`
	TradersCommLongOl           int     `json:"tradersCommLongOl"`
	TradersCommShortOl          int     `json:"tradersCommShortOl"`
	TradersTotReptLongOl        int     `json:"tradersTotReptLongOl"`
	TradersTotReptShortOl       int     `json:"tradersTotReptShortOl"`
	TradersTotOther             int     `json:"tradersTotOther"`
	TradersNoncommLongOther     int     `json:"tradersNoncommLongOther"`
	TradersNoncommShortOther    int     `json:"tradersNoncommShortOther"`
	TradersNoncommSpreadOther   int     `json:"tradersNoncommSpreadOther"`
	TradersCommLongOther        int     `json:"tradersCommLongOther"`
	TradersCommShortOther       int     `json:"tradersCommShortOther"`
	TradersTotReptLongOther     int     `json:"tradersTotReptLongOther"`
	TradersTotReptShortOther    int     `json:"tradersTotReptShortOther"`
	ConcGrossLe4TdrLongAll      float64 `json:"concGrossLe4TdrLongAll"`
	ConcGrossLe4TdrShortAll     float64 `json:"concGrossLe4TdrShortAll"`
	ConcGrossLe8TdrLongAll      float64 `json:"concGrossLe8TdrLongAll"`
	ConcGrossLe8TdrShortAll     float64 `json:"concGrossLe8TdrShortAll"`
	ConcNetLe4TdrLongAll        float64 `json:"concNetLe4TdrLongAll"`
	ConcNetLe4TdrShortAll       float64 `json:"concNetLe4TdrShortAll"`
	ConcNetLe8TdrLongAll        float64 `json:"concNetLe8TdrLongAll"`
	ConcNetLe8TdrShortAll       float64 `json:"concNetLe8TdrShortAll"`
	ConcGrossLe4TdrLongOl       float64 `json:"concGrossLe4TdrLongOl"`
	ConcGrossLe4TdrShortOl      float64 `json:"concGrossLe4TdrShortOl"`
	ConcGrossLe8TdrLongOl       float64 `json:"concGrossLe8TdrLongOl"`
	ConcGrossLe8TdrShortOl      float64 `json:"concGrossLe8TdrShortOl"`
	ConcNetLe4TdrLongOl         float64 `json:"concNetLe4TdrLongOl"`
	ConcNetLe4TdrShortOl        float64 `json:"concNetLe4TdrShortOl"`
	ConcNetLe8TdrLongOl         float64 `json:"concNetLe8TdrLongOl"`
	ConcNetLe8TdrShortOl        float64 `json:"concNetLe8TdrShortOl"`
	ConcGrossLe4TdrLongOther    float64 `json:"concGrossLe4TdrLongOther"`
	ConcGrossLe4TdrShortOther   float64 `json:"concGrossLe4TdrShortOther"`
	ConcGrossLe8TdrLongOther    float64 `json:"concGrossLe8TdrLongOther"`
	ConcGrossLe8TdrShortOther   float64 `json:"concGrossLe8TdrShortOther"`
	ConcNetLe4TdrLongOther      float64 `json:"concNetLe4TdrLongOther"`
	ConcNetLe4TdrShortOther     float64 `json:"concNetLe4TdrShortOther"`
	ConcNetLe8TdrLongOther      float64 `json:"concNetLe8TdrLongOther"`
	ConcNetLe8TdrShortOther     float64 `json:"concNetLe8TdrShortOther"`
	ContractUnits               string  `json:"contractUnits"`
}

// COTReport retrieves comprehensive Commitment of Traders (COT) reports
func (c *Client) COTReport(params COTReportParams) ([]COTReportResponse, error) {
	if params.From == nil {
		return nil, fmt.Errorf("from parameter is required")
	}

	if params.To == nil {
		return nil, fmt.Errorf("to parameter is required")
	}

	urlParams := map[string]string{
		"from": *params.From,
		"to":   *params.To,
	}

	if params.Symbol != nil {
		urlParams["symbol"] = *params.Symbol
	}

	var result []COTReportResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/commitment-of-traders-report", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
