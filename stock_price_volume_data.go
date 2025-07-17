package go_fmp

import (
	"fmt"
	"time"
)

// StockPriceVolumeDataParams represents the parameters for the Stock Price Volume Data API
type StockPriceVolumeDataParams struct {
	Symbol string  `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// StockPriceVolumeDataResponse represents the response from the Stock Price Volume Data API
type StockPriceVolumeDataResponse struct {
	Date          time.Time  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	AdjClose      float64 `json:"adjClose"`
	Volume        int64   `json:"volume"`
	UnadjustedVol int64   `json:"unadjustedVolume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Vwap          float64 `json:"vwap"`
	Label         string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
}

// StockPriceVolumeData retrieves full price and volume data for any stock symbol
func (c *Client) StockPriceVolumeData(params StockPriceVolumeDataParams) ([]StockPriceVolumeDataResponse, error) {
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

	url := c.BaseURL + "/historical-price-eod/full"
	var result []StockPriceVolumeDataResponse
	if err := c.doRequest(url, urlParams, &result); err != nil {
		return nil, err
	}
	return result, nil
}
