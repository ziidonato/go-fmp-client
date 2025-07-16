package go_fmp

import (
	"fmt"
)

// CrowdfundingOfferingsLatestParams represents the parameters for the Latest Crowdfunding Campaigns API
type CrowdfundingOfferingsLatestParams struct {
	Page  *int `json:"page"`  // Required: Page number (e.g., 0)
	Limit *int `json:"limit"` // Required: Number of results (e.g., 100)
}

// CrowdfundingOfferingsLatestResponse represents the response from the Latest Crowdfunding Campaigns API
type CrowdfundingOfferingsLatestResponse struct {
	CIK                                       string  `json:"cik"`
	CompanyName                               string  `json:"companyName"`
	Date                                      string  `json:"date"`
	FilingDate                                string  `json:"filingDate"`
	AcceptedDate                              string  `json:"acceptedDate"`
	FormType                                  string  `json:"formType"`
	FormSignification                         string  `json:"formSignification"`
	NameOfIssuer                              string  `json:"nameOfIssuer"`
	LegalStatusForm                           string  `json:"legalStatusForm"`
	JurisdictionOrganization                  string  `json:"jurisdictionOrganization"`
	IssuerStreet                              string  `json:"issuerStreet"`
	IssuerCity                                string  `json:"issuerCity"`
	IssuerStateOrCountry                      string  `json:"issuerStateOrCountry"`
	IssuerZipCode                             string  `json:"issuerZipCode"`
	IssuerWebsite                             string  `json:"issuerWebsite"`
	IntermediaryCompanyName                   string  `json:"intermediaryCompanyName"`
	IntermediaryCommissionCik                 string  `json:"intermediaryCommissionCik"`
	IntermediaryCommissionFileNumber          string  `json:"intermediaryCommissionFileNumber"`
	CompensationAmount                        string  `json:"compensationAmount"`
	FinancialInterest                         string  `json:"financialInterest"`
	SecurityOfferedType                       string  `json:"securityOfferedType"`
	SecurityOfferedOtherDescription           string  `json:"securityOfferedOtherDescription"`
	NumberOfSecurityOffered                   int64   `json:"numberOfSecurityOffered"`
	OfferingPrice                             float64 `json:"offeringPrice"`
	OfferingAmount                            int64   `json:"offeringAmount"`
	OverSubscriptionAccepted                  string  `json:"overSubscriptionAccepted"`
	OverSubscriptionAllocationType            string  `json:"overSubscriptionAllocationType"`
	MaximumOfferingAmount                     int64   `json:"maximumOfferingAmount"`
	OfferingDeadlineDate                      string  `json:"offeringDeadlineDate"`
	CurrentNumberOfEmployees                  int     `json:"currentNumberOfEmployees"`
	TotalAssetMostRecentFiscalYear            int64   `json:"totalAssetMostRecentFiscalYear"`
	TotalAssetPriorFiscalYear                 int64   `json:"totalAssetPriorFiscalYear"`
	CashAndCashEquiValentMostRecentFiscalYear int64   `json:"cashAndCashEquiValentMostRecentFiscalYear"`
	CashAndCashEquiValentPriorFiscalYear      int64   `json:"cashAndCashEquiValentPriorFiscalYear"`
	AccountsReceivableMostRecentFiscalYear    int64   `json:"accountsReceivableMostRecentFiscalYear"`
	AccountsReceivablePriorFiscalYear         int64   `json:"accountsReceivablePriorFiscalYear"`
	ShortTermDebtMostRecentFiscalYear         int64   `json:"shortTermDebtMostRecentFiscalYear"`
	ShortTermDebtPriorFiscalYear              int64   `json:"shortTermDebtPriorFiscalYear"`
	LongTermDebtMostRecentFiscalYear          int64   `json:"longTermDebtMostRecentFiscalYear"`
	LongTermDebtPriorFiscalYear               int64   `json:"longTermDebtPriorFiscalYear"`
	RevenueMostRecentFiscalYear               int64   `json:"revenueMostRecentFiscalYear"`
	RevenuePriorFiscalYear                    int64   `json:"revenuePriorFiscalYear"`
	CostGoodsSoldMostRecentFiscalYear         int64   `json:"costGoodsSoldMostRecentFiscalYear"`
	CostGoodsSoldPriorFiscalYear              int64   `json:"costGoodsSoldPriorFiscalYear"`
	TaxesPaidMostRecentFiscalYear             int64   `json:"taxesPaidMostRecentFiscalYear"`
	TaxesPaidPriorFiscalYear                  int64   `json:"taxesPaidPriorFiscalYear"`
	NetIncomeMostRecentFiscalYear             int64   `json:"netIncomeMostRecentFiscalYear"`
	NetIncomePriorFiscalYear                  int64   `json:"netIncomePriorFiscalYear"`
}

// CrowdfundingOfferingsLatest retrieves the most recent crowdfunding campaigns
func (c *Client) CrowdfundingOfferingsLatest(params CrowdfundingOfferingsLatestParams) ([]CrowdfundingOfferingsLatestResponse, error) {
	if params.Page == nil {
		return nil, fmt.Errorf("page parameter is required")
	}

	if params.Limit == nil {
		return nil, fmt.Errorf("limit parameter is required")
	}

	urlParams := map[string]string{
		"page":  fmt.Sprintf("%d", *params.Page),
		"limit": fmt.Sprintf("%d", *params.Limit),
	}

	var result []CrowdfundingOfferingsLatestResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/crowdfunding-offerings-latest", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
