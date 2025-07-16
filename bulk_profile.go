package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ProfileBulkParams represents the parameters for the Profile Bulk API
type ProfileBulkParams struct {
	Part string `json:"part"` // Required: part number (e.g., "0")
}

// ProfileBulkResponse represents the response from the Profile Bulk API
type ProfileBulkResponse struct {
	Symbol            string `json:"symbol"`
	Price             string `json:"price"`
	MarketCap         string `json:"marketCap"`
	Beta              string `json:"beta"`
	LastDividend      string `json:"lastDividend"`
	Range             string `json:"range"`
	Change            string `json:"change"`
	ChangePercentage  string `json:"changePercentage"`
	Volume            string `json:"volume"`
	AverageVolume     string `json:"averageVolume"`
	CompanyName       string `json:"companyName"`
	Currency          string `json:"currency"`
	CIK               string `json:"cik"`
	ISIN              string `json:"isin"`
	CUSIP             string `json:"cusip"`
	ExchangeFullName  string `json:"exchangeFullName"`
	Exchange          string `json:"exchange"`
	Industry          string `json:"industry"`
	Website           string `json:"website"`
	Description       string `json:"description"`
	CEO               string `json:"ceo"`
	Sector            string `json:"sector"`
	Country           string `json:"country"`
	FullTimeEmployees string `json:"fullTimeEmployees"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	City              string `json:"city"`
	State             string `json:"state"`
	Zip               string `json:"zip"`
	Image             string `json:"image"`
	IPODate           string `json:"ipoDate"`
	DefaultImage      string `json:"defaultImage"`
	IsETF             string `json:"isEtf"`
	IsActivelyTrading string `json:"isActivelyTrading"`
	IsADR             string `json:"isAdr"`
	IsFund            string `json:"isFund"`
}

// GetProfileBulk retrieves comprehensive company profile data in bulk
func (c *Client) GetProfileBulk(params ProfileBulkParams) ([]ProfileBulkResponse, error) {
	// Validate required parameters
	if params.Part == "" {
		return nil, fmt.Errorf("part parameter is required")
	}

	urlParams := map[string]string{
		"part": params.Part,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/profile-bulk", urlParams)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result []ProfileBulkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
}
