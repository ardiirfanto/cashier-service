package repository

import (
	"service-cashier/internal/model"

	"gorm.io/gorm"
)

// UserRepository handles user data access operations
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// FindByUsername retrieves a user by username
func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID retrieves a user by ID
func (r *UserRepository) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Create creates a new user
func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// Update updates an existing user
func (r *UserRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

// Delete deletes a user by ID
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

// GetAll retrieves all users
func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}
