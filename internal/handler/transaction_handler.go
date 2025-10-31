package handler

import (
	"service-cashier/internal/middleware"
	"service-cashier/internal/service"
	"service-cashier/pkg/utils"

	"github.com/gin-gonic/gin"
)

// TransactionHandler handles transaction-related HTTP requests
type TransactionHandler struct {
	transactionService *service.TransactionService
}

// NewTransactionHandler creates a new TransactionHandler instance
func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

// Checkout handles the checkout endpoint
// POST /api/checkout
func (h *TransactionHandler) Checkout(c *gin.Context) {
	var req service.CheckoutRequest

	// Bind JSON request body
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequestResponse(c, "Invalid request payload")
		return
	}

	// Get cashier ID from JWT middleware context
	cashierID, ok := middleware.GetUserID(c)
	if !ok {
		utils.UnauthorizedResponse(c, "Unable to retrieve user information")
		return
	}

	// Process checkout with concurrent item processing
	response, err := h.transactionService.Checkout(cashierID, &req)
	if err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// Return success response
	utils.SuccessResponse(c, "Checkout successful", response)
}

// GetTransactions handles the get transaction history endpoint
// GET /api/transactions
func (h *TransactionHandler) GetTransactions(c *gin.Context) {
	// Get cashier ID from JWT middleware context
	cashierID, ok := middleware.GetUserID(c)
	if !ok {
		utils.UnauthorizedResponse(c, "Unable to retrieve user information")
		return
	}

	// Retrieve transactions for the cashier
	transactions, err := h.transactionService.GetTransactionsByCashier(cashierID)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to retrieve transactions")
		return
	}

	// Return success response
	utils.SuccessResponse(c, "Transactions retrieved successfully", transactions)
}
