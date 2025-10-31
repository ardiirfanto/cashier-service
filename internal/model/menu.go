package model

import (
	"time"
)

// Menu represents a menu item available for purchase
type Menu struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Price     float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock     int       `gorm:"type:int;default:0" json:"stock"`
	Image     string    `gorm:"type:varchar(255)" json:"image"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// TableName specifies the table name for the Menu model
func (Menu) TableName() string {
	return "menus"
}
