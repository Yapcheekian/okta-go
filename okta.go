package okta

import (
	"context"
	"fmt"
	"net/http"
)

type Client struct {
	OktaDomain string
	Client     *http.Client
	Apps       *AppsService
	Users      *UsersService
	Logs       *LogsService
}

func New(ctx context.Context, token string, oktaDomain string, httpClient *http.Client) *Client {
	var rt http.RoundTripper
	if httpClient.Transport != nil {
		rt = httpClient.Transport
	} else {
		rt = http.DefaultTransport
	}

	transport := &customTransport{
		rt:    rt,
		token: token,
	}
	httpClient.Transport = transport
	s := &Client{
		OktaDomain: oktaDomain,
		Client:     httpClient,
	}

	s.Users = NewUsersService(s)
	s.Logs = NewLogsService(s)
	s.Apps = NewAppsService(s)

	return s
}

type customTransport struct {
	rt    http.RoundTripper
	token string
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := *req
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")
	reqHeaders.Set("Content-Type", "application/json")
	reqHeaders.Set("Authorization", fmt.Sprintf("SSWS %s", t.token))

	newReq.Header = reqHeaders
	for k, vv := range req.Header {
		newReq.Header[k] = vv
	}

	for k, v := range newReq.Header {
		fmt.Println(k, v)
	}
	return t.rt.RoundTrip(&newReq)
}
