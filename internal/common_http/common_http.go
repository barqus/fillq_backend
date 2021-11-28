package common_http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/barqus/fillq_backend/config"
	"github.com/go-chi/chi"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client interface {
	Get(ctx context.Context, address, token string, clientID string, params url.Values) ([]byte, int, error)
	Post(ctx context.Context, address, token string, params url.Values, body []byte) ([]byte, int, error)
}

type client struct {
	baseUrl string
	client  *http.Client
}

type HttpError struct {
	Message string `json:"message"`
	ErrorCode int `json:"error_code"`
}

func NewClient(httpClient *http.Client) Client {
	return &client{client: httpClient}
}

func (r *client) Get(ctx context.Context, address, token string, clientID string, params url.Values) ([]byte, int, error) {
	return r.doRequest(ctx, "GET", address, token, params, nil, clientID)
}

func (r *client) Post(ctx context.Context, address, token string, params url.Values, body []byte) ([]byte, int, error) {
	return r.doRequest(ctx, "POST", address, token, params, bytes.NewBuffer(body), "")
}

func (r *client) doRequest(ctx context.Context, method, address, token string, params url.Values, buf io.Reader, clientID string) ([]byte, int, error) {
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

	if clientID != "" {
		req.Header.Set("Client-Id", clientID)
	}

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
	}
}

func WriteErrorResponse(w http.ResponseWriter, errorMessage error) {
	code := http.StatusInternalServerError
	http, ok := errorMessage.(*config.ErrorHttp)
	if ok {
		code = http.HTTPCode()
	}
	w.Header().Add("error_code", errorMessage.Error())
	WriteJSONResponse(w, code, &HttpError{
		Message: errorMessage.Error(),
		ErrorCode: code,
	})
}
func ParseURLParamToString(r *http.Request, paramName string) (string, error) {
	stringOfParam := chi.URLParam(r, paramName)

	if len(stringOfParam) == 0 {
		return "", config.INVALID_PARAMETER_IN_URL
	}

	return stringOfParam, nil
}

func ParseURLParamToSInt(r *http.Request, paramName string) (int, error) {
	stringOfParam := chi.URLParam(r, paramName)
	if len(stringOfParam) == 0 {
		return 0, config.INVALID_PARAMETER_IN_URL
	}

	intOfParam, err := strconv.Atoi(stringOfParam)
	if err != nil {
		return 0, config.INVALID_PARAMETER_IN_URL
	}

	return intOfParam, nil
}

func ParseJSON(r *http.Request, object interface{}) error {
	if r == nil {
		return config.NO_BODY_FOUND
	}
	if r.Body == nil {
		return config.NO_BODY_FOUND
	}

	err := json.NewDecoder(r.Body).Decode(&object)
	if err != nil {
		return config.FAILED_TO_READ_BODY
	}
	if object == nil {
		return config.FAILED_TO_READ_BODY
	}

	return nil
}