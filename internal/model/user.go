package model

import (
	"time"
)

// User represents a cashier user in the system
type User struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username     string    `gorm:"type:varchar(50);uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"type:varchar(255);not null" json:"-"` // "-" prevents password from being serialized to JSON
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
