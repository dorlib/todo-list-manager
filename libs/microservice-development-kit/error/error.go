package error

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const DefaultErrorMsg = "unexpected error"

type AppError struct {
	Code   int    `json:"-"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

func (a *AppError) Error() string {
	var parts []string

	if a.Title != "" {
		parts = append(parts, a.Title)
	}
	if a.Detail != "" {
		parts = append(parts, a.Detail)
	}

	return strings.Join(parts, " - ")
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

func NewInternalServerError(bodyError string) *AppError {
	return &AppError{
		Code:   http.StatusInternalServerError,
		Title:  DefaultErrorMsg,
		Detail: bodyError,
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
