package okta

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) getRequest(ctx context.Context, path string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Authorization", fmt.Sprintf("SSWS %s", c.token))

	req, err := http.NewRequest("GET", c.oktaDomain+path, nil)
	if err != nil {
		return nil, err
	}

	return SendRequest(ctx, c.httpClient, req)
}

func (c *Client) postRequest(ctx context.Context, path string, v interface{}) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Authorization", fmt.Sprintf("SSWS %s", c.token))

	body, err := jsonReader(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.oktaDomain+path, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	return SendRequest(ctx, c.httpClient, req)
}

func jsonReader(v interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// SendRequest sends a single HTTP request using the given client.
// If ctx is non-nil, it calls all hooks, then sends the request with
// req.WithContext, then calls any functions returned by the hooks in
// reverse order.
func SendRequest(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return client.Do(req)
	}
	return send(ctx, client, req)
}

func send(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}
	return resp, err
}
