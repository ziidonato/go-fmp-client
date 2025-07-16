package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CompanyScreenerParams represents the parameters for the Stock Screener API
type CompanyScreenerParams struct {
	MarketCapMoreThan      *int64   `json:"marketCapMoreThan"`      // Optional: Minimum market cap
	MarketCapLowerThan     *int64   `json:"marketCapLowerThan"`     // Optional: Maximum market cap
	Sector                 *string  `json:"sector"`                 // Optional: Sector filter (e.g., "Technology")
	Industry               *string  `json:"industry"`               // Optional: Industry filter (e.g., "Consumer Electronics")
	BetaMoreThan           *float64 `json:"betaMoreThan"`           // Optional: Minimum beta
	BetaLowerThan          *float64 `json:"betaLowerThan"`          // Optional: Maximum beta
	PriceMoreThan          *float64 `json:"priceMoreThan"`          // Optional: Minimum price
	PriceLowerThan         *float64 `json:"priceLowerThan"`         // Optional: Maximum price
	DividendMoreThan       *float64 `json:"dividendMoreThan"`       // Optional: Minimum dividend
	DividendLowerThan      *float64 `json:"dividendLowerThan"`      // Optional: Maximum dividend
	VolumeMoreThan         *int64   `json:"volumeMoreThan"`         // Optional: Minimum volume
	VolumeLowerThan        *int64   `json:"volumeLowerThan"`        // Optional: Maximum volume
	Exchange               *string  `json:"exchange"`               // Optional: Exchange filter (e.g., "NASDAQ")
	Country                *string  `json:"country"`                // Optional: Country filter (e.g., "US")
	IsETF                  *bool    `json:"isEtf"`                  // Optional: ETF filter
	IsFund                 *bool    `json:"isFund"`                 // Optional: Fund filter
	IsActivelyTrading      *bool    `json:"isActivelyTrading"`      // Optional: Actively trading filter
	Limit                  *int     `json:"limit"`                  // Optional: Number of results (default: 1000)
	IncludeAllShareClasses *bool    `json:"includeAllShareClasses"` // Optional: Include all share classes

// CompanyScreenerResponse represents the response from the Stock Screener API
type CompanyScreenerResponse struct {
	Symbol             string  `json:"symbol"`
	CompanyName        string  `json:"companyName"`
	MarketCap          int64   `json:"marketCap"`
	Sector             string  `json:"sector"`
	Industry           string  `json:"industry"`
	Beta               float64 `json:"beta"`
	Price              float64 `json:"price"`
	LastAnnualDividend float64 `json:"lastAnnualDividend"`
	Volume             int64   `json:"volume"`
	Exchange           string  `json:"exchange"`
	ExchangeShortName  string  `json:"exchangeShortName"`
	Country            string  `json:"country"`
	IsETF              bool    `json:"isEtf"`
	IsFund             bool    `json:"isFund"`
	IsActivelyTrading  bool    `json:"isActivelyTrading"`

// CompanyScreener discovers stocks that align with investment strategy using various filters
func (c *Client) CompanyScreener(params CompanyScreenerParams) ([]CompanyScreenerResponse, error) {
	urlParams := make(map[string]string)

	if params.MarketCapMoreThan != nil {
		urlParams["marketCapMoreThan"] = strconv.FormatInt(*params.MarketCapMoreThan, 10)
	}

	if params.MarketCapLowerThan != nil {
		urlParams["marketCapLowerThan"] = strconv.FormatInt(*params.MarketCapLowerThan, 10)
	}

	if params.Sector != nil {
		urlParams["sector"] = *params.Sector
	}

	if params.Industry != nil {
		urlParams["industry"] = *params.Industry
	}

	if params.BetaMoreThan != nil {
		urlParams["betaMoreThan"] = strconv.FormatFloat(*params.BetaMoreThan, 'f', -1, 64)
	}

	if params.BetaLowerThan != nil {
		urlParams["betaLowerThan"] = strconv.FormatFloat(*params.BetaLowerThan, 'f', -1, 64)
	}

	if params.PriceMoreThan != nil {
		urlParams["priceMoreThan"] = strconv.FormatFloat(*params.PriceMoreThan, 'f', -1, 64)
	}

	if params.PriceLowerThan != nil {
		urlParams["priceLowerThan"] = strconv.FormatFloat(*params.PriceLowerThan, 'f', -1, 64)
	}

	if params.DividendMoreThan != nil {
		urlParams["dividendMoreThan"] = strconv.FormatFloat(*params.DividendMoreThan, 'f', -1, 64)
	}

	if params.DividendLowerThan != nil {
		urlParams["dividendLowerThan"] = strconv.FormatFloat(*params.DividendLowerThan, 'f', -1, 64)
	}

	if params.VolumeMoreThan != nil {
		urlParams["volumeMoreThan"] = strconv.FormatInt(*params.VolumeMoreThan, 10)
	}

	if params.VolumeLowerThan != nil {
		urlParams["volumeLowerThan"] = strconv.FormatInt(*params.VolumeLowerThan, 10)
	}

	if params.Exchange != nil {
		urlParams["exchange"] = *params.Exchange
	}

	if params.Country != nil {
		urlParams["country"] = *params.Country
	}

	if params.IsETF != nil {
		urlParams["isEtf"] = strconv.FormatBool(*params.IsETF)
	}

	if params.IsFund != nil {
		urlParams["isFund"] = strconv.FormatBool(*params.IsFund)
	}

	if params.IsActivelyTrading != nil {
		urlParams["isActivelyTrading"] = strconv.FormatBool(*params.IsActivelyTrading)
	}

	if params.Limit != nil {
		urlParams["limit"] = strconv.Itoa(*params.Limit)
	}

	if params.IncludeAllShareClasses != nil {
		urlParams["includeAllShareClasses"] = strconv.FormatBool(*params.IncludeAllShareClasses)
	}

	return doRequest[[]CompanyScreenerResponse](c, "https://financialmodelingprep.com/stable/analyst-estimates", urlParams)
