package go_fmp

import (
	"fmt"
	"strconv"
	"time"
)

// SearchCryptoNewsResponse represents the response from the search crypto news API
type SearchCryptoNewsResponse struct {
	Symbol        string    `json:"symbol"`
	PublishedDate time.Time `json:"publishedDate"`
	Publisher     string    `json:"publisher"`
	Title         string    `json:"title"`
	Image         string    `json:"image"`
	Site          string    `json:"site"`
	Text          string    `json:"text"`
	URL           string    `json:"url"`
}

// SearchCryptoNews searches for cryptocurrency news by symbol
func (c *Client) SearchCryptoNews(symbols string, page, limit int, from, to string) ([]SearchCryptoNewsResponse, error) {
	if symbols == "" {
		return nil, fmt.Errorf("symbols is required")
	}
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
		"symbols": symbols,
		"page":    strconv.Itoa(page),
		"limit":   strconv.Itoa(limit),
	}

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}

	url := c.BaseURL + "/news/crypto"

	var result []SearchCryptoNewsResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
