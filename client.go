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
	body, err := doRequest[[]map[string]any](client, "https://financialmodelingprep.com/api/v3/quote/AAPL", map[string]string{})
	if err != nil {
		return fmt.Errorf("failed to perform test request: %v", err)
	}

	if len(body) > 0 && body[0]["error"] != nil {
		return fmt.Errorf("failed to perform test request: %v", body[0]["error"])
	}

	return nil
}

// doRequest is a unified method for making HTTP requests to the Financial Modeling Prep API
// It handles the complete request-response cycle including JSON unmarshaling
func doRequest[T any](c *Client, url string, params map[string]string) (T, error) {
	var result T
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, fmt.Errorf("failed to create request: %v", err)
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
		return result, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return result, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}

	// Unmarshal the response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("error decoding response: %w", err)
	}

	return result, nil
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
