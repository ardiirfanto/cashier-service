package repository

import (
	"service-cashier/internal/model"

	"gorm.io/gorm"
)

// TransactionRepository handles transaction data access operations
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new TransactionRepository instance
func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create creates a new transaction within a database transaction
func (r *TransactionRepository) Create(tx *gorm.DB, transaction *model.Transaction) error {
	return tx.Create(transaction).Error
}

// CreateDetail creates a new transaction detail within a database transaction
func (r *TransactionRepository) CreateDetail(tx *gorm.DB, detail *model.TransactionDetail) error {
	return tx.Create(detail).Error
}

// CreateDetails creates multiple transaction details within a database transaction
func (r *TransactionRepository) CreateDetails(tx *gorm.DB, details []model.TransactionDetail) error {
	return tx.Create(&details).Error
}

// FindByID retrieves a transaction by ID with its details
func (r *TransactionRepository) FindByID(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.Preload("Details").Preload("Details.Menu").First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

// GetByCashierID retrieves all transactions for a specific cashier with details
func (r *TransactionRepository) GetByCashierID(cashierID uint) ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.
		Where("cashier_id = ?", cashierID).
		Preload("Details").
		Preload("Details.Menu").
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, err
}

// GetAll retrieves all transactions with details
func (r *TransactionRepository) GetAll() ([]model.Transaction, error) {
	var transactions []model.Transaction
	err := r.db.
		Preload("Details").
		Preload("Details.Menu").
		Preload("Cashier").
		Order("created_at DESC").
		Find(&transactions).Error
	return transactions, err
}

// BeginTransaction starts a new database transaction
func (r *TransactionRepository) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}
