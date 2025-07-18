package fmp_client

// CIKListParams represents parameters for CIK list endpoint.
type CIKListParams struct {
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// SymbolChangesParams represents parameters for symbol changes endpoint.
type SymbolChangesParams struct {
	// Invalid filters for invalid symbols when false (optional).
	Invalid *string `json:"invalid,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// StockListResponse represents a company in stock list.
type StockListResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
}

// FinancialStatementSymbolResponse represents a company with financial statements.
type FinancialStatementSymbolResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// TradingCurrency is the currency used for trading.
	TradingCurrency string `json:"tradingCurrency"`
	// ReportingCurrency is the currency used in financial reports.
	ReportingCurrency string `json:"reportingCurrency"`
}

// CIKListResponse represents a CIK entry.
type CIKListResponse struct {
	// CIK is the Central Index Key with leading zeros.
	CIK string `json:"cik"`
	// CompanyName is the registered company name.
	CompanyName string `json:"companyName"`
}

// SymbolChangeResponse represents a symbol change event.
type SymbolChangeResponse struct {
	// Date is when the symbol change occurred.
	Date string `json:"date"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// OldSymbol is the previous ticker symbol.
	OldSymbol string `json:"oldSymbol"`
	// NewSymbol is the new ticker symbol.
	NewSymbol string `json:"newSymbol"`
}

// ETFListResponse represents an ETF entry.
type ETFListResponse struct {
	// Symbol is the ETF ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the ETF name.
	Name string `json:"name"`
}

// ActivelyTradingResponse represents an actively trading security.
type ActivelyTradingResponse struct {
	// Symbol is the ticker symbol.
	Symbol string `json:"symbol"`
	// Name is the security name.
	Name string `json:"name"`
}

// EarningsTranscriptListResponse represents earnings transcript availability.
type EarningsTranscriptListResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CompanyName is the full company name.
	CompanyName string `json:"companyName"`
	// NoOfTranscripts is the number of available transcripts.
	NoOfTranscripts string `json:"noOfTranscripts"`
}

// ExchangeResponse represents exchange information.
type ExchangeResponse struct {
	// Exchange is the exchange code.
	Exchange string `json:"exchange"`
	// Name is the full exchange name.
	Name string `json:"name"`
	// CountryName is the country where the exchange operates.
	CountryName string `json:"countryName"`
	// CountryCode is the ISO country code.
	CountryCode string `json:"countryCode"`
	// SymbolSuffix is the suffix used for symbols on this exchange.
	SymbolSuffix string `json:"symbolSuffix"`
	// Delay describes the data delay for this exchange.
	Delay string `json:"delay"`
}

// SectorResponse represents a market sector.
type SectorResponse struct {
	// Sector is the sector name.
	Sector string `json:"sector"`
}

// IndustryResponse represents an industry category.
type IndustryResponse struct {
	// Industry is the industry name.
	Industry string `json:"industry"`
}

// CountryResponse represents a country code.
type CountryResponse struct {
	// Country is the country code.
	Country string `json:"country"`
}

// StockList retrieves a comprehensive list of all available stock symbols.
// Returns symbols and company names from various global exchanges.
//
// Endpoint: /stock-list
func (c *Client) StockList() ([]StockListResponse, error) {
	pathName := "/stock-list"

	var result []StockListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// FinancialStatementSymbolList retrieves companies with available financial statements.
// Returns symbols with trading and reporting currency information.
//
// Endpoint: /financial-statement-symbol-list
func (c *Client) FinancialStatementSymbolList() ([]FinancialStatementSymbolResponse, error) {
	pathName := "/financial-statement-symbol-list"

	var result []FinancialStatementSymbolResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// CIKList retrieves a comprehensive list of CIK numbers.
// Returns Central Index Keys for SEC-registered entities.
//
// Endpoint: /cik-list
func (c *Client) CIKList(params CIKListParams) ([]CIKListResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/cik-list"

	var result []CIKListResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SymbolChanges retrieves recent stock symbol changes.
// Returns changes due to mergers, acquisitions, and rebranding.
//
// Endpoint: /symbol-change
func (c *Client) SymbolChanges(params SymbolChangesParams) ([]SymbolChangeResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/symbol-change"

	var result []SymbolChangeResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ETFList retrieves all available ETF symbols.
// Returns Exchange Traded Funds with their symbols and names.
//
// Endpoint: /etf-list
func (c *Client) ETFList() ([]ETFListResponse, error) {
	pathName := "/etf-list"

	var result []ETFListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// ActivelyTradingList retrieves all actively trading securities.
// Returns symbols currently being traded on public exchanges.
//
// Endpoint: /actively-trading-list
func (c *Client) ActivelyTradingList() ([]ActivelyTradingResponse, error) {
	pathName := "/actively-trading-list"

	var result []ActivelyTradingResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// EarningsTranscriptList retrieves companies with available earnings transcripts.
// Returns symbols with the count of available transcripts.
//
// Endpoint: /earnings-transcript-list
func (c *Client) EarningsTranscriptList() ([]EarningsTranscriptListResponse, error) {
	pathName := "/earnings-transcript-list"

	var result []EarningsTranscriptListResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// AvailableExchanges retrieves all supported stock exchanges.
// Returns exchange codes with country and delay information.
//
// Endpoint: /available-exchanges
func (c *Client) AvailableExchanges() ([]ExchangeResponse, error) {
	pathName := "/available-exchanges"

	var result []ExchangeResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// AvailableSectors retrieves all available market sectors.
// Returns a list of sector categories for classification.
//
// Endpoint: /available-sectors
func (c *Client) AvailableSectors() ([]SectorResponse, error) {
	pathName := "/available-sectors"

	var result []SectorResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// AvailableIndustries retrieves all available industries.
// Returns a list of industry categories for detailed classification.
//
// Endpoint: /available-industries
func (c *Client) AvailableIndustries() ([]IndustryResponse, error) {
	pathName := "/available-industries"

	var result []IndustryResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// AvailableCountries retrieves all available country codes.
// Returns countries where securities are available for trading.
//
// Endpoint: /available-countries
func (c *Client) AvailableCountries() ([]CountryResponse, error) {
	pathName := "/available-countries"

	var result []CountryResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}