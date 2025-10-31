package service

import (
	"errors"
	"fmt"
	"service-cashier/internal/model"
	"service-cashier/internal/repository"

	"gorm.io/gorm"
)

// TransactionService handles transaction business logic
type TransactionService struct {
	transactionRepo *repository.TransactionRepository
	menuRepo        *repository.MenuRepository
}

// NewTransactionService creates a new TransactionService instance
func NewTransactionService(transactionRepo *repository.TransactionRepository, menuRepo *repository.MenuRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		menuRepo:        menuRepo,
	}
}

// CheckoutItem represents a single item in a checkout request
type CheckoutItem struct {
	MenuID uint `json:"menu_id" binding:"required"`
	Qty    int  `json:"qty" binding:"required,min=1"`
}

// CheckoutRequest represents the checkout request payload
type CheckoutRequest struct {
	Items []CheckoutItem `json:"items" binding:"required,min=1"`
}

// CheckoutResponse represents the checkout response payload
type CheckoutResponse struct {
	TransactionID uint                    `json:"transaction_id"`
	TotalAmount   float64                 `json:"total_amount"`
	Items         []CheckoutItemResponse  `json:"items"`
}

// CheckoutItemResponse represents a single item in the checkout response
type CheckoutItemResponse struct {
	MenuID   uint    `json:"menu_id"`
	Qty      int     `json:"qty"`
	Subtotal float64 `json:"subtotal"`
}

// ProcessedItem represents a processed checkout item from a goroutine
type ProcessedItem struct {
	MenuID   uint
	Qty      int
	Subtotal float64
	Menu     *model.Menu
	Error    error
}

// Checkout processes a checkout request sequentially within a database transaction
func (s *TransactionService) Checkout(cashierID uint, req *CheckoutRequest) (*CheckoutResponse, error) {
	// Start a database transaction
	tx := s.transactionRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Process each item sequentially
	var processedItems []ProcessedItem
	var totalAmount float64

	for _, item := range req.Items {
		// Process the item
		processedItem := s.processCheckoutItem(tx, item)

		// Check for errors
		if processedItem.Error != nil {
			tx.Rollback()
			return nil, processedItem.Error
		}

		processedItems = append(processedItems, processedItem)
		totalAmount += processedItem.Subtotal
	}

	// Create the transaction record
	transaction := &model.Transaction{
		CashierID:   cashierID,
		TotalAmount: totalAmount,
	}

	err := s.transactionRepo.Create(tx, transaction)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	// Create transaction details
	var details []model.TransactionDetail
	var responseItems []CheckoutItemResponse

	for _, item := range processedItems {
		detail := model.TransactionDetail{
			TransactionID: transaction.ID,
			MenuID:        item.MenuID,
			Qty:           item.Qty,
			Subtotal:      item.Subtotal,
		}
		details = append(details, detail)

		responseItems = append(responseItems, CheckoutItemResponse{
			MenuID:   item.MenuID,
			Qty:      item.Qty,
			Subtotal: item.Subtotal,
		})

		// Update stock for each item
		newStock := item.Menu.Stock - item.Qty
		err = s.menuRepo.UpdateStock(tx, item.MenuID, newStock)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("failed to update stock: %w", err)
		}
	}

	// Save all transaction details
	err = s.transactionRepo.CreateDetails(tx, details)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create transaction details: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Return successful response
	return &CheckoutResponse{
		TransactionID: transaction.ID,
		TotalAmount:   totalAmount,
		Items:         responseItems,
	}, nil
}

// processCheckoutItem processes a single checkout item
// This function is called within a goroutine for concurrent processing
func (s *TransactionService) processCheckoutItem(tx *gorm.DB, item CheckoutItem) ProcessedItem {
	// Fetch menu item with row-level lock to prevent race conditions
	menu, err := s.menuRepo.FindByIDWithLock(tx, item.MenuID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ProcessedItem{
				Error: fmt.Errorf("menu item with ID %d not found", item.MenuID),
			}
		}
		return ProcessedItem{
			Error: fmt.Errorf("failed to fetch menu item: %w", err),
		}
	}

	// Validate stock availability
	if menu.Stock < item.Qty {
		return ProcessedItem{
			Error: fmt.Errorf("insufficient stock for menu item '%s' (available: %d, requested: %d)",
				menu.Name, menu.Stock, item.Qty),
		}
	}

	// Calculate subtotal
	subtotal := menu.Price * float64(item.Qty)

	// Return processed item
	return ProcessedItem{
		MenuID:   item.MenuID,
		Qty:      item.Qty,
		Subtotal: subtotal,
		Menu:     menu,
		Error:    nil,
	}
}

// GetTransactionsByCashier retrieves all transactions for a specific cashier
func (s *TransactionService) GetTransactionsByCashier(cashierID uint) ([]model.Transaction, error) {
	return s.transactionRepo.GetByCashierID(cashierID)
}

// GetTransactionByID retrieves a transaction by ID
func (s *TransactionService) GetTransactionByID(id uint) (*model.Transaction, error) {
	return s.transactionRepo.FindByID(id)
}

// GetAllTransactions retrieves all transactions (admin function)
func (s *TransactionService) GetAllTransactions() ([]model.Transaction, error) {
	return s.transactionRepo.GetAll()
}
