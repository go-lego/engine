package error

import (
	"encoding/json"
	"fmt"

	"github.com/go-lego/engine/log"
)

// Error business error, with code and message
type Error struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// New create new Error
func New(code int64, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Error to string
func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

var (
	// SystemMin minimum system error code
	SystemMin = 1

	// SystemMax maximum system error code
	SystemMax = 100

	/**
	 * @apiDefine ErrSystemError
	 * @apiError (Error <100) {Number} ErrSystemError <code>1</code>. System error
	 */

	// ErrSystemError 1
	ErrSystemError = New(1, "System Error")

	/**
	 * @apiDefine ErrInputInvalid
	 * @apiError (Error <100) {Number} ErrInputInvalid <code>2</code>. Invalid input
	 */

	// ErrInputInvalid 2
	ErrInputInvalid = New(2, "Input invalid")
)

// NewSystemError create system error
func NewSystemError(code int, message string) *Error {
	if code < SystemMin || code > SystemMax {
		log.Error("Bad system error code provided(%d), must between %d - %d", code, SystemMin, SystemMax)
	}
	return &Error{
		Code:    1,
		Message: fmt.Sprintf("System error:%s", message),
	}
}
