package go_fmp

import (
	"time"
	"fmt"
)

// CryptocurrencyChart1MinParams represents the parameters for the 1-Minute Cryptocurrency Chart API
type CryptocurrencyChart1MinParams struct {
	Symbol string  `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
	From   *string `json:"from"`   // Optional: Start date (e.g., "2024-01-01")
	To     *string `json:"to"`     // Optional: End date (e.g., "2024-03-01")
}

type CryptocurrencyChart5MinParams = CryptocurrencyChart1MinParams

type CryptocurrencyChart1HourParams = CryptocurrencyChart1MinParams

// CryptocurrencyChart1MinResponse represents the response from the 1-Minute Cryptocurrency Chart API
type CryptocurrencyChart1MinResponse struct {
	Date time.Time `json:"date"`
	Open   float64 `json:"open"`
	Low    float64 `json:"low"`
	High   float64 `json:"high"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

type CryptocurrencyChart5MinResponse = CryptocurrencyChart1MinResponse

type CryptocurrencyChart1HourResponse = CryptocurrencyChart1MinResponse

// CryptocurrencyChart1Min retrieves real-time, 1-minute interval price data for cryptocurrencies
func (c *Client) CryptocurrencyChart1Min(params CryptocurrencyChart1MinParams) ([]CryptocurrencyChart1MinResponse, error) {
	return c.getCryptocurrencyChartInterval(params, "1min")
}

// CryptocurrencyChart5Min retrieves real-time, 5-minute interval price data for cryptocurrencies
func (c *Client) CryptocurrencyChart5Min(params CryptocurrencyChart5MinParams) ([]CryptocurrencyChart5MinResponse, error) {
	return c.getCryptocurrencyChartInterval(params, "5min")
}

// CryptocurrencyChart1Hour retrieves real-time, 1-hour interval price data for cryptocurrencies
func (c *Client) CryptocurrencyChart1Hour(params CryptocurrencyChart1HourParams) ([]CryptocurrencyChart1HourResponse, error) {
	return c.getCryptocurrencyChartInterval(params, "1hour")
}

func (c *Client) getCryptocurrencyChartInterval(params CryptocurrencyChart1MinParams, interval string) ([]CryptocurrencyChart1MinResponse, error) {
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
	var result []CryptocurrencyChart1MinResponse
	url := c.BaseURL + "/historical-chart/" + interval
	err := c.doRequest(url, urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	return result, nil
}
