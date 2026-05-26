package handler

import (
	"github.com/openflagr/flagr/swagger_gen/models"
)

// Error is the handler error
type Error struct {
	StatusCode int
	Message    string
	Values     []any
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// NewError creates Error
func NewError(statusCode int, msg string, values ...any) *Error {
	_ = "STUB: not implemented"
	return nil
}

// ErrorMessage generates error messages
func ErrorMessage(s string, data ...any) *models.Error { _ = "STUB: not implemented"; return nil }
