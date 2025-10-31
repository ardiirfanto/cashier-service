# Project Structure

This document explains the organization and purpose of each directory and file in the Cashier API project.

## Directory Tree

```
service-cashier/
├── cmd/
│   └── server/
│       └── main.go                    # Application entry point
│
├── config/
│   └── config.go                      # Viper configuration loader
│
├── internal/
│   ├── database/
│   │   └── mysql.go                   # Database connection and migrations
│   │
│   ├── handler/
│   │   ├── auth_handler.go           # Authentication endpoints (login)
│   │   ├── menu_handler.go           # Menu endpoints (get menus)
│   │   └── transaction_handler.go    # Transaction endpoints (checkout, history)
│   │
│   ├── middleware/
│   │   └── jwt.go                    # JWT authentication middleware
│   │
│   ├── model/
│   │   ├── menu.go                   # Menu model (GORM)
│   │   ├── transaction.go            # Transaction & TransactionDetail models
│   │   └── user.go                   # User model
│   │
│   ├── repository/
│   │   ├── menu_repository.go        # Menu data access layer
│   │   ├── transaction_repository.go # Transaction data access layer
│   │   └── user_repository.go        # User data access layer
│   │
│   ├── router/
│   │   └── router.go                 # Route definitions and middleware setup
│   │
│   └── service/
│       ├── menu_service.go           # Menu business logic
│       ├── transaction_service.go    # Transaction logic with concurrency
│       └── user_service.go           # User authentication logic
│
├── pkg/
│   └── utils/
│       ├── jwt.go                    # JWT token generation and validation
│       └── response.go               # Standard API response helpers
│
├── .env.example                       # Environment variables template
├── .gitignore                         # Git ignore rules
├── API_TESTING.md                     # API testing guide with curl examples
├── database_setup.sql                 # Database schema and sample data
├── go.mod                             # Go module dependencies
├── Makefile                           # Common commands (run, build, test)
├── PROMPT_Cashier_API_Gin.md         # Original specification
├── PROJECT_STRUCTURE.md               # This file
├── QUICKSTART.md                      # Quick start guide
└── README.md                          # Main documentation
```

## Directory Purposes

### `/cmd`
Contains application entry points. Each subdirectory represents a different executable.
- **`/cmd/server`**: Main API server application

### `/config`
Configuration management using Viper to load environment variables.

### `/internal`
Private application code that cannot be imported by other projects.

#### `/internal/database`
Database connection setup and GORM auto-migrations.

#### `/internal/handler`
HTTP request handlers (controllers). Each handler:
- Parses request data
- Calls appropriate service methods
- Returns formatted JSON responses

#### `/internal/middleware`
HTTP middleware functions:
- **JWT middleware**: Validates tokens and extracts user information

#### `/internal/model`
GORM data models representing database tables:
- **User**: Cashier accounts
- **Menu**: Available items for purchase
- **Transaction**: Checkout transactions
- **TransactionDetail**: Individual items in a transaction

#### `/internal/repository`
Data access layer (Repository pattern):
- Abstracts database operations
- Provides clean interface for services
- Handles GORM queries

#### `/internal/router`
Route definitions and middleware registration:
- Public routes (login)
- Protected routes (menus, checkout, transactions)

#### `/internal/service`
Business logic layer:
- **UserService**: Authentication, password hashing
- **MenuService**: Menu management
- **TransactionService**: Checkout with **goroutines and channels**

### `/pkg`
Public packages that can be imported by other projects.

#### `/pkg/utils`
Utility functions:
- **JWT utilities**: Token generation and validation
- **Response helpers**: Standardized JSON responses

## Key Design Patterns

### Clean Architecture
The project follows clean architecture principles:
1. **Handler** (Presentation Layer) → receives HTTP requests
2. **Service** (Business Layer) → processes business logic
3. **Repository** (Data Layer) → manages data persistence
4. **Model** (Domain Layer) → defines data structures

### Dependency Injection
Dependencies are injected through constructors:
```go
userRepo := repository.NewUserRepository(db)
userService := service.NewUserService(userRepo, jwtSecret)
authHandler := handler.NewAuthHandler(userService)
```

### Concurrent Processing
The checkout process uses:
- **Goroutines**: Each item processed in parallel
- **Channels**: Collect results from goroutines
- **WaitGroups**: Synchronize completion
- **Database Transactions**: Ensure atomicity

## Data Flow Example: Checkout

```
1. Client sends POST /api/checkout with items
   ↓
2. JWT Middleware validates token, extracts cashier_id
   ↓
3. TransactionHandler.Checkout() receives request
   ↓
4. TransactionService.Checkout() processes items:
   - Starts database transaction
   - Launches goroutine for each item
   - Each goroutine validates stock and calculates subtotal
   - Results collected via channel
   - Creates transaction and details in database
   - Updates stock for each item
   - Commits transaction
   ↓
5. Returns checkout response with transaction details
```

## File Naming Conventions

- **Models**: Singular noun (e.g., `user.go`, `menu.go`)
- **Repositories**: `{model}_repository.go`
- **Services**: `{model}_service.go`
- **Handlers**: `{feature}_handler.go`
- **Utilities**: Descriptive name (e.g., `jwt.go`, `response.go`)

## Import Paths

All internal imports use the module name:
```go
import (
    "service-cashier/config"
    "service-cashier/internal/model"
    "service-cashier/pkg/utils"
)
```

## Configuration

Environment variables are loaded via Viper from:
1. `.env` file (if present)
2. System environment variables
3. Default values in `config.go`

## Database Migrations

GORM auto-migrations run on server startup:
- Creates tables if they don't exist
- Updates schema for model changes
- Does NOT delete columns or tables

For production, consider using migration tools like `golang-migrate`.

## Testing Structure (Future)

Recommended test organization:
```
internal/
├── handler/
│   ├── auth_handler.go
│   └── auth_handler_test.go
├── service/
│   ├── transaction_service.go
│   └── transaction_service_test.go
└── repository/
    ├── user_repository.go
    └── user_repository_test.go
```

## Adding New Features

To add a new feature (e.g., "Categories"):

1. **Create Model**: `internal/model/category.go`
2. **Create Repository**: `internal/repository/category_repository.go`
3. **Create Service**: `internal/service/category_service.go`
4. **Create Handler**: `internal/handler/category_handler.go`
5. **Register Routes**: Update `internal/router/router.go`
6. **Update Main**: Wire dependencies in `cmd/server/main.go`

## Performance Considerations

- **Goroutines**: Used in checkout for parallel item processing
- **Database Connection Pooling**: Handled by GORM
- **Row-level Locking**: Prevents race conditions during stock updates
- **Database Transactions**: Ensures data consistency

## Security Features

- **Password Hashing**: bcrypt with default cost
- **JWT Tokens**: HS256 signing, 24-hour expiration
- **Parameterized Queries**: GORM prevents SQL injection
- **Middleware Protection**: Sensitive endpoints require valid JWT

---

This structure provides a solid foundation for a scalable, maintainable API backend!
