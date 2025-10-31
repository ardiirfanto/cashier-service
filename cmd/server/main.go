package main

import (
	"fmt"
	"log"
	"service-cashier/config"
	"service-cashier/internal/database"
	"service-cashier/internal/handler"
	"service-cashier/internal/repository"
	"service-cashier/internal/router"
	"service-cashier/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Println("Configuration loaded successfully")

	// Connect to database
	err = database.Connect(cfg.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	// Get database instance
	db := database.GetDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	menuRepo := repository.NewMenuRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo, cfg.JWT.Secret)
	menuService := service.NewMenuService(menuRepo)
	transactionService := service.NewTransactionService(transactionRepo, menuRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(userService)
	menuHandler := handler.NewMenuHandler(menuService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Setup router with all handlers
	r := router.SetupRouter(&router.RouterConfig{
		AuthHandler:        authHandler,
		MenuHandler:        menuHandler,
		TransactionHandler: transactionHandler,
		JWTSecret:          cfg.JWT.Secret,
	})

	// Start server
	serverAddr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("Starting server on %s", serverAddr)

	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
