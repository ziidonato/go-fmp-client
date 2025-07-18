package fmp_client

import "time"

// LatestInsiderTradingParams represents parameters for latest insider trading endpoint.
type LatestInsiderTradingParams struct {
	// Date filters trades by specific date (optional).
	Date *time.Time `json:"date,omitempty"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
}

// SearchInsiderTradingParams represents parameters for searching insider trades.
type SearchInsiderTradingParams struct {
	// Symbol is the stock ticker symbol (optional).
	Symbol *string `json:"symbol,omitempty"`
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results per page (optional).
	Limit *int `json:"limit,omitempty"`
	// ReportingCIK is the reporting person's CIK (optional).
	ReportingCIK *string `json:"reportingCik,omitempty"`
	// CompanyCIK is the company's CIK (optional).
	CompanyCIK *string `json:"companyCik,omitempty"`
	// TransactionType is the type of transaction (optional).
	TransactionType *string `json:"transactionType,omitempty"`
}

// SearchInsiderTradingByNameParams represents parameters for searching by name.
type SearchInsiderTradingByNameParams struct {
	// Name is the reporting person's name to search for (required).
	Name string `json:"name"`
}

// InsiderTradeStatisticsParams represents parameters for insider trade statistics.
type InsiderTradeStatisticsParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// AcquisitionOwnershipParams represents parameters for acquisition ownership.
type AcquisitionOwnershipParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// InsiderTradingResponse represents an insider trading transaction.
type InsiderTradingResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// FilingDate is when the form was filed.
	FilingDate string `json:"filingDate"`
	// TransactionDate is when the transaction occurred.
	TransactionDate string `json:"transactionDate"`
	// ReportingCIK is the CIK of the reporting person.
	ReportingCIK string `json:"reportingCik"`
	// CompanyCIK is the CIK of the company.
	CompanyCIK string `json:"companyCik"`
	// TransactionType describes the type of transaction.
	TransactionType string `json:"transactionType"`
	// SecuritiesOwned is the number of securities owned after transaction.
	SecuritiesOwned int64 `json:"securitiesOwned"`
	// ReportingName is the name of the reporting person.
	ReportingName string `json:"reportingName"`
	// TypeOfOwner describes the relationship to the company.
	TypeOfOwner string `json:"typeOfOwner"`
	// AcquisitionOrDisposition indicates if acquiring (A) or disposing (D).
	AcquisitionOrDisposition string `json:"acquisitionOrDisposition"`
	// DirectOrIndirect indicates direct (D) or indirect (I) ownership.
	DirectOrIndirect string `json:"directOrIndirect"`
	// FormType is the SEC form type.
	FormType string `json:"formType"`
	// SecuritiesTransacted is the number of securities in the transaction.
	SecuritiesTransacted int64 `json:"securitiesTransacted"`
	// Price is the transaction price per share.
	Price float64 `json:"price"`
	// SecurityName is the type of security.
	SecurityName string `json:"securityName"`
	// URL is the link to the SEC filing.
	URL string `json:"url"`
}

// InsiderTradingNameSearchResponse represents a reporting person search result.
type InsiderTradingNameSearchResponse struct {
	// ReportingCIK is the CIK of the reporting person.
	ReportingCIK string `json:"reportingCik"`
	// ReportingName is the name of the reporting person.
	ReportingName string `json:"reportingName"`
}

// InsiderTransactionTypeResponse represents a transaction type.
type InsiderTransactionTypeResponse struct {
	// TransactionType is the type of insider transaction.
	TransactionType string `json:"transactionType"`
}

// InsiderTradeStatisticsResponse represents insider trading statistics.
type InsiderTradeStatisticsResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// CIK is the company's Central Index Key.
	CIK string `json:"cik"`
	// Year is the year of the statistics.
	Year int `json:"year"`
	// Quarter is the quarter of the year.
	Quarter int `json:"quarter"`
	// AcquiredTransactions is the number of acquisition transactions.
	AcquiredTransactions int `json:"acquiredTransactions"`
	// DisposedTransactions is the number of disposal transactions.
	DisposedTransactions int `json:"disposedTransactions"`
	// AcquiredDisposedRatio is the ratio of acquisitions to disposals.
	AcquiredDisposedRatio float64 `json:"acquiredDisposedRatio"`
	// TotalAcquired is the total shares acquired.
	TotalAcquired int64 `json:"totalAcquired"`
	// TotalDisposed is the total shares disposed.
	TotalDisposed int64 `json:"totalDisposed"`
	// AverageAcquired is the average shares per acquisition.
	AverageAcquired float64 `json:"averageAcquired"`
	// AverageDisposed is the average shares per disposal.
	AverageDisposed float64 `json:"averageDisposed"`
	// TotalPurchases is the total purchase transactions.
	TotalPurchases int `json:"totalPurchases"`
	// TotalSales is the total sale transactions.
	TotalSales int `json:"totalSales"`
}

// AcquisitionOwnershipResponse represents beneficial ownership data.
type AcquisitionOwnershipResponse struct {
	// CIK is the company's Central Index Key.
	CIK string `json:"cik"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// FilingDate is when the form was filed.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is when the filing was accepted.
	AcceptedDate string `json:"acceptedDate"`
	// CUSIP is the CUSIP identifier.
	CUSIP string `json:"cusip"`
	// NameOfReportingPerson is the beneficial owner's name.
	NameOfReportingPerson string `json:"nameOfReportingPerson"`
	// CitizenshipOrPlaceOfOrganization indicates citizenship or organization location.
	CitizenshipOrPlaceOfOrganization string `json:"citizenshipOrPlaceOfOrganization"`
	// SoleVotingPower is shares with sole voting power.
	SoleVotingPower string `json:"soleVotingPower"`
	// SharedVotingPower is shares with shared voting power.
	SharedVotingPower string `json:"sharedVotingPower"`
	// SoleDispositivePower is shares with sole dispositive power.
	SoleDispositivePower string `json:"soleDispositivePower"`
	// SharedDispositivePower is shares with shared dispositive power.
	SharedDispositivePower string `json:"sharedDispositivePower"`
	// AmountBeneficiallyOwned is total shares beneficially owned.
	AmountBeneficiallyOwned string `json:"amountBeneficiallyOwned"`
	// PercentOfClass is percentage of class owned.
	PercentOfClass string `json:"percentOfClass"`
	// TypeOfReportingPerson describes the type of reporting person.
	TypeOfReportingPerson string `json:"typeOfReportingPerson"`
	// URL is the link to the SEC filing.
	URL string `json:"url"`
}

