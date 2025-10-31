package handler

import (
	"service-cashier/internal/service"
	"service-cashier/pkg/utils"

	"github.com/gin-gonic/gin"
)

// AuthHandler handles authentication-related HTTP requests
type AuthHandler struct {
	userService *service.UserService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

// Login handles the login endpoint
// POST /api/login
func (h *AuthHandler) Login(c *gin.Context) {
	var req service.LoginRequest

	// Bind JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request payload")
		return
	}

	// Authenticate user
	response, err := h.userService.Login(&req)
	if err != nil {
		utils.UnauthorizedResponse(c, err.Error())
		return
	}

	// Return success response with token
	utils.SuccessResponse(c, "Login successful", response)
}
