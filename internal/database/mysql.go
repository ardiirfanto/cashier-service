package database

import (
	"fmt"
	"log"
	"service-cashier/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is the global database connection instance
var DB *gorm.DB

// Connect initializes the database connection using GORM
func Connect(dsn string) error {
	var err error

	// Configure GORM with logger
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully")

	// Run auto-migrations for all models
	if err := AutoMigrate(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	return nil
}

// AutoMigrate runs GORM auto-migration for all models
func AutoMigrate() error {
	log.Println("Running database migrations...")

	err := DB.AutoMigrate(
		&model.User{},
		&model.Menu{},
		&model.Transaction{},
		&model.TransactionDetail{},
	)

	if err != nil {
		return err
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// Close closes the database connection
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
