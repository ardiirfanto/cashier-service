package repository

import (
	"service-cashier/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// MenuRepository handles menu data access operations
type MenuRepository struct {
	db *gorm.DB
}

// NewMenuRepository creates a new MenuRepository instance
func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

// GetAll retrieves all menu items
func (r *MenuRepository) GetAll() ([]model.Menu, error) {
	var menus []model.Menu
	err := r.db.Find(&menus).Error
	return menus, err
}

// FindByID retrieves a menu item by ID
func (r *MenuRepository) FindByID(id uint) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// FindByIDWithLock retrieves a menu item by ID with row-level locking for updates
// This is used during checkout to prevent race conditions when updating stock
func (r *MenuRepository) FindByIDWithLock(tx *gorm.DB, id uint) (*model.Menu, error) {
	var menu model.Menu
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&menu, id).Error
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

// Create creates a new menu item
func (r *MenuRepository) Create(menu *model.Menu) error {
	return r.db.Create(menu).Error
}

// Update updates an existing menu item
func (r *MenuRepository) Update(menu *model.Menu) error {
	return r.db.Save(menu).Error
}

// UpdateStock updates the stock of a menu item within a transaction
func (r *MenuRepository) UpdateStock(tx *gorm.DB, menuID uint, newStock int) error {
	return tx.Model(&model.Menu{}).Where("id = ?", menuID).Update("stock", newStock).Error
}

// Delete deletes a menu item by ID
func (r *MenuRepository) Delete(id uint) error {
	return r.db.Delete(&model.Menu{}, id).Error
}
