package fmp_client

import "time"

// MarketPerformanceSnapshotParams represents parameters for snapshot endpoints.
type MarketPerformanceSnapshotParams struct {
	// Date is the date for the snapshot (required).
	Date string `json:"date"`
	// Exchange filters by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
	// Sector filters by specific sector (optional).
	Sector *string `json:"sector,omitempty"`
}

// IndustryPerformanceSnapshotParams represents parameters for industry snapshots.
type IndustryPerformanceSnapshotParams struct {
	// Date is the date for the snapshot (required).
	Date string `json:"date"`
	// Exchange filters by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
	// Industry filters by specific industry (optional).
	Industry *string `json:"industry,omitempty"`
}

// HistoricalSectorParams represents parameters for historical sector data.
type HistoricalSectorParams struct {
	// Sector is the sector name (required).
	Sector string `json:"sector"`
	// From is the start date for filtering (optional).
	From *string `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *string `json:"to,omitempty"`
	// Exchange filters by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
}

// HistoricalIndustryParams represents parameters for historical industry data.
type HistoricalIndustryParams struct {
	// Industry is the industry name (required).
	Industry string `json:"industry"`
	// From is the start date for filtering (optional).
	From *string `json:"from,omitempty"`
	// To is the end date for filtering (optional).
	To *string `json:"to,omitempty"`
	// Exchange filters by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
}

// SectorPESnapshotParams represents parameters for sector PE snapshots.
type SectorPESnapshotParams struct {
	// Date is the date for the snapshot (required).
	Date string `json:"date"`
	// Exchange filters by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
	// Sector filters by specific sector (optional).
	Sector *string `json:"sector,omitempty"`
}

// IndustryPESnapshotParams represents parameters for industry PE snapshots.
type IndustryPESnapshotParams struct {
	// Date is the date for the snapshot (required).
	Date string `json:"date"`
	// Exchange filters by specific exchange (optional).
	Exchange *string `json:"exchange,omitempty"`
	// Industry filters by specific industry (optional).
	Industry *string `json:"industry,omitempty"`
}

// SectorPerformanceResponse represents sector performance data.
type SectorPerformanceResponse struct {
	// Date is the date of the data.
	Date string `json:"date"`
	// Sector is the sector name.
	Sector string `json:"sector"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
	// AverageChange is the average price change percentage.
	AverageChange float64 `json:"averageChange"`
}

// IndustryPerformanceResponse represents industry performance data.
type IndustryPerformanceResponse struct {
	// Date is the date of the data.
	Date string `json:"date"`
	// Industry is the industry name.
	Industry string `json:"industry"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
	// AverageChange is the average price change percentage.
	AverageChange float64 `json:"averageChange"`
}

// SectorPEResponse represents sector PE ratio data.
type SectorPEResponse struct {
	// Date is the date of the data.
	Date string `json:"date"`
	// Sector is the sector name.
	Sector string `json:"sector"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
	// PE is the price-to-earnings ratio.
	PE float64 `json:"pe"`
}

// IndustryPEResponse represents industry PE ratio data.
type IndustryPEResponse struct {
	// Date is the date of the data.
	Date string `json:"date"`
	// Industry is the industry name.
	Industry string `json:"industry"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
	// PE is the price-to-earnings ratio.
	PE float64 `json:"pe"`
}

// StockMoverResponse represents a stock mover (gainer/loser).
type StockMoverResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// Price is the current stock price.
	Price float64 `json:"price"`
	// Name is the company name.
	Name string `json:"name"`
	// Change is the price change amount.
	Change float64 `json:"change"`
	// ChangesPercentage is the percentage change.
	ChangesPercentage float64 `json:"changesPercentage"`
	// Exchange is the exchange name.
	Exchange string `json:"exchange"`
}

// SectorPerformanceSnapshot retrieves sector performance for a specific date.
// Returns average changes across sectors.
//
// Endpoint: /sector-performance-snapshot
func (c *Client) SectorPerformanceSnapshot(params MarketPerformanceSnapshotParams) ([]SectorPerformanceResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sector-performance-snapshot"

	var result []SectorPerformanceResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndustryPerformanceSnapshot retrieves industry performance for a specific date.
// Returns average changes across industries.
//
// Endpoint: /industry-performance-snapshot
func (c *Client) IndustryPerformanceSnapshot(params IndustryPerformanceSnapshotParams) ([]IndustryPerformanceResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/industry-performance-snapshot"

	var result []IndustryPerformanceResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HistoricalSectorPerformance retrieves historical sector performance data.
// Returns sector performance over a date range.
//
// Endpoint: /historical-sector-performance
func (c *Client) HistoricalSectorPerformance(params HistoricalSectorParams) ([]SectorPerformanceResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-sector-performance"

	var result []SectorPerformanceResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HistoricalIndustryPerformance retrieves historical industry performance data.
// Returns industry performance over a date range.
//
// Endpoint: /historical-industry-performance
func (c *Client) HistoricalIndustryPerformance(params HistoricalIndustryParams) ([]IndustryPerformanceResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-industry-performance"

	var result []IndustryPerformanceResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SectorPESnapshot retrieves sector PE ratios for a specific date.
// Returns price-to-earnings ratios across sectors.
//
// Endpoint: /sector-pe-snapshot
func (c *Client) SectorPESnapshot(params SectorPESnapshotParams) ([]SectorPEResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/sector-pe-snapshot"

	var result []SectorPEResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// IndustryPESnapshot retrieves industry PE ratios for a specific date.
// Returns price-to-earnings ratios across industries.
//
// Endpoint: /industry-pe-snapshot
func (c *Client) IndustryPESnapshot(params IndustryPESnapshotParams) ([]IndustryPEResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/industry-pe-snapshot"

	var result []IndustryPEResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HistoricalSectorPE retrieves historical sector PE ratio data.
// Returns sector PE ratios over a date range.
//
// Endpoint: /historical-sector-pe
func (c *Client) HistoricalSectorPE(params HistoricalSectorParams) ([]SectorPEResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-sector-pe"

	var result []SectorPEResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// HistoricalIndustryPE retrieves historical industry PE ratio data.
// Returns industry PE ratios over a date range.
//
// Endpoint: /historical-industry-pe
func (c *Client) HistoricalIndustryPE(params HistoricalIndustryParams) ([]IndustryPEResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/historical-industry-pe"

	var result []IndustryPEResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// BiggestGainers retrieves stocks with the largest price increases.
// Returns top performing stocks by percentage gain.
//
// Endpoint: /biggest-gainers
func (c *Client) BiggestGainers() ([]StockMoverResponse, error) {
	pathName := "/biggest-gainers"

	var result []StockMoverResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// BiggestLosers retrieves stocks with the largest price decreases.
// Returns worst performing stocks by percentage loss.
//
// Endpoint: /biggest-losers
func (c *Client) BiggestLosers() ([]StockMoverResponse, error) {
	pathName := "/biggest-losers"

	var result []StockMoverResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// MostActives retrieves the most actively traded stocks.
// Returns stocks with the highest trading volumes.
//
// Endpoint: /most-actives
func (c *Client) MostActives() ([]StockMoverResponse, error) {
	pathName := "/most-actives"

	var result []StockMoverResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}