package fmp_client

// CrowdfundingLatestParams represents parameters for latest crowdfunding campaigns.
type CrowdfundingLatestParams struct {
	// Page is the page number for pagination (optional).
	Page *int `json:"page,omitempty"`
	// Limit is the maximum number of results (optional).
	Limit *int `json:"limit,omitempty"`
}

// CrowdfundingSearchParams represents parameters for crowdfunding search.
type CrowdfundingSearchParams struct {
	// Name is the company or campaign name to search for (required).
	Name string `json:"name"`
}

// CrowdfundingByCIKParams represents parameters for crowdfunding by CIK.
type CrowdfundingByCIKParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
}

// EquityFundingSearchParams represents parameters for equity funding search.
type EquityFundingSearchParams struct {
	// Name is the company or funding name to search for (required).
	Name string `json:"name"`
}

// EquityFundingByCIKParams represents parameters for equity funding by CIK.
type EquityFundingByCIKParams struct {
	// CIK is the Central Index Key (required).
	CIK string `json:"cik"`
}

// CrowdfundingResponse represents a crowdfunding campaign.
type CrowdfundingResponse struct {
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// CompanyName is the company name.
	CompanyName string `json:"companyName"`
	// Date is the campaign date.
	Date string `json:"date"`
	// FilingDate is the filing date.
	FilingDate string `json:"filingDate"`
	// AcceptedDate is the accepted date.
	AcceptedDate string `json:"acceptedDate"`
	// FormType is the form type.
	FormType string `json:"formType"`
	// FormSignification is the form description.
	FormSignification string `json:"formSignification"`
	// NameOfIssuer is the issuer name.
	NameOfIssuer string `json:"nameOfIssuer"`
	// LegalStatusForm is the legal entity type.
	LegalStatusForm string `json:"legalStatusForm"`
	// JurisdictionOrganization is the jurisdiction.
	JurisdictionOrganization string `json:"jurisdictionOrganization"`
	// IssuerStreet is the street address.
	IssuerStreet string `json:"issuerStreet"`
	// IssuerCity is the city.
	IssuerCity string `json:"issuerCity"`
	// IssuerStateOrCountry is the state or country.
	IssuerStateOrCountry string `json:"issuerStateOrCountry"`
	// IssuerZipCode is the zip code.
	IssuerZipCode string `json:"issuerZipCode"`
	// IssuerWebsite is the website.
	IssuerWebsite string `json:"issuerWebsite"`
	// IntermediaryCompanyName is the platform name.
	IntermediaryCompanyName string `json:"intermediaryCompanyName"`
	// IntermediaryCommissionCIK is the platform CIK.
	IntermediaryCommissionCIK string `json:"intermediaryCommissionCik"`
	// IntermediaryCommissionFileNumber is the file number.
	IntermediaryCommissionFileNumber string `json:"intermediaryCommissionFileNumber"`
	// CompensationAmount is the compensation details.
	CompensationAmount string `json:"compensationAmount"`
	// FinancialInterest is the financial interest.
	FinancialInterest string `json:"financialInterest"`
	// SecurityOfferedType is the security type.
	SecurityOfferedType string `json:"securityOfferedType"`
	// SecurityOfferedOtherDescription is the security description.
	SecurityOfferedOtherDescription string `json:"securityOfferedOtherDescription"`
	// NumberOfSecurityOffered is the number of securities.
	NumberOfSecurityOffered int64 `json:"numberOfSecurityOffered"`
	// OfferingPrice is the price per security.
	OfferingPrice float64 `json:"offeringPrice"`
	// OfferingAmount is the offering amount.
	OfferingAmount int64 `json:"offeringAmount"`
	// OverSubscriptionAccepted indicates if over-subscription is accepted.
	OverSubscriptionAccepted string `json:"overSubscriptionAccepted"`
	// OverSubscriptionAllocationType is the allocation method.
	OverSubscriptionAllocationType string `json:"overSubscriptionAllocationType"`
	// MaximumOfferingAmount is the maximum amount.
	MaximumOfferingAmount int64 `json:"maximumOfferingAmount"`
	// OfferingDeadlineDate is the deadline date.
	OfferingDeadlineDate string `json:"offeringDeadlineDate"`
	// CurrentNumberOfEmployees is the employee count.
	CurrentNumberOfEmployees int `json:"currentNumberOfEmployees"`
	// TotalAssetMostRecentFiscalYear is the recent total assets.
	TotalAssetMostRecentFiscalYear int64 `json:"totalAssetMostRecentFiscalYear"`
	// TotalAssetPriorFiscalYear is the prior year total assets.
	TotalAssetPriorFiscalYear int64 `json:"totalAssetPriorFiscalYear"`
	// CashAndCashEquivalentMostRecentFiscalYear is the recent cash.
	CashAndCashEquivalentMostRecentFiscalYear int64 `json:"cashAndCashEquiValentMostRecentFiscalYear"`
	// CashAndCashEquivalentPriorFiscalYear is the prior year cash.
	CashAndCashEquivalentPriorFiscalYear int64 `json:"cashAndCashEquiValentPriorFiscalYear"`
	// AccountsReceivableMostRecentFiscalYear is the recent AR.
	AccountsReceivableMostRecentFiscalYear int64 `json:"accountsReceivableMostRecentFiscalYear"`
	// AccountsReceivablePriorFiscalYear is the prior year AR.
	AccountsReceivablePriorFiscalYear int64 `json:"accountsReceivablePriorFiscalYear"`
	// ShortTermDebtMostRecentFiscalYear is the recent short-term debt.
	ShortTermDebtMostRecentFiscalYear int64 `json:"shortTermDebtMostRecentFiscalYear"`
	// ShortTermDebtPriorFiscalYear is the prior year short-term debt.
	ShortTermDebtPriorFiscalYear int64 `json:"shortTermDebtPriorFiscalYear"`
	// LongTermDebtMostRecentFiscalYear is the recent long-term debt.
	LongTermDebtMostRecentFiscalYear int64 `json:"longTermDebtMostRecentFiscalYear"`
	// LongTermDebtPriorFiscalYear is the prior year long-term debt.
	LongTermDebtPriorFiscalYear int64 `json:"longTermDebtPriorFiscalYear"`
	// RevenueMostRecentFiscalYear is the recent revenue.
	RevenueMostRecentFiscalYear int64 `json:"revenueMostRecentFiscalYear"`
	// RevenuePriorFiscalYear is the prior year revenue.
	RevenuePriorFiscalYear int64 `json:"revenuePriorFiscalYear"`
	// CostGoodsSoldMostRecentFiscalYear is the recent COGS.
	CostGoodsSoldMostRecentFiscalYear int64 `json:"costGoodsSoldMostRecentFiscalYear"`
	// CostGoodsSoldPriorFiscalYear is the prior year COGS.
	CostGoodsSoldPriorFiscalYear int64 `json:"costGoodsSoldPriorFiscalYear"`
	// TaxesPaidMostRecentFiscalYear is the recent taxes paid.
	TaxesPaidMostRecentFiscalYear int64 `json:"taxesPaidMostRecentFiscalYear"`
	// TaxesPaidPriorFiscalYear is the prior year taxes paid.
	TaxesPaidPriorFiscalYear int64 `json:"taxesPaidPriorFiscalYear"`
	// NetIncomeMostRecentFiscalYear is the recent net income.
	NetIncomeMostRecentFiscalYear int64 `json:"netIncomeMostRecentFiscalYear"`
	// NetIncomePriorFiscalYear is the prior year net income.
	NetIncomePriorFiscalYear int64 `json:"netIncomePriorFiscalYear"`
}

