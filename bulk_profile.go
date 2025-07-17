package go_fmp

import (
	"fmt"
	"time"
)

// BulkProfileParams represents the parameters for the Bulk Profile API
type BulkProfileParams struct {
	Date string `json:"date"` // Required: date (e.g., "2024-10-22")
}

// BulkProfileResponse represents the response from the Bulk Profile API
type BulkProfileResponse struct {
	Symbol            string  `json:"symbol"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	Exchange          string  `json:"exchange"`
	MicCode           string  `json:"micCode"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Industry          string  `json:"industry"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	CEO               string  `json:"ceo"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	FullTimeEmployees int     `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	DCFDiff           float64 `json:"dcfDiff"`
	DCF               float64 `json:"dcf"`
	Image             string  `json:"image"`
	IpoDate           time.Time  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsEtf             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsAdr             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
	CIK               string  `json:"cik"`
	CUSIP             string  `json:"cusip"`
	ISIN              string  `json:"isin"`
	ISINBase          string  `json:"isinBase"`
	SEDOL             string  `json:"sedol"`
}

// BulkProfile retrieves comprehensive company profiles for multiple stocks
func (c *Client) BulkProfile(params BulkProfileParams) ([]BulkProfileResponse, error) {
	urlParams := map[string]string{
		"date": params.Date,
	}

	var result []BulkProfileResponse
	err := c.doRequest(c.BaseURL+"/bulk-profile", urlParams, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to get bulk profile: %w", err)
	}

	return result, nil
}
