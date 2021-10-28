package common_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client interface {
	Get(ctx context.Context, address, token string, params url.Values) ([]byte, int, error)
}

type client struct {
	baseUrl string
	client  *http.Client
}

func NewClient(httpClient *http.Client) Client {
	return &client{client: httpClient}
}

func (r *client) Get(ctx context.Context, address, token string, params url.Values) ([]byte, int, error) {
	panic("implement me")
	return r.doRequest(ctx, "GET", address, token, params, nil)
}

func (r *client) doRequest(ctx context.Context, method, address, token string, params url.Values, buf io.Reader) ([]byte, int, error) {
	u, _ := url.Parse(address)
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, 0, err
	}
	req = req.WithContext(ctx)
	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return data, resp.StatusCode, err
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
	}
}
