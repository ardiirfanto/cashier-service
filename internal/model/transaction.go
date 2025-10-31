package model

import (
	"time"
)

// Transaction represents a completed checkout transaction
type Transaction struct {
	ID          uint                `gorm:"primaryKey;autoIncrement" json:"id"`
	CashierID   uint                `gorm:"not null;index" json:"cashier_id"`
	TotalAmount float64             `gorm:"type:decimal(10,2);not null" json:"total_amount"`
	CreatedAt   time.Time           `gorm:"autoCreateTime" json:"created_at"`
	Details     []TransactionDetail `gorm:"foreignKey:TransactionID" json:"details,omitempty"`
	Cashier     User                `gorm:"foreignKey:CashierID" json:"cashier,omitempty"`
}

// TableName specifies the table name for the Transaction model
func (Transaction) TableName() string {
	return "transactions"
}

// TransactionDetail represents individual items in a transaction
type TransactionDetail struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	TransactionID uint      `gorm:"not null;index" json:"transaction_id"`
	MenuID        uint      `gorm:"not null;index" json:"menu_id"`
	Qty           int       `gorm:"not null" json:"qty"`
	Subtotal      float64   `gorm:"type:decimal(10,2);not null" json:"subtotal"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	Menu          Menu      `gorm:"foreignKey:MenuID" json:"menu,omitempty"`
}

// TableName specifies the table name for the TransactionDetail model
func (TransactionDetail) TableName() string {
	return "transaction_details"
}
