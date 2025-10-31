package service

import (
	"service-cashier/internal/model"
	"service-cashier/internal/repository"
)

// MenuService handles menu business logic
type MenuService struct {
	menuRepo *repository.MenuRepository
}

// NewMenuService creates a new MenuService instance
func NewMenuService(menuRepo *repository.MenuRepository) *MenuService {
	return &MenuService{menuRepo: menuRepo}
}

// GetAllMenus retrieves all menu items
func (s *MenuService) GetAllMenus() ([]model.Menu, error) {
	return s.menuRepo.GetAll()
}

// GetMenuByID retrieves a menu item by ID
func (s *MenuService) GetMenuByID(id uint) (*model.Menu, error) {
	return s.menuRepo.FindByID(id)
}

// CreateMenu creates a new menu item
func (s *MenuService) CreateMenu(menu *model.Menu) error {
	return s.menuRepo.Create(menu)
}

// UpdateMenu updates an existing menu item
func (s *MenuService) UpdateMenu(menu *model.Menu) error {
	return s.menuRepo.Update(menu)
}

// DeleteMenu deletes a menu item by ID
func (s *MenuService) DeleteMenu(id uint) error {
	return s.menuRepo.Delete(id)
}
