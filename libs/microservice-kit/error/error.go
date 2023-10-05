package error

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type AppError struct {
	Code   int    `json:"-"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

const DefaultErrorMsg = "unexpected error"

func (a *AppError) Error() string {
	sb := strings.Builder{}
	appendDash := false

	if a.Title != "" {
		sb.WriteString(a.Title)

		appendDash = true
	}

	if a.Detail != "" {
		if appendDash {
			sb.WriteString(" - ")
		}

		sb.WriteString(a.Detail)

		appendDash = true
	}

	return sb.String()
}

func (a *AppError) IsResourceNotFound() bool {
	return a.Code == http.StatusNotFound
}

func (a *AppError) IsInternalServerError() bool {
	return a.Code == http.StatusInternalServerError
}

func (a *AppError) IsInvalidRequest() bool {
	return a.Code == http.StatusBadRequest
}

func (a *AppError) IsConflictError() bool {
	return a.Code == http.StatusConflict
}

func NewResourceNotFound(bodyError string) *AppError {
	return &AppError{
		Code:   http.StatusNotFound,
		Title:  "resource not found",
		Detail: bodyError,
	}
}

func NewInternalServerError() *AppError {
	return &AppError{
		Code:  http.StatusInternalServerError,
		Title: DefaultErrorMsg,
	}
}

func NewInvalidRequest(bodyError string) *AppError {
	return &AppError{
		Code:   http.StatusBadRequest,
		Title:  "Invalid Request",
		Detail: bodyError,
	}
}

func NewConflictError(bodyError string) *AppError {
	return &AppError{
		Code:   http.StatusConflict,
		Title:  "Conflict",
		Detail: bodyError,
	}
}

func NewHTTPErrorToAppError(echoErr *echo.HTTPError) *AppError {
	return &AppError{
		Code:  echoErr.Code,
		Title: echoErr.Message.(string),
	}
}
