package go_fmp

import (
	"encoding/json"
	"fmt"
)

// FundDisclosureParams represents the parameters for the Mutual Fund Disclosures API
type FundDisclosureParams struct {
	Symbol  string  `json:"symbol"`  // Required: Fund symbol (e.g., "VWO")
	Year    string  `json:"year"`    // Required: Year (e.g., "2023")
	Quarter string  `json:"quarter"` // Required: Quarter (e.g., "4")
	CIK     *string `json:"cik"`     // Optional: CIK number
}

// FundDisclosureResponse represents the response from the Mutual Fund Disclosures API
type FundDisclosureResponse struct {
	CIK                 string  `json:"cik"`
	Date                string  `json:"date"`
	AcceptedDate        string  `json:"acceptedDate"`
	Symbol              string  `json:"symbol"`
	Name                string  `json:"name"`
	LEI                 string  `json:"lei"`
	Title               string  `json:"title"`
	CUSIP               string  `json:"cusip"`
	ISIN                string  `json:"isin"`
	Balance             int64   `json:"balance"`
	Units               string  `json:"units"`
	CurCd               string  `json:"cur_cd"`
	ValUsd              float64 `json:"valUsd"`
	PctVal              float64 `json:"pctVal"`
	PayoffProfile       string  `json:"payoffProfile"`
	AssetCat            string  `json:"assetCat"`
	IssuerCat           string  `json:"issuerCat"`
	InvCountry          string  `json:"invCountry"`
	IsRestrictedSec     string  `json:"isRestrictedSec"`
	FairValLevel        string  `json:"fairValLevel"`
	IsCashCollateral    string  `json:"isCashCollateral"`
	IsNonCashCollateral string  `json:"isNonCashCollateral"`
	IsLoanByFund        string  `json:"isLoanByFund"`
}

// FundDisclosure retrieves comprehensive disclosure data for mutual funds
func (c *Client) FundDisclosure(params FundDisclosureParams) ([]FundDisclosureResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	if params.Year == "" {
		return nil, fmt.Errorf("year parameter is required")
	}

	if params.Quarter == "" {
		return nil, fmt.Errorf("quarter parameter is required")
	}

	urlParams := map[string]string{
		"symbol":  params.Symbol,
		"year":    params.Year,
		"quarter": params.Quarter,
	}

	if params.CIK != nil {
		urlParams["cik"] = *params.CIK
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/funds/disclosure", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []FundDisclosureResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
