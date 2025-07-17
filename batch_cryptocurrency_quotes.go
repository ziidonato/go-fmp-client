package go_fmp

// BatchCryptocurrencyQuotesResponse represents the response from the Batch Cryptocurrency Quotes API
type BatchCryptocurrencyQuotesResponse struct {
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

// BatchCryptocurrencyQuotes retrieves live price data for a wide range of cryptocurrencies in a single request
func (c *Client) BatchCryptocurrencyQuotes() ([]BatchCryptocurrencyQuotesResponse, error) {
	url := c.BaseURL + "/batch-crypto-quotes"
	var result []BatchCryptocurrencyQuotesResponse
	if err := c.doRequest(url, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
