package middleware

import (
	"fmt"
	"log/slog"

	domainErrors "github.com/Diegonr1791/GymBro/internal/domain/errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// ErrorHandler is a middleware to handle errors centrally.
// It logs the error and returns a standardized JSON response.
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			// Log the full error with stack trace
			slog.Error("An error occurred", "error", fmt.Sprintf("%+v", err))

			var appErr *domainErrors.AppError
			if errors.As(err, &appErr) {
				// The error is a known application error
				response := domainErrors.ErrorResponse{
					Code:    appErr.Code,
					Message: appErr.Message,
				}
				c.JSON(appErr.HTTPStatus, gin.H{"error": response})
			} else {
				// The error is an unexpected, non-application error
				// Respond with a generic 500 internal server error
				genericError := domainErrors.ErrInternalServer
				response := domainErrors.ErrorResponse{
					Code:    genericError.Code,
					Message: genericError.Message,
				}
				c.JSON(genericError.HTTPStatus, gin.H{"error": response})
			}
		}
	}
}
