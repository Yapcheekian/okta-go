package okta

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func getRequest(ctx context.Context, client *http.Client, oktaDomain, path string) (*http.Response, error) {
	req, err := http.NewRequest("GET", oktaDomain+path, nil)
	if err != nil {
		return nil, err
	}

	return sendRequest(ctx, client, req)
}

func postRequest(ctx context.Context, client *http.Client, oktaDomain, path string, v interface{}) (*http.Response, error) {
	body, err := jsonReader(v)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", oktaDomain+path, body)
	if err != nil {
		return nil, err
	}

	return sendRequest(ctx, client, req)
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
func sendRequest(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
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

func parseResponse(body io.ReadCloser, intf interface{}, debug bool) error {
	var decoder *json.Decoder
	if debug {
		response, err := ioutil.ReadAll(body)
		if err != nil {
			return err
		}
		log.Println(string(response))
		decoder = json.NewDecoder(bytes.NewReader(response))
	} else {
		decoder = json.NewDecoder(body)
	}
	if err := decoder.Decode(&intf); err != nil {
		return err
	}
	return nil
}
