package go_fmp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HTTPClient is an interface for making HTTP requests.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client is the main struct for interacting with the Financial Modeling Prep API.
type Client struct {
	HTTPClient HTTPClient
	APIKey     string
}

// NewClient creates a new Client. If httpClient is nil, http.DefaultClient is used.
func NewClient(httpClient HTTPClient, apiKey string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	newClient := &Client{
		HTTPClient: httpClient,
		APIKey:     apiKey,
	}

	err := testClient(newClient)
	if err != nil {
		return nil, fmt.Errorf("failed to test client: %v", err)
	}

	return newClient, nil
}

func testClient(client *Client) error {
	testResp, err := client.doRequest("https://financialmodelingprep.com/api/v3/quote/AAPL", map[string]string{})
	if err != nil {
		return fmt.Errorf("failed to perform test request: %v", err)
	}

	defer testResp.Body.Close()
	var body map[string]any
	err = json.NewDecoder(testResp.Body).Decode(&body)

	if err != nil {
		return fmt.Errorf("failed to decode test response body: %v", err)
	}

	if body["error"] != nil {
		return fmt.Errorf("failed to perform test request: %v", body["error"])
	}

	return nil
}

// doRequest is a unified method for making HTTP requests to the Financial Modeling Prep API
func (c *Client) doRequest(url string, params map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Add API key to parameters
	if params == nil {
		params = make(map[string]string)
	}
	params["apikey"] = c.APIKey

	// Add query parameters
	query := req.URL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// Set standard headers
	req.Header.Set("User-Agent", "fmp-go-client")
	req.Header.Set("Accept", "application/json")

	// Make the request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	// Check response status
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	return resp, nil
}

func editQueryParams(req *http.Request, params map[string]string) {
	query := req.URL.Query()
	for key, value := range params {
		query.Set(key, value)
	}
	req.URL.RawQuery = query.Encode()
}

func (c *Client) get(url string, params map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	params["apikey"] = c.APIKey
	editQueryParams(req, params)
	return c.HTTPClient.Do(req)
}