// CrowdfundingSearchResponse represents crowdfunding search results.
type CrowdfundingSearchResponse struct {
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// Name is the company name.
	Name string `json:"name"`
	// Date is the campaign date.
	Date *string `json:"date"`
}

// EquityFundingResponse represents equity funding round.
type EquityFundingResponse struct {
	// CIK is the Central Index Key.
	CIK string `json:"cik"`
	// Name is the company name.
	Name string `json:"name"`
	// FundingRound is the funding round type.
	FundingRound string `json:"fundingRound"`
	// FundingDate is the funding date.
	FundingDate string `json:"fundingDate"`
	// RaisedAmount is the amount raised.
	RaisedAmount int64 `json:"raisedAmount"`
	// InvestorCount is the number of investors.
	InvestorCount int `json:"investorCount"`
	// ShareClass is the class of shares.
	ShareClass string `json:"shareClass"`
	// SharesIssued is the number of shares issued.
	SharesIssued int64 `json:"sharesIssued"`
	// SharePrice is the price per share.
	SharePrice float64 `json:"sharePrice"`
	// PreMoneyValuation is the pre-money valuation.
	PreMoneyValuation int64 `json:"preMoneyValuation"`
	// PostMoneyValuation is the post-money valuation.
	PostMoneyValuation int64 `json:"postMoneyValuation"`
}

// LatestCrowdfundingCampaigns retrieves recent crowdfunding campaigns.
// Returns active crowdfunding offerings with financial details.
//
// Endpoint: /crowdfunding-offerings-latest
func (c *Client) LatestCrowdfundingCampaigns(params CrowdfundingLatestParams) ([]CrowdfundingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/crowdfunding-offerings-latest"

	var result []CrowdfundingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchCrowdfundingCampaigns searches for crowdfunding campaigns.
// Returns campaigns matching the search criteria.
//
// Endpoint: /crowdfunding-offerings-search
func (c *Client) SearchCrowdfundingCampaigns(params CrowdfundingSearchParams) ([]CrowdfundingSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/crowdfunding-offerings-search"

	var result []CrowdfundingSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// CrowdfundingByCIK retrieves crowdfunding campaigns by CIK.
// Returns all campaigns for a specific company.
//
// Endpoint: /crowdfunding-offerings
func (c *Client) CrowdfundingByCIK(params CrowdfundingByCIKParams) ([]CrowdfundingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/crowdfunding-offerings"

	var result []CrowdfundingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// LatestEquityFundingRounds retrieves recent equity funding rounds.
// Returns private equity and VC funding information.
//
// Endpoint: /equity-funding-latest
func (c *Client) LatestEquityFundingRounds(params CrowdfundingLatestParams) ([]EquityFundingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/equity-funding-latest"

	var result []EquityFundingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// SearchEquityFundingRounds searches for equity funding rounds.
// Returns funding rounds matching the search criteria.
//
// Endpoint: /equity-funding-search
func (c *Client) SearchEquityFundingRounds(params EquityFundingSearchParams) ([]CrowdfundingSearchResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/equity-funding-search"

	var result []CrowdfundingSearchResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}

// EquityFundingByCIK retrieves equity funding rounds by CIK.
// Returns all funding rounds for a specific company.
//
// Endpoint: /equity-funding
func (c *Client) EquityFundingByCIK(params EquityFundingByCIKParams) ([]EquityFundingResponse, error) {
	queryParams := StructToMap(params)
	pathName := "/equity-funding"

	var result []EquityFundingResponse
	err := c.doRequest(pathName, queryParams, &result)
	return result, err
}