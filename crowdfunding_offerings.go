package go_fmp

import (
	"encoding/json"
	"fmt"
)

// CrowdfundingOfferingsParams represents the parameters for the Crowdfunding By CIK API
type CrowdfundingOfferingsParams struct {
	CIK string `json:"cik"` // Required: CIK number (e.g., "0001916078")
}

// CrowdfundingOfferingsResponse represents the response from the Crowdfunding By CIK API
type CrowdfundingOfferingsResponse struct {
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

// CrowdfundingOfferings retrieves detailed information on all crowdfunding campaigns launched by a specific company
func (c *Client) CrowdfundingOfferings(params CrowdfundingOfferingsParams) ([]CrowdfundingOfferingsResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	urlParams := map[string]string{
		"cik": params.CIK,
	}

	resp, err := c.get("https://financialmodelingprep.com/stable/crowdfunding-offerings", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []CrowdfundingOfferingsResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
