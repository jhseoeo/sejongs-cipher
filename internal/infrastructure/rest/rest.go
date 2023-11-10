package rest

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/go-errors/errors"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

// Post sends a POST request to the given URL with the given token and payload
func (c *Client) Post(ctx context.Context, url string, token string, payload []byte) ([]byte, error) {
	return c.createRequest(ctx, "POST", url, token, payload)
}

// Get sends a GET request to the given URL with the given token and payload
func (c *Client) Get(ctx context.Context, url string, token string) ([]byte, error) {
	return c.createRequest(ctx, "GET", url, token, nil)
}

func (c *Client) RawPost(ctx context.Context, url string, payload []byte) (*http.Response, error) {
	return c.createRawRequest(ctx, "POST", url, payload)
}

func (c *Client) RawGet(ctx context.Context, url string) (*http.Response, error) {
	return c.createRawRequest(ctx, "GET", url, nil)
}

// createRequest sends a request to the given URL with the given token and payload
func (c *Client) createRequest(ctx context.Context, method string, url string, token string, payload []byte) ([]byte, error) {
	// Create a POST Client
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Errorf("failed to create request: %w", err)
	}

	// Set the Content-Type header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("referer", "https://ko.dict.naver.com/")

	// Add headers if needed
	if token != "" {
		req.Header.Add("Authorization", "Bearer "+token)
	}

	// Send the Client
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Errorf("failed to send request: %w", err)
	}

	// Close the response body when we're done
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Errorf("failed to read response body: %w", err)
	}

	// Print the response
	return body, nil
}

func (c *Client) createRawRequest(ctx context.Context, method string, url string, payload []byte) (*http.Response, error) {
	// Create a POST Client
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Accept-Encoding", "*/*")

	// Send the Client
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Errorf("failed to send request: %w", err)
	}

	// Print the response
	return resp, nil
}
