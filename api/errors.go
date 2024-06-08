package api

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	if apiError, ok := err.(*Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}

	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements error interface
func (e Error) Error() string {
	return e.Message
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func ErrInvalidID() *Error {
	return NewError(http.StatusBadRequest, "invalid id")
}

func ErrUnauthorized() *Error {
	return NewError(http.StatusUnauthorized, "unauthorized")
}

func ErrBadRequest() *Error {
	return NewError(http.StatusBadRequest, "invalid JSON request")
}

func ErrResourceNotFound(resource string) *Error {
	return NewError(http.StatusNotFound, fmt.Sprintf("%s not found", resource))
}
