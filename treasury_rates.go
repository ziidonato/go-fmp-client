package go_fmp

// TreasuryRatesResponse represents the response from the Treasury Rates API
type TreasuryRatesResponse struct {
	Date      string  `json:"date"`
	BC_1MONTH float64 `json:"BC_1MONTH"`
	BC_2MONTH float64 `json:"BC_2MONTH"`
	BC_3MONTH float64 `json:"BC_3MONTH"`
	BC_6MONTH float64 `json:"BC_6MONTH"`
	BC_1YEAR  float64 `json:"BC_1YEAR"`
	BC_2YEAR  float64 `json:"BC_2YEAR"`
	BC_3YEAR  float64 `json:"BC_3YEAR"`
	BC_5YEAR  float64 `json:"BC_5YEAR"`
	BC_7YEAR  float64 `json:"BC_7YEAR"`
	BC_10YEAR float64 `json:"BC_10YEAR"`
	BC_20YEAR float64 `json:"BC_20YEAR"`
	BC_30YEAR float64 `json:"BC_30YEAR"`
}

// TreasuryRates retrieves real-time and historical Treasury rates for all maturities
func (c *Client) TreasuryRates() ([]TreasuryRatesResponse, error) {
	var result []TreasuryRatesResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/treasury-rates", map[string]string{}, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}
