package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code"`
	Msg     string   `json:"msg"`
	Details []string `json:"details"`
}

var codes = map[int]struct{}{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("code %d is exsit, please change one", code))
	}
	codes[code] = struct{}{}
	return &Error{Code: code, Msg: msg}
}

// Error return a error string
func (e *Error) Error() string {
	return fmt.Sprintf("code：%d, msg:：%s", e.Code, e.Msg)
}

// Code return error code
func (e *Error) GetCode() int {
	return e.Code
}

// Msg return error msg
func (e *Error) GetMsg() string {
	return e.Msg
}

// Details return more error details
func (e *Error) GetDetails() []string {
	return e.Details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Details = []string{}
	newError.Details = append(newError.Details, details...)
	return &newError
}

// Err represents an error
type Err struct {
	Code int
	Msg  string
	Err  error
}

func (e *Err) Error() string {
	return fmt.Sprintf("code：%d, msg:：%s, error:：%s", e.Code, e.Msg, e.Err)
}

// DecodeErr 对错误进行解码，返回错误code和错误提示
func DecodeErr(err error) (int, string) {
	if err == nil {
		return Success.Code, Success.Msg
	}

	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Msg
	case *Error:
		return typed.Code, typed.Msg
	default:
	}

	return ServerError.Code, ServerError.Msg
}

// StatusCode trans err code to http status code
func (e *Error) StatusCode() int {
	switch e.GetCode() {
	case Success.GetCode():
		return http.StatusOK
	case ServerError.GetCode():
		return http.StatusInternalServerError
	case ParamBindError.GetCode():
		return http.StatusBadRequest
	case InvalidTokenError.GetCode():
		fallthrough
	case TokenTimeoutError.GetCode():
		return http.StatusUnauthorized
	case TooManyRequest.GetCode():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
