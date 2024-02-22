package httpResponse

import (
	"fmt"
	"net/http"
	"strings"

	"InitiaRe-website/constant"
	"InitiaRe-website/pkg/utils"
)

// Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Rest error struct
type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCause  interface{} `json:"cause,omitempty"`
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - cause: %v", e.ErrStatus, e.ErrError, e.ErrCause)
}

// Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrCause
}

// New Rest Error
func NewRestError(status int, err string, cause interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCause:  cause,
	}
}

// New Bad Request Error
func NewBadRequestError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  constant.STATUS_CODE_BAD_REQUEST,
		ErrCause:  cause,
	}
}

// New Not Found Error
func NewNotFoundError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  constant.STATUS_CODE_NOT_FOUND,
		ErrCause:  cause,
	}
}

// New Unauthorized Error
func NewUnauthorizedError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  constant.STATUS_CODE_UNAUTHORIZED,
		ErrCause:  cause,
	}
}

// New Forbidden Error
func NewForbiddenError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  constant.STATUS_CODE_FORBIDDEN,
		ErrCause:  cause,
	}
}

// New Internal Server Error
func NewInternalServerError(cause interface{}) RestErr {
	result := RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  constant.STATUS_CODE_INTERNAL_SERVER,
		ErrCause:  cause,
	}
	return result
}

func NewRequestTimeoutError(cause interface{}) RestErr {
	result := RestError{
		ErrStatus: http.StatusRequestTimeout,
		ErrError:  constant.STATUS_CODE_REQUEST_TIMEOUT,
		ErrCause:  cause,
	}
	return result
}

func ParseError(err error) RestErr {
	if strings.Contains(err.Error(), constant.STATUS_CODE_BAD_REQUEST) {
		return NewBadRequestError(utils.GetErrorMessage(err))
	}
	if strings.Contains(err.Error(), constant.STATUS_CODE_NOT_FOUND) {
		return NewNotFoundError(utils.GetErrorMessage(err))
	}
	if strings.Contains(err.Error(), constant.STATUS_CODE_UNAUTHORIZED) {
		return NewUnauthorizedError(utils.GetErrorMessage(err))
	}
	if strings.Contains(err.Error(), constant.STATUS_CODE_FORBIDDEN) {
		return NewForbiddenError(utils.GetErrorMessage(err))
	}
	if strings.Contains(err.Error(), constant.STATUS_CODE_REQUEST_TIMEOUT) {
		return NewRequestTimeoutError(utils.GetErrorMessage(err))
	}
	return NewInternalServerError(utils.GetErrorMessage(err))
}
