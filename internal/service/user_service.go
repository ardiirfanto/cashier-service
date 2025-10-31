package service

import (
	"errors"
	"service-cashier/internal/model"
	"service-cashier/internal/repository"
	"service-cashier/pkg/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserService handles user business logic
type UserService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo *repository.UserRepository, jwtSecret string) *UserService {
	return &UserService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

// LoginRequest represents the login request payload
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse represents the login response payload
type LoginResponse struct {
	Token string `json:"token"`
}

// Login authenticates a user and returns a JWT token
func (s *UserService) Login(req *LoginRequest) (*LoginResponse, error) {
	// Find user by username
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID, user.Username, s.jwtSecret)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &LoginResponse{Token: token}, nil
}

// CreateUser creates a new user with hashed password
func (s *UserService) CreateUser(username, password string) (*model.User, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByUsername(username)
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := &model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID retrieves a user by ID
func (s *UserService) GetUserByID(id uint) (*model.User, error) {
	return s.userRepo.FindByID(id)
}
