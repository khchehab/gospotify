package gospotify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

// Client is a Spotify API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
	logger     *slog.Logger
}

// NewClient creates a new Spotify API client.
func NewClient(ts oauth2.TokenSource, logger *slog.Logger) *Client {
	if logger == nil {
		logger = slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	return &Client{
		baseURL:    fmt.Sprintf("%s%s", spotifyBaseUrl, spotifyAPIVersion),
		httpClient: oauth2.NewClient(context.Background(), ts),
		logger:     logger,
	}
}

// get performs a GET request to the specified URL and unmarshals the response into the provided response object.
func (c *Client) get(ctx context.Context, endpoint string, response any, opts ...QueryOption) error {
	req, err := c.prepareRequest(ctx, http.MethodGet, endpoint, nil, "", opts...)
	if err != nil {
		return err
	}
	return c.execute(req, response)
}

// post performs a POST request to the specified URL.
func (c *Client) post(ctx context.Context, endpoint string, body any, response any, opts ...QueryOption) error {
	req, err := c.prepareRequest(ctx, http.MethodPost, endpoint, body, "", opts...)
	if err != nil {
		return err
	}
	return c.execute(req, response)
}

// put performs a PUT request to the specified URL.
func (c *Client) put(ctx context.Context, endpoint string, body any, contentType string, response any, opts ...QueryOption) error {
	req, err := c.prepareRequest(ctx, http.MethodPut, endpoint, body, contentType, opts...)
	if err != nil {
		return err
	}
	return c.execute(req, response)
}

// delete performs a DELETE request to the specified URL.
func (c *Client) delete(ctx context.Context, endpoint string, body any, response any, opts ...QueryOption) error {
	req, err := c.prepareRequest(ctx, http.MethodDelete, endpoint, body, "", opts...)
	if err != nil {
		return err
	}
	return c.execute(req, response)
}

// prepareRequest prepares the request object to call.
func (c *Client) prepareRequest(ctx context.Context, method string, endpoint string, body any, contentType string, opts ...QueryOption) (*http.Request, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("no endpoint provided")
	}
	url := c.buildURL(endpoint, opts...)

	c.logger.Debug("preparing request object", "method", method, "url", url)

	bodyReader, contentType, err := prepareBody(body, contentType)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	return req, nil
}

// execute will execute a given request, read the response and return it if found.
func (c *Client) execute(req *http.Request, response any) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			c.logger.Error("failed to close response body", "error", closeErr)
		}
	}(res.Body)

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	c.logger.Debug("response", "status code", res.StatusCode)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		var errResponse *ErrorResponse
		if err = json.Unmarshal(b, &errResponse); err != nil {
			return fmt.Errorf("unexpected Spotify error response (HTTP %d): %w - body: %s", res.StatusCode, err, string(b))
		}
		return errResponse
	}

	if res.StatusCode == http.StatusNoContent {
		return nil
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

// prepareBody prepares the request body and the content type to pass in to the request.
func prepareBody(body any, contentType string) (io.Reader, string, error) {
	if body == nil {
		return nil, "", nil
	}

	var b []byte
	var err error

	switch t := body.(type) {
	case []byte:
		b = t
	default:
		if b, err = json.Marshal(body); err != nil {
			return nil, "", err
		}
	}

	if contentType == "" {
		contentType = "application/json"
	}

	return bytes.NewReader(b), contentType, nil
}
