package go_fmp

import (
	"fmt"
	"time"
)

// CompanyProfileCIKParams represents the parameters for the Company Profile CIK API
type CompanyProfileCIKParams struct {
	CIK string `json:"cik"` // Required: Company CIK number (e.g., "0000320193")
}

// CompanyProfileCIKResponse represents the response from the Company Profile CIK API
type CompanyProfileCIKResponse struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Beta              float64 `json:"beta"`
	VolAvg            int64   `json:"volAvg"`
	MarketCap         int64   `json:"mktCap"`
	LastDiv           float64 `json:"lastDiv"`
	Range             string  `json:"range"`
	Changes           float64 `json:"changes"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	CIK               string  `json:"cik"`
	ISIN              string  `json:"isin"`
	CUSIP             string  `json:"cusip"`
	Exchange          Exchange  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
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
	DCFDiff           float64 `json:"dcfDiff"`
	DCF               float64 `json:"dcf"`
	Image             string  `json:"image"`
	IPODate           time.Time  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsAdr             bool    `json:"isAdr"`
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
	err := c.doRequest(c.BaseURL+"/profile-cik", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
