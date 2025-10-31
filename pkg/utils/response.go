package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse sends a successful JSON response
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error JSON response
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Data:    nil,
	})
}

// CreatedResponse sends a created (201) JSON response
func CreatedResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// BadRequestResponse sends a 400 Bad Request error response
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusBadRequest, message)
}

// UnauthorizedResponse sends a 401 Unauthorized error response
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusUnauthorized, message)
}

// NotFoundResponse sends a 404 Not Found error response
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusNotFound, message)
}

// InternalServerErrorResponse sends a 500 Internal Server Error response
func InternalServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, http.StatusInternalServerError, message)
}
