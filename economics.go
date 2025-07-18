package fmp_client

import "time"

// TreasuryRatesParams represents parameters for treasury rates endpoint.
type TreasuryRatesParams struct {
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// EconomicIndicatorsParams represents parameters for economic indicators endpoint.
type EconomicIndicatorsParams struct {
	// Name is the indicator name (required).
	Name string `json:"name"`
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// EconomicCalendarParams represents parameters for economic calendar endpoint.
type EconomicCalendarParams struct {
	// From is the start date for filtering (optional).
	From *time.Time `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *time.Time `json:"to,omitempty"`
}

// TreasuryRatesResponse represents treasury rates for various maturities.
type TreasuryRatesResponse struct {
	// Date is the date of the rates.
	Date string `json:"date"`
	// Month1 is the 1-month treasury rate.
	Month1 float64 `json:"month1"`
	// Month2 is the 2-month treasury rate.
	Month2 float64 `json:"month2"`
	// Month3 is the 3-month treasury rate.
	Month3 float64 `json:"month3"`
	// Month6 is the 6-month treasury rate.
	Month6 float64 `json:"month6"`
	// Year1 is the 1-year treasury rate.
	Year1 float64 `json:"year1"`
	// Year2 is the 2-year treasury rate.
	Year2 float64 `json:"year2"`
	// Year3 is the 3-year treasury rate.
	Year3 float64 `json:"year3"`
	// Year5 is the 5-year treasury rate.
	Year5 float64 `json:"year5"`
	// Year7 is the 7-year treasury rate.
	Year7 float64 `json:"year7"`
	// Year10 is the 10-year treasury rate.
	Year10 float64 `json:"year10"`
	// Year20 is the 20-year treasury rate.
	Year20 float64 `json:"year20"`
	// Year30 is the 30-year treasury rate.
	Year30 float64 `json:"year30"`
}

// EconomicIndicatorResponse represents an economic indicator data point.
type EconomicIndicatorResponse struct {
	// Name is the indicator name.
	Name string `json:"name"`
	// Date is the date of the data point.
	Date string `json:"date"`
	// Value is the indicator value.
	Value float64 `json:"value"`
}

// EconomicCalendarResponse represents an economic calendar event.
type EconomicCalendarResponse struct {
	// Date is the date and time of the event.
	Date string `json:"date"`
	// Country is the country code.
	Country string `json:"country"`
	// Event is the event description.
	Event string `json:"event"`
	// Currency is the currency affected.
	Currency string `json:"currency"`
	// Previous is the previous value.
	Previous float64 `json:"previous"`
	// Estimate is the consensus estimate.
	Estimate *float64 `json:"estimate"`
	// Actual is the actual reported value.
	Actual float64 `json:"actual"`
	// Change is the absolute change from previous.
	Change float64 `json:"change"`
	// Impact is the expected market impact level.
	Impact string `json:"impact"`
	// ChangePercentage is the percentage change.
	ChangePercentage float64 `json:"changePercentage"`
}

// MarketRiskPremiumResponse represents market risk premium by country.
type MarketRiskPremiumResponse struct {
	// Country is the country name.
	Country string `json:"country"`
	// Continent is the continent name.
	Continent string `json:"continent"`
	// CountryRiskPremium is the country-specific risk premium.
	CountryRiskPremium float64 `json:"countryRiskPremium"`
	// TotalEquityRiskPremium is the total equity risk premium.
	TotalEquityRiskPremium float64 `json:"totalEquityRiskPremium"`
}

// TreasuryRates retrieves US Treasury rates for all maturities.
// Returns yields from 1-month to 30-year treasuries.
//
// Endpoint: /treasury-rates
func (c *Client) TreasuryRates(params TreasuryRatesParams) ([]TreasuryRatesResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/treasury-rates"

	var result []TreasuryRatesResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// EconomicIndicators retrieves economic indicators data.
// Supports indicators like GDP, CPI, unemployment rate, retail sales, and more.
//
// Endpoint: /economic-indicators
func (c *Client) EconomicIndicators(params EconomicIndicatorsParams) ([]EconomicIndicatorResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/economic-indicators"

	var result []EconomicIndicatorResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// EconomicCalendar retrieves scheduled economic data releases.
// Returns upcoming and past economic events with actual vs estimate data.
//
// Endpoint: /economic-calendar
func (c *Client) EconomicCalendar(params EconomicCalendarParams) ([]EconomicCalendarResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/economic-calendar"

	var result []EconomicCalendarResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// MarketRiskPremium retrieves equity risk premiums by country.
// Shows country risk premiums and total equity risk premiums globally.
//
// Endpoint: /market-risk-premium
func (c *Client) MarketRiskPremium() ([]MarketRiskPremiumResponse, error) {
	pathName := "/market-risk-premium"

	var result []MarketRiskPremiumResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}