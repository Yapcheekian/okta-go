package okta

import (
	"context"
	"net/http"
)

type Client struct {
	token      string
	oktaDomain string
	httpClient *http.Client
}

func New(ctx context.Context, token string, oktaDomain string, httpClient *http.Client) *Client {
	return &Client{
		token:      token,
		oktaDomain: oktaDomain,
		httpClient: httpClient,
	}
}
