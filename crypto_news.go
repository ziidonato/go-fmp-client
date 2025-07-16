package go_fmp

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// CryptoNewsResponse represents the response from the crypto news API
type CryptoNewsResponse struct {
	Symbol        string `json:"symbol"`
	PublishedDate string `json:"publishedDate"`
	Publisher     string `json:"publisher"`
	Title         string `json:"title"`
	Image         string `json:"image"`
	Site          string `json:"site"`
	Text          string `json:"text"`
	URL           string `json:"url"`
}

// GetCryptoNews retrieves the latest cryptocurrency news
func (c *Client) GetCryptoNews(page, limit int, from, to string) ([]CryptoNewsResponse, error) {
	if page < 0 {
		return nil, fmt.Errorf("page must be non-negative")
	}
	if limit <= 0 || limit > 250 {
		return nil, fmt.Errorf("limit must be between 1 and 250")
	}
	if page > 100 {
		return nil, fmt.Errorf("page cannot exceed 100")
	}

	params := map[string]string{
		"page":  strconv.Itoa(page),
		"limit": strconv.Itoa(limit),
	}

	if from != " {
		params["from"] = from
	}
	if to != " {
		params["to"] = to
	}

	url := "https://financialmodelingprep.com/stable/news/crypto-latest"

	return doRequest[[]CryptoNewsResponse](c, url, params)
}
