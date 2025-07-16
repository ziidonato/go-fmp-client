package go_fmp

import (
	"fmt"
)

// CompanyProfileCIKParams represents the parameters for the Company Profile by CIK API
type CompanyProfileCIKParams struct {
	CIK string `json:"cik"` // Required: Central Index Key (e.g., "320193")
}

// CompanyProfileCIKResponse represents the response from the Company Profile by CIK API
type CompanyProfileCIKResponse struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	MarketCap         int64   `json:"marketCap"`
	Beta              float64 `json:"beta"`
	LastDividend      float64 `json:"lastDividend"`
	Range             string  `json:"range"`
	Change            float64 `json:"change"`
	ChangePercentage  float64 `json:"changePercentage"`
	Volume            int64   `json:"volume"`
	AverageVolume     int64   `json:"averageVolume"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	CIK               string  `json:"cik"`
	ISIN              string  `json:"isin"`
	CUSIP             string  `json:"cusip"`
	ExchangeFullName  string  `json:"exchangeFullName"`
	Exchange          string  `json:"exchange"`
	Industry          string  `json:"industry"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	CEO               string  `json:"ceo"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	Image             string  `json:"image"`
	IPODate           string  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsETF             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsADR             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
}

// CompanyProfileCIK retrieves detailed company profile data by CIK (Central Index Key)
func (c *Client) CompanyProfileCIK(params CompanyProfileCIKParams) ([]CompanyProfileCIKResponse, error) {
	if params.CIK == "" {
		return nil, fmt.Errorf("cik parameter is required")
	}

	urlParams := map[string]string{
		"cik": params.CIK,
	}

	var result []CompanyProfileCIKResponse
	err := c.doRequest("https://financialmodelingprep.com/stable/profile-cik", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
