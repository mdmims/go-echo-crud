package repository

import (
	"encoding/json"
	"fmt"
)

const (
	ECONFLICT       = "conflict"
	EINTERNAL       = "internal error"
	EINVALID        = "invalid"
	ENOTFOUND       = "not found"
	ENOTIMPLEMENTED = "not implemented"
	EUNAUTHORIZED   = "unauthorized"
)

type MyError struct {
	MyError error
}

func (err MyError) Error() string {
	return fmt.Sprintf("error: message=%s", err.MyError)
}

func (err MyError) MarshalJSON() ([]byte, error) {
	return json.Marshal(err.MyError.Error())
}

type HTTPError struct {
	Cause  error  `json:"Cause"`
	Detail string `json:"Detail"`
	Status int    `json:"Status"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

func NewHTTPError(err error, status int, detail string) error {
	e := MyError{MyError: err}
	return &HTTPError{
		Cause:  e,
		Detail: detail,
		Status: status,
	}
}
