package gospotify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

// Client is a Spotify API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Spotify API client.
func NewClient(ts oauth2.TokenSource) *Client {
	return &Client{
		baseURL:    fmt.Sprintf("%s%s", SpotifyBaseUrl, SpotifyAPIVersion),
		httpClient: oauth2.NewClient(context.Background(), ts),
	}
}

// get performs a GET request to the specified URL and unmarshals the response into the provided response object.
func (c *Client) get(ctx context.Context, endpoint string, response any, opts ...QueryOption) error {
	url := c.buildURL(endpoint, opts...)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			fmt.Println("error closing the response body:", closeErr)
		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		var errResponse *ErrorResponse
		if err = json.Unmarshal(b, &errResponse); err != nil {
			return err
		}
		return errResponse
	}

	fmt.Println(string(b))

	if err = json.Unmarshal(b, response); err != nil {
		return err
	}
	return nil
}

// buildURL builds a URL for the specified endpoint and query parameters.
func (c *Client) buildURL(endpoint string, opts ...QueryOption) string {
	url := fmt.Sprintf("%s%s", c.baseURL, endpoint)

	p := applyQueryParameters(opts...)
	if q := p.toQuery(); q != "" {
		if strings.ContainsRune(endpoint, '?') {
			url += "&" + q
		} else {
			url += "?" + q
		}
	}

	return url
}
