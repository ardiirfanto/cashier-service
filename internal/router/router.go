package router

import (
	"service-cashier/internal/handler"
	"service-cashier/internal/middleware"

	"github.com/gin-gonic/gin"
)

// RouterConfig holds the configuration needed to set up routes
type RouterConfig struct {
	AuthHandler        *handler.AuthHandler
	MenuHandler        *handler.MenuHandler
	TransactionHandler *handler.TransactionHandler
	JWTSecret          string
}

// SetupRouter configures and returns the Gin router with all routes
func SetupRouter(config *RouterConfig) *gin.Engine {
	// Create a new Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// API group
	api := router.Group("/api")
	{
		// Public routes (no authentication required)
		api.POST("/login", config.AuthHandler.Login)

		// Protected routes (require JWT authentication)
		protected := api.Group("")
		protected.Use(middleware.JWTAuth(config.JWTSecret))
		{
			// Menu routes
			protected.GET("/menus", config.MenuHandler.GetMenus)

			// Transaction routes
			protected.POST("/checkout", config.TransactionHandler.Checkout)
			protected.GET("/transactions", config.TransactionHandler.GetTransactions)
		}
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Cashier API is running",
		})
	})

	return router
}
