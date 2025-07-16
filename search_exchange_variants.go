package go_fmp

import (
	"encoding/json"
	"fmt"
)

// SearchExchangeVariantsParams represents the parameters for the Exchange Variants Search API
type SearchExchangeVariantsParams struct {
	Symbol string `json:"symbol"` // Required: Stock symbol (e.g., "AAPL")
}

// SearchExchangeVariantsResponse represents the response from the Exchange Variants Search API
type SearchExchangeVariantsResponse struct {
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Beta              float64 `json:"beta"`
	VolAvg            int64   `json:"volAvg"`
	MktCap            int64   `json:"mktCap"`
	LastDiv           float64 `json:"lastDiv"`
	Range             string  `json:"range"`
	Changes           float64 `json:"changes"`
	CompanyName       string  `json:"companyName"`
	Currency          string  `json:"currency"`
	CIK               string  `json:"cik"`
	ISIN              string  `json:"isin"`
	CUSIP             string  `json:"cusip"`
	Exchange          string  `json:"exchange"`
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
	IPODate           string  `json:"ipoDate"`
	DefaultImage      bool    `json:"defaultImage"`
	IsETF             bool    `json:"isEtf"`
	IsActivelyTrading bool    `json:"isActivelyTrading"`
	IsADR             bool    `json:"isAdr"`
	IsFund            bool    `json:"isFund"`
}

// SearchExchangeVariants searches across multiple public exchanges to find where a given stock symbol is listed
func (c *Client) SearchExchangeVariants(params SearchExchangeVariantsParams) ([]SearchExchangeVariantsResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	return doRequest[[]SearchExchangeVariantsResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
}
