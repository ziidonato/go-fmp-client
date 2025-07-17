package go_fmp

import (
	"fmt"
	"strconv"
	"time"
)

// SearchPressReleasesResponse represents the response from the Search Press Releases API
type SearchPressReleasesResponse struct {
	Symbol        string    `json:"symbol"`
	PublishedDate time.Time `json:"publishedDate"`
	Title         string    `json:"title"`
	Image         string    `json:"image"`
	Site          string    `json:"site"`
	Text          string    `json:"text"`
	URL           string    `json:"url"`
}

// SearchPressReleases searches for company press releases by symbol
func (c *Client) SearchPressReleases(symbols string, page, limit int, from, to string) ([]SearchPressReleasesResponse, error) {
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

	url := c.BaseURL + "/news/press-releases"

	var result []SearchPressReleasesResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return result, nil
}
