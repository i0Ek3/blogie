package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	msg     string
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		message := fmt.Sprintf("error code %d already exist, change to another", code)
		panic(any(message))
	}
	codes[code] = msg

	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code: %d, error message: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

// StatusCode converts common error code into http status code
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case InternalServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code(): // 400
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case DatabaseError.Code():
		return http.StatusInternalServerError
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code(): // 401
		return http.StatusUnauthorized
	case TooManyRequests.Code(): // 429
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
