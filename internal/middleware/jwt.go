package middleware

import (
	"service-cashier/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuth is a middleware that validates JWT tokens
func JWTAuth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.UnauthorizedResponse(c, "Authorization header is required")
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.UnauthorizedResponse(c, "Invalid authorization header format")
			c.Abort()
			return
		}

		// Extract token
		tokenString := parts[1]

		// Validate token
		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			utils.UnauthorizedResponse(c, "Invalid or expired token")
			c.Abort()
			return
		}

		// Attach user information to context
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)

		// Continue to next handler
		c.Next()
	}
}

// GetUserID retrieves the user ID from the Gin context
func GetUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	id, ok := userID.(uint)
	return id, ok
}

// GetUsername retrieves the username from the Gin context
func GetUsername(c *gin.Context) (string, bool) {
	username, exists := c.Get("username")
	if !exists {
		return "", false
	}

	name, ok := username.(string)
	return name, ok
}
