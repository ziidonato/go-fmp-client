package go_fmp

import (
	"fmt"
)

// CompanyProfileParams represents the parameters for the Company Profile Data API
type CompanyProfileParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// CompanyProfileResponse represents the response from the Company Profile Data API
type CompanyProfileResponse struct {
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

// CompanyProfile retrieves detailed company profile data for a specific stock symbol
func (c *Client) CompanyProfile(params CompanyProfileParams) ([]CompanyProfileResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	var result []CompanyProfileResponse
	err := c.doRequest(c.BaseURL+"/profile", urlParams, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
