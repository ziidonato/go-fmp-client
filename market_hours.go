package fmp_client

import "time"

// ExchangeMarketHoursParams represents parameters for exchange market hours endpoint.
type ExchangeMarketHoursParams struct {
	// Exchange is the exchange name (required).
	Exchange string `json:"exchange"`
}

// HolidaysByExchangeParams represents parameters for holidays by exchange endpoint.
type HolidaysByExchangeParams struct {
	// Exchange is the exchange name (required).
	Exchange string `json:"exchange"`
	// From is the start date for filtering holidays (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering holidays (optional).
	To *time.Time `json:"to,omitempty"`
}

// MarketHoursResponse represents market hours information for an exchange.
type MarketHoursResponse struct {
	// Exchange is the exchange code.
	Exchange string `json:"exchange"`
	// Name is the full name of the exchange.
	Name string `json:"name"`
	// OpeningHour is the market opening time with timezone.
	OpeningHour string `json:"openingHour"`
	// ClosingHour is the market closing time with timezone.
	ClosingHour string `json:"closingHour"`
	// Timezone is the exchange's timezone.
	Timezone string `json:"timezone"`
	// IsMarketOpen indicates if the market is currently open.
	IsMarketOpen bool `json:"isMarketOpen"`
}

// HolidayResponse represents holiday information for an exchange.
type HolidayResponse struct {
	// Exchange is the exchange code.
	Exchange string `json:"exchange"`
	// Date is the holiday date.
	Date string `json:"date"`
	// Name is the holiday name.
	Name string `json:"name"`
	// IsClosed indicates if the market is fully closed.
	IsClosed bool `json:"isClosed"`
	// AdjOpenTime is the adjusted opening time if partially open.
	AdjOpenTime *string `json:"adjOpenTime"`
	// AdjCloseTime is the adjusted closing time if early close.
	AdjCloseTime *string `json:"adjCloseTime"`
}

// ExchangeMarketHours retrieves trading hours for a specific exchange.
// Returns opening and closing times with timezone information.
//
// Endpoint: /exchange-market-hours
func (c *Client) ExchangeMarketHours(params ExchangeMarketHoursParams) ([]MarketHoursResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/exchange-market-hours"

	var result []MarketHoursResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HolidaysByExchange retrieves holidays and special hours for an exchange.
// Returns dates when the market is closed or has adjusted hours.
//
// Endpoint: /holidays-by-exchange
func (c *Client) HolidaysByExchange(params HolidaysByExchangeParams) ([]HolidayResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/holidays-by-exchange"

	var result []HolidayResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AllExchangeMarketHours retrieves market hours for all exchanges.
// Returns trading times for exchanges worldwide.
//
// Endpoint: /all-exchange-market-hours
func (c *Client) AllExchangeMarketHours() ([]MarketHoursResponse, error) {
	pathName := "/all-exchange-market-hours"

	var result []MarketHoursResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}