package gospotify

import (
	"bytes"
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

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var errResponse *ErrorResponse
		if err = json.Unmarshal(b, &errResponse); err != nil {
			return err
		}
		return errResponse
	}

	if err = json.Unmarshal(b, response); err != nil {
		return err
	}
	return nil
}

// post performs a POST request to the specified URL.
func (c *Client) post(ctx context.Context, endpoint string, body any, response any, opts ...QueryOption) error {
	url := c.buildURL(endpoint, opts...)

	var bodyReader io.Reader = nil
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bodyReader)
	if err != nil {
		return err
	}

	if bodyReader != nil {
		req.Header.Set("Content-Type", "application/json")
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

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var errResponse *ErrorResponse
		if err = json.Unmarshal(b, &errResponse); err != nil {
			return err
		}
		return errResponse
	}

	if response != nil {
		if err = json.Unmarshal(b, response); err != nil {
			return err
		}
	}

	return nil
}

// put performs a PUT request to the specified URL.
func (c *Client) put(ctx context.Context, endpoint string, body any, contentType string, response any, opts ...QueryOption) error {
	url := c.buildURL(endpoint, opts...)

	var bodyReader io.Reader = nil
	if body != nil {
		var b []byte
		var err error

		if raw, ok := body.([]byte); ok {
			b = raw
		} else {
			if b, err = json.Marshal(body); err != nil {
				return err
			}
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bodyReader)
	if err != nil {
		return err
	}

	if bodyReader != nil {
		if contentType != "" {
			req.Header.Set("Content-Type", contentType)
		} else {
			req.Header.Set("Content-Type", "application/json")
		}
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

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var errResponse *ErrorResponse
		if err = json.Unmarshal(b, &errResponse); err != nil {
			return err
		}
		return errResponse
	}

	if response != nil {
		if err = json.Unmarshal(b, response); err != nil {
			return err
		}
	}

	return nil
}

// delete performs a DELETE request to the specified URL.
func (c *Client) delete(ctx context.Context, endpoint string, body any, response any, opts ...QueryOption) error {
	url := c.buildURL(endpoint, opts...)

	var bodyReader io.Reader = nil
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, bodyReader)
	if err != nil {
		return err
	}

	if bodyReader != nil {
		req.Header.Set("Content-Type", "application/json")
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

	if res.StatusCode < 200 || res.StatusCode > 299 {
		var errResponse *ErrorResponse
		if err = json.Unmarshal(b, &errResponse); err != nil {
			return err
		}
		return errResponse
	}

	if response != nil {
		if err = json.Unmarshal(b, response); err != nil {
			return err
		}
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
