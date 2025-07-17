package go_fmp

import (
	"fmt"
	"strconv"
	"time"
)

// PressReleasesResponse represents the response from the press releases API
type PressReleasesResponse struct {
	Symbol        string    `json:"symbol"`
	PublishedDate time.Time `json:"publishedDate"`
	Publisher     string    `json:"publisher"`
	Title         string    `json:"title"`
	Image         string    `json:"image"`
	Site          string    `json:"site"`
	Text          string    `json:"text"`
	URL           string    `json:"url"`
}

// GetPressReleases retrieves official company press releases
func (c *Client) GetPressReleases(page, limit int, from, to string) ([]PressReleasesResponse, error) {
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

	if from != "" {
		params["from"] = from
	}
	if to != "" {
		params["to"] = to
	}

	url := c.BaseURL + "/news/press-releases-latest"

	var result []PressReleasesResponse
	err := c.doRequest(url, params, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
