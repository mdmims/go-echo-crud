package repository

import (
	"fmt"
	"net/http"
)

const (
	ENOTFOUND  = "not found"
	EINTERNAL  = "server error"
	EMALFORMED = "schema invalid"
)

type Error struct {
	Status  int    `json:"Status"`
	Cause   string `json:"-"`
	Message string `json:"Message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("wtf error: status=%v cause=%s message=%s", e.Status, e.Cause, e.Message)
}

func NewError(err error, code int, errMessage string) Error {
	e := Error{}
	e.Status = code
	e.Cause = err.Error()
	e.Message = errMessage
	return e
}

func NotFound() Error {
	return Error{
		Status:  http.StatusNotFound,
		Cause:   ENOTFOUND,
		Message: ENOTFOUND,
	}
}
