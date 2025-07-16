package go_fmp

import (
	"encoding/json"
	"fmt"
)

// ETFAssetExposureParams represents the parameters for the ETF Asset Exposure API
type ETFAssetExposureParams struct {
	Symbol string `json:"symbol"` // Required: Asset symbol (e.g., "AAPL")
}

// ETFAssetExposureResponse represents the response from the ETF Asset Exposure API
type ETFAssetExposureResponse struct {
	Symbol           string  `json:"symbol"`
	Asset            string  `json:"asset"`
	SharesNumber     int64   `json:"sharesNumber"`
	WeightPercentage float64 `json:"weightPercentage"`
	MarketValue      float64 `json:"marketValue"`
}

// ETFAssetExposure retrieves information about which ETFs hold specific stocks
func (c *Client) ETFAssetExposure(params ETFAssetExposureParams) ([]ETFAssetExposureResponse, error) {
	if params.Symbol == "" {
		return nil, fmt.Errorf("symbol parameter is required")
	}

	urlParams := map[string]string{
		"symbol": params.Symbol,
	}

	resp, err := c.doRequest("https://financialmodelingprep.com/stable/etf/asset-exposure", urlParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result []ETFAssetExposureResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
