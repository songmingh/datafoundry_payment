package pkg

import (
	"fmt"
)

const (
	ErrCodeOK                 = 1200
	ErrCodeBadRequest         = 1400
	ErrCodeActionNotSupport   = 14003
	ErrCodeInvalidToken       = 14004
	ErrCodeForbidden          = 1403
	ErrCodePermissionDenied   = 14030
	ErrCodeNotFound           = 1404
	ErrCodePlanNotFound       = 14040
	ErrCodeMethodNotAllowed   = 1405
	ErrCodeServiceUnavailable = 1503

	ErrCodeUnknownError = 140010
)

var errText = map[int]string{
	ErrCodeOK:                 "OK",
	ErrCodeBadRequest:         "Bad request",
	ErrCodeActionNotSupport:   "Not supported action",
	ErrCodeInvalidToken:       "Invalid token",
	ErrCodeForbidden:          "Forbidden",
	ErrCodePermissionDenied:   "Permission denied",
	ErrCodeNotFound:           "Not found",
	ErrCodePlanNotFound:       "No such plan",
	ErrCodeMethodNotAllowed:   "Method not allowed",
	ErrCodeServiceUnavailable: "Service unavailable",

	ErrCodeUnknownError: "Unknown error",
}

func ErrText(code int) string {
	return errText[code]
}

type ErrorMessage struct {
	Code    int
	Message string
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("[%v] %v", e.Code, e.Message)
}

func (e *ErrorMessage) New(code int) error {
	e.Code = code

	e.Message = ErrText(code)
	return e
}

func ErrorNew(code int) error {
	var e ErrorMessage
	return e.New(code)
}
