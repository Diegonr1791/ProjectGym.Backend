package errors

import (
	"fmt"
	"net/http"
)

// =================================================================
// ||                ERRORES DE NEGOCIO ESTÁNDAR                  ||
// =================================================================
// Estos errores son la "lingua franca" entre los casos de uso y los handlers.
// Representan conceptos de negocio, no fallos de implementación.

// ErrorResponse defines the structure of the JSON error response.
// ErrorResponse is a standard error response for API endpoints.
// @Description Standard error response
// @name ErrorResponse
// @produce json
// @Success 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// AppError represents a custom application error.
type AppError struct {
	HTTPStatus int
	Code       string
	Message    string
	Err        error
}

// Error implements the error interface for AppError.
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// Unwrap provides compatibility with errors.Unwrap.
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewAppError creates a new AppError.
func NewAppError(httpStatus int, code, message string, err error) *AppError {
	return &AppError{
		HTTPStatus: httpStatus,
		Code:       code,
		Message:    message,
		Err:        err,
	}
}

// Standard application errors.
var (
	ErrNotFound               = NewAppError(http.StatusNotFound, "NOT_FOUND", "Resource not found.", nil)
	ErrConflict               = NewAppError(http.StatusConflict, "CONFLICT", "A conflict occurred.", nil)
	ErrInternalServer         = NewAppError(http.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "An internal server error occurred.", nil)
	ErrBadRequest             = NewAppError(http.StatusBadRequest, "BAD_REQUEST", "The request is invalid.", nil)
	ErrUnauthorized           = NewAppError(http.StatusUnauthorized, "UNAUTHORIZED", "Request is not authorized.", nil)
	ErrForbidden              = NewAppError(http.StatusForbidden, "FORBIDDEN", "Access forbidden. Insufficient permissions.", nil)
	ErrEmailAlreadyExists     = NewAppError(http.StatusConflict, "EMAIL_ALREADY_EXISTS", "The provided email is already in use.", nil)
	ErrInvalidCredentials     = NewAppError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid credentials provided.", nil)
	ErrInvalidRefreshToken    = NewAppError(http.StatusUnauthorized, "INVALID_REFRESH_TOKEN", "The provided refresh token is invalid or has expired.", nil)
	ErrSystemRoleNotDeletable = NewAppError(http.StatusBadRequest, "SYSTEM_ROLE_NOT_DELETABLE", "System roles cannot be deleted.", nil)

	// User validation errors
	ErrEmailRequired         = NewAppError(http.StatusBadRequest, "EMAIL_REQUIRED", "Email is required.", nil)
	ErrInvalidEmailFormat    = NewAppError(http.StatusBadRequest, "INVALID_EMAIL_FORMAT", "Invalid email format.", nil)
	ErrPasswordRequired      = NewAppError(http.StatusBadRequest, "PASSWORD_REQUIRED", "Password is required.", nil)
	ErrPasswordTooShort      = NewAppError(http.StatusBadRequest, "PASSWORD_TOO_SHORT", "Password must be at least 8 characters long.", nil)
	ErrPasswordTooLong       = NewAppError(http.StatusBadRequest, "PASSWORD_TOO_LONG", "Password must not exceed 128 characters.", nil)
	ErrPasswordWeak          = NewAppError(http.StatusBadRequest, "PASSWORD_WEAK", "Password must contain at least one uppercase letter, one lowercase letter, and one number.", nil)
	ErrNameRequired          = NewAppError(http.StatusBadRequest, "NAME_REQUIRED", "Name is required.", nil)
	ErrNameTooShort          = NewAppError(http.StatusBadRequest, "NAME_TOO_SHORT", "Name must be at least 2 characters long.", nil)
	ErrNameTooLong           = NewAppError(http.StatusBadRequest, "NAME_TOO_LONG", "Name must not exceed 100 characters.", nil)
	ErrInvalidNameCharacters = NewAppError(http.StatusBadRequest, "INVALID_NAME_CHARACTERS", "Name contains invalid characters.", nil)
	ErrRoleRequired          = NewAppError(http.StatusBadRequest, "ROLE_REQUIRED", "Role is required.", nil)
	ErrUserInactive          = NewAppError(http.StatusUnauthorized, "USER_INACTIVE", "User account is inactive.", nil)
	ErrUserAlreadyDeleted    = NewAppError(http.StatusBadRequest, "USER_ALREADY_DELETED", "User is already deleted.", nil)
	ErrUserNotDeleted        = NewAppError(http.StatusBadRequest, "USER_NOT_DELETED", "User is not deleted.", nil)
	ErrEmailSoftDeleted      = NewAppError(http.StatusConflict, "EMAIL_SOFT_DELETED", "Email belongs to a deleted user.", nil)
)
