package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/amburskui/httpserver/internal/infrastructure/persistence"
)

type resultError struct {
	Code    Code `json:"code"`
	Message any  `json:"message"`
}

type Code int32

const (
	CodeUnexpectedError Code = iota
	CodeBadRequestError
	CodeNotFoundError
)

func customErrorHandler(err error, c echo.Context) {
	errHTTP := &echo.HTTPError{}

	switch {
	case errors.As(err, &errHTTP):
		c.JSON(errHTTP.Code, resultError{
			Code:    CodeBadRequestError,
			Message: errHTTP.Message,
		})

	case errors.Is(err, persistence.ErrNotFound):
		c.JSON(http.StatusNotFound, resultError{
			Code:    CodeBadRequestError,
			Message: err.Error(),
		})

	default:
		c.JSON(http.StatusInternalServerError, resultError{
			Code:    CodeUnexpectedError,
			Message: err.Error(),
		})
	}
}