// LatestInsiderTrading retrieves the most recent insider trading transactions.
// Returns insider trades across all companies or filtered by date.
//
// Endpoint: /insider-trading/latest
func (c *Client) LatestInsiderTrading(params LatestInsiderTradingParams) ([]InsiderTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/insider-trading/latest"

	var result []InsiderTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchInsiderTrading searches for insider trading transactions by various criteria.
// Can filter by symbol, CIK, transaction type, and more.
//
// Endpoint: /insider-trading/search
func (c *Client) SearchInsiderTrading(params SearchInsiderTradingParams) ([]InsiderTradingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/insider-trading/search"

	var result []InsiderTradingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchInsiderTradingByName searches for reporting persons by name.
// Returns CIK and name matches for insider trading searches.
//
// Endpoint: /insider-trading/reporting-name
func (c *Client) SearchInsiderTradingByName(params SearchInsiderTradingByNameParams) ([]InsiderTradingNameSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/insider-trading/reporting-name"

	var result []InsiderTradingNameSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AllInsiderTransactionTypes retrieves all possible insider transaction types.
// Returns a list of transaction type codes used in insider trading reports.
//
// Endpoint: /insider-trading-transaction-type
func (c *Client) AllInsiderTransactionTypes() ([]InsiderTransactionTypeResponse, error) {
	pathName := "/insider-trading-transaction-type"

	var result []InsiderTransactionTypeResponse
	err := c.doRequest(pathName, map[string]string{}, &result)
	return result, err
}

// InsiderTradeStatistics retrieves aggregated insider trading statistics.
// Returns purchase/sale ratios and transaction counts by quarter.
//
// Endpoint: /insider-trading/statistics
func (c *Client) InsiderTradeStatistics(params InsiderTradeStatisticsParams) ([]InsiderTradeStatisticsResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/insider-trading/statistics"

	var result []InsiderTradeStatisticsResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// AcquisitionOwnership retrieves beneficial ownership reports (Schedule 13D/G).
// Shows major shareholders and ownership changes.
//
// Endpoint: /acquisition-of-beneficial-ownership
func (c *Client) AcquisitionOwnership(params AcquisitionOwnershipParams) ([]AcquisitionOwnershipResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/acquisition-of-beneficial-ownership"

	var result []AcquisitionOwnershipResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}