package fmp_client

import "time"

// LatestFilingsParams represents parameters for latest filings.
type LatestFilingsParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// FilingsExtractParams represents parameters for filings extract.
type FilingsExtractParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
	// Year is the filing year (required).
	Year string `json:"year"`
	// Quarter is the filing quarter (required).
	Quarter string `json:"quarter"`
}

// FilingsDatesParams represents parameters for filings dates.
type FilingsDatesParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
}

// InstitutionalOwnershipParams represents parameters for institutional ownership.
type InstitutionalOwnershipParams struct {
	// Symbol is the stock ticker symbol (required).
	Symbol string `json:"symbol"`
}

// InstitutionalHolderParams represents parameters for institutional holder.
type InstitutionalHolderParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
	// Date is the filing date (optional).
	Date *time.Time `json:"date,omitempty"`
}

// InstitutionalOwnershipResponse represents institutional ownership filing.
type InstitutionalOwnershipResponse struct {
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// Name is the institution name.
	Name string `json:"name"`
	// Date is the reporting date.
	Date string `json:"date"`
	// FilingDate is the filing date.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is the accepted date.
	AcceptedDate string `json:"acceptedDate"`
	// FormType is the form type.
	FormType string `json:"formType"`
	// Link is the filing link.
	Link string `json:"link"`
	// FinalLink is the final document link.
	FinalLink string `json:"finalLink"`
}

// FilingExtractResponse represents extracted filing data.
type FilingExtractResponse struct {
	// Date is the reporting date.
	Date string `json:"date"`
	// FilingDate is the filing date.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is the accepted date.
	AcceptedDate string `json:"acceptedDate"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// SecurityCUSIP is the CUSIP identifier.
	SecurityCUSIP string `json:"securityCusip"`
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// NameOfIssuer is the issuer name.
	NameOfIssuer string `json:"nameOfIssuer"`
	// Shares is the number of shares.
	Shares int64 `json:"shares"`
	// TitleOfClass is the security class.
	TitleOfClass string `json:"titleOfClass"`
	// SharesType is the type of shares.
	SharesType string `json:"sharesType"`
	// PutCallShare indicates put/call option.
	PutCallShare string `json:"putCallShare"`
	// Value is the position value.
	Value int64 `json:"value"`
	// Link is the filing link.
	Link string `json:"link"`
	// FinalLink is the final document link.
	FinalLink string `json:"finalLink"`
}

// FilingDateResponse represents filing date information.
type FilingDateResponse struct {
	// FilingDate is the filing date.
	FilingDate string `json:"filingDate"`
	// Date is the reporting date.
	Date string `json:"date"`
}

// InstitutionalOwnershipBySymbolResponse represents ownership by symbol.
type InstitutionalOwnershipBySymbolResponse struct {
	// InvestorName is the institutional investor name.
	InvestorName string `json:"investorName"`
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// DateReported is the report date.
	DateReported string `json:"dateReported"`
	// Shares is the number of shares held.
	Shares int64 `json:"shares"`
	// PutCallShare indicates put/call option.
	PutCallShare string `json:"putCallShare"`
	// Change is the change in shares.
	Change int64 `json:"change"`
	// PercentageChange is the percentage change.
	PercentageChange float64 `json:"percentageChange"`
}

// InstitutionalHolderResponse represents institutional holder portfolio.
type InstitutionalHolderResponse struct {
	// Symbol is the stock ticker symbol.
	Symbol string `json:"symbol"`
	// SecurityCUSIP is the CUSIP identifier.
	SecurityCUSIP string `json:"securityCusip"`
	// NameOfIssuer is the issuer name.
	NameOfIssuer string `json:"nameOfIssuer"`
	// Shares is the number of shares.
	Shares int64 `json:"shares"`
	// TitleOfClass is the security class.
	TitleOfClass string `json:"titleOfClass"`
	// Value is the position value.
	Value int64 `json:"value"`
	// SharesType is the type of shares.
	SharesType string `json:"sharesType"`
	// PutCallShare indicates put/call option.
	PutCallShare string `json:"putCallShare"`
	// InvestmentDiscretion is the investment discretion.
	InvestmentDiscretion string `json:"investmentDiscretion"`
	// OtherManager is other manager info.
	OtherManager string `json:"otherManager"`
	// VotingAuthoritySole is sole voting authority.
	VotingAuthoritySole int64 `json:"votingAuthoritySole"`
	// VotingAuthorityShared is shared voting authority.
	VotingAuthorityShared int64 `json:"votingAuthorityShared"`
	// VotingAuthorityNone is no voting authority.
	VotingAuthorityNone int64 `json:"votingAuthorityNone"`
}

// CIKListResponse represents CIK list entry.
type Form13FCIKListResponse struct {
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// Name is the institution name.
	Name string `json:"name"`
}

// LatestInstitutionalOwnershipFilings retrieves recent 13F filings.
// Returns latest institutional ownership disclosures.
//
// Endpoint: /institutional-ownership/latest
func (c *Client) LatestInstitutionalOwnershipFilings(params LatestFilingsParams) ([]InstitutionalOwnershipResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/institutional-ownership/latest"

	var result []InstitutionalOwnershipResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// ExtractFilingsData extracts detailed 13F filing data.
// Returns holdings details for a specific filing.
//
// Endpoint: /institutional-ownership/extract
func (c *Client) ExtractFilingsData(params FilingsExtractParams) ([]FilingExtractResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/institutional-ownership/extract"

	var result []FilingExtractResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// Form13FFilingsDates retrieves filing dates for an institution.
// Returns list of 13F filing dates for tracking.
//
// Endpoint: /institutional-ownership/dates
func (c *Client) Form13FFilingsDates(params FilingsDatesParams) ([]FilingDateResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/institutional-ownership/dates"

	var result []FilingDateResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// InstitutionalOwnershipBySymbol retrieves institutional holders of a stock.
// Returns list of institutions holding the given symbol.
//
// Endpoint: /institutional-ownership/symbol
func (c *Client) InstitutionalOwnershipBySymbol(params InstitutionalOwnershipParams) ([]InstitutionalOwnershipBySymbolResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/institutional-ownership/symbol"

	var result []InstitutionalOwnershipBySymbolResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// InstitutionalHolderPortfolio retrieves an institution's portfolio.
// Returns all holdings for a specific institutional investor.
//
// Endpoint: /institutional-ownership/holder
func (c *Client) InstitutionalHolderPortfolio(params InstitutionalHolderParams) ([]InstitutionalHolderResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/institutional-ownership/holder"

	var result []InstitutionalHolderResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// Form13FCIKList retrieves list of institutional CIKs.
// Returns institutions that file Form 13F.
//
// Endpoint: /institutional-ownership/cik-list
func (c *Client) Form13FCIKList() ([]Form13FCIKListResponse, error) {
	pathName := "/institutional-ownership/cik-list"

	var result []Form13FCIKListResponse
	err := c.doRequest(pathName, nil, &result)
	return result, err
}
