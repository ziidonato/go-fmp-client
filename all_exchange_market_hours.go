package go_fmp

// AllExchangeMarketHoursResponse represents the response from the all exchange market hours API
type AllExchangeMarketHoursResponse struct {
	// Note: The exact structure will depend on the actual API response
	// This is a placeholder structure that should be updated based on actual response
	Exchange string `json:"exchange"`
	// Add other fields as needed based on actual API response
}

// AllExchangeMarketHours retrieves market hours for all exchanges
func (c *Client) AllExchangeMarketHours() ([]AllExchangeMarketHoursResponse, error) {
	url := "https://financialmodelingprep.com/stable/all-exchange-market-hours"

	var result []AllExchangeMarketHoursResponse
	if err := c.doRequest(url, map[string]string{}, &result); err != nil {
		return nil, err
	}
	return result, nil
}
