package go_fmp

import (
	"fmt"
	"time"
)

// StockPriceChartsParams represents the parameters for the Stock Price Charts API
type StockPriceChartsParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// StockPriceChartsResponse represents the response from the Stock Price Charts API
type StockPriceChartsResponse struct {
	Date     time.Time  `json:"date"`
	Open     float64 `json:"open"`
	Low      float64 `json:"low"`
	High     float64 `json:"high"`
	Close    float64 `json:"close"`
	AdjClose float64 `json:"adjClose"`
	Volume   int64   `json:"volume"`
	UnadjustedVolume int64 `json:"unadjustedVolume"`
	Change   float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Vwap     float64 `json:"vwap"`
	Label    string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
}

// UnadjustedStockPriceChart retrieves stock price and volume data without adjustments for stock splits
func (c *Client) UnadjustedStockPriceChart(params StockPriceChartsParams) ([]StockPriceChartsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.From != nil {
		urlParams["from"] = *params.From
	}

	if params.To != nil {
		urlParams["to"] = *params.To
	}

	url := c.BaseURL + "/historical-price-eod/non-split-adjusted"
	var result []StockPriceChartsResponse
	if err := c.doRequest(url, urlParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DividendAdjustedPriceChart retrieves stock price and volume data with dividend adjustments
func (c *Client) DividendAdjustedPriceChart(params StockPriceChartsParams) ([]StockPriceChartsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	if params.From != nil {
		urlParams["from"] = *params.From
	}

	if params.To != nil {
		urlParams["to"] = *params.To
	}

	var result []StockPriceChartsResponse
	err := c.doRequest(c.BaseURL+"/historical-price-eod/dividend-adjusted", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
