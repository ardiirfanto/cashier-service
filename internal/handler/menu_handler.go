package handler

import (
	"service-cashier/internal/service"
	"service-cashier/pkg/utils"

	"github.com/gin-gonic/gin"
)

// MenuHandler handles menu-related HTTP requests
type MenuHandler struct {
	menuService *service.MenuService
}

// NewMenuHandler creates a new MenuHandler instance
func NewMenuHandler(menuService *service.MenuService) *MenuHandler {
	return &MenuHandler{menuService: menuService}
}

// GetMenus handles the get all menus endpoint
// GET /api/menus
func (h *MenuHandler) GetMenus(c *gin.Context) {
	// Retrieve all menus
	menus, err := h.menuService.GetAllMenus()
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve menus")
		return
	}

	// Return success response
	utils.SuccessResponse(c, "Menus retrieved successfully", menus)
}
