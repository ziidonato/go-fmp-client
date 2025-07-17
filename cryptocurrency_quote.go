package go_fmp

// CryptocurrencyQuoteParams represents the parameters for the Cryptocurrency Quote API
type CryptocurrencyQuoteParams struct {
	Symbol string `json:"symbol"` // Required: Cryptocurrency symbol (e.g., "BTCUSD")
}

// CryptocurrencyQuoteResponse represents the response from the Cryptocurrency Quote API
type CryptocurrencyQuoteResponse struct {
	Symbol           string  `json:"symbol"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	ChangePercentage float64 `json:"changePercentage"`
	Change           float64 `json:"change"`
	Volume           int64   `json:"volume"`
	DayLow           float64 `json:"dayLow"`
	DayHigh          float64 `json:"dayHigh"`
	YearHigh         float64 `json:"yearHigh"`
	YearLow          float64 `json:"yearLow"`
	MarketCap        *int64  `json:"marketCap"`
	PriceAvg50       float64 `json:"priceAvg50"`
	PriceAvg200      float64 `json:"priceAvg200"`
	Exchange Exchange `json:"exchange"`
	Open             float64 `json:"open"`
	PreviousClose    float64 `json:"previousClose"`
	Timestamp        int64   `json:"timestamp"`
}

// CryptocurrencyQuote retrieves real-time quotes for cryptocurrencies
func (c *Client) CryptocurrencyQuote(symbol string) ([]CryptocurrencyQuoteResponse, error) {
	url := c.BaseURL + "/quote"
	params := map[string]string{"symbol": symbol}
	var result []CryptocurrencyQuoteResponse
	if err := c.doRequest(url, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
