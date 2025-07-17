package go_fmp

import (
	"fmt"
	"time"
)

// CryptoHistoricalFullParams represents the parameters for the Crypto Historical Full API
type CryptoHistoricalFullParams struct {
	Symbol string  `json:"symbol"` // Required: Crypto symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (YYYY-MM-DD format)
	To     *string `json:"to"`     // Optional: End date (YYYY-MM-DD format)
}

// CryptoHistoricalFullResponse represents the response from the Crypto Historical Full API
type CryptoHistoricalFullResponse struct {
	Date          time.Time  `json:"date"`
	Open          float64 `json:"open"`
	Low           float64 `json:"low"`
	High          float64 `json:"high"`
	Close         float64 `json:"close"`
	AdjClose      float64 `json:"adjClose"`
	Volume        float64 `json:"volume"`
	UnadjustedVolume float64 `json:"unadjustedVolume"`
	Change        float64 `json:"change"`
	ChangePercent float64 `json:"changePercent"`
	Vwap          float64 `json:"vwap"`
	Label         string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
}

// CryptoHistoricalFull retrieves comprehensive historical price data for cryptocurrencies
func (c *Client) CryptoHistoricalFull(params CryptoHistoricalFullParams) ([]CryptoHistoricalFullResponse, error) {
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

	url := fmt.Sprintf("%s/historical-price-full/%s", c.BaseURL, params.Symbol)
	var result []CryptoHistoricalFullResponse
	err := c.doRequest(url, urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}