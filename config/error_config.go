package config

import "net/http"

type ErrorHttp struct {
	errorCode string
	httpCode  int
}

func (e *ErrorHttp) Error() string {
	return e.errorCode
}

func (e *ErrorHttp) HTTPCode() int {
	return e.httpCode
}

func New(errorCode string, httpCode int) error {
	return &ErrorHttp{errorCode: errorCode, httpCode: httpCode}
}

var INVALID_PARAMETER_IN_URL = New("invalid_parameter_in_url", http.StatusBadRequest)
var TWITCH_CODE_INVALID = New("invalid_twitch_code", http.StatusBadRequest)
var USER_NOT_FOUND = New("user_not_found", http.StatusNotFound)
var PICKEMS_NOT_FOUND = New("pickems_not_found", http.StatusNotFound)
var NO_BODY_FOUND = New("no_body_found", http.StatusBadRequest)
var FAILED_TO_READ_BODY = New("failed_to_read_body", http.StatusBadRequest)
var NOT_YOUR_PICKEMS = New("not_your_pickems", http.StatusForbidden)
var QUESTIONS_NOT_FOUND = New("questions_not_found", http.StatusNotFound)
