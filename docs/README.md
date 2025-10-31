# Cashier API Backend

A RESTful API backend built with **Go (Golang)** and the **Gin** framework for a Cashier Application. This backend supports both React/Next.js web frontend and Flutter mobile client.

## Features

- **Authentication**: JWT-based authentication for cashiers
- **Menu Management**: List all available menu items
- **Checkout Process**: Concurrent transaction processing using goroutines and channels
- **Transaction History**: View transaction history per cashier
- **High Performance**: Built with concurrency in mind for handling multiple simultaneous checkouts

## Tech Stack

- **Gin** - HTTP web framework
- **GORM** - ORM for MySQL
- **MySQL** - Primary database
- **JWT** - Authentication middleware
- **Viper** - Environment configuration
- **Goroutines & Channels** - Concurrent transaction processing
- **Go Modules** - Dependency management

## Prerequisites

- Go 1.21 or higher
- MySQL 5.7 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd service-cashier
```

2. Install dependencies:
```bash
go mod download
```

3. Create a `.env` file from the example:
```bash
cp .env.example .env
```

4. Update the `.env` file with your database credentials and JWT secret:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=cashier_db
JWT_SECRET=supersecretkey
SERVER_PORT=8080
```

5. Create the MySQL database:
```sql
CREATE DATABASE cashier_db;
```

6. Run the database schema (or let GORM auto-migrate):
```sql
USE cashier_db;

CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE menus (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  stock INT DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
  id INT AUTO_INCREMENT PRIMARY KEY,
  cashier_id INT NOT NULL,
  total_amount DECIMAL(10,2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (cashier_id) REFERENCES users(id)
);

CREATE TABLE transaction_details (
  id INT AUTO_INCREMENT PRIMARY KEY,
  transaction_id INT NOT NULL,
  menu_id INT NOT NULL,
  qty INT NOT NULL,
  subtotal DECIMAL(10,2) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  FOREIGN KEY (menu_id) REFERENCES menus(id)
);
```

7. Insert sample data (optional):
```sql
-- Insert a test user (password: "password123")
INSERT INTO users (username, password_hash) VALUES
('cashier1', '$2a$10$YourBcryptHashHere');

-- Insert sample menu items
INSERT INTO menus (name, price, stock) VALUES
('Coffee', 25000.00, 100),
('Tea', 15000.00, 100),
('Sandwich', 35000.00, 50),
('Burger', 45000.00, 50),
('Juice', 20000.00, 75);
```

## Running the Application

Start the server:
```bash
go run ./cmd/server
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Authentication

#### Login
```http
POST /api/login
Content-Type: application/json

{
  "username": "cashier1",
  "password": "password123"
}
```

Response:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIs..."
  }
}
```

### Menu (Protected - Requires JWT)

#### Get All Menus
```http
GET /api/menus
Authorization: Bearer <your-jwt-token>
```

Response:
```json
{
  "success": true,
  "message": "Menus retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Espresso",
      "price": 25000.00,
      "stock": 100,
      "image": "https://images.unsplash.com/photo-1510591509098-f4fdc6d0ff04?w=400",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

### Transactions (Protected - Requires JWT)

#### Checkout
```http
POST /api/checkout
Authorization: Bearer <your-jwt-token>
Content-Type: application/json

{
  "items": [
    {
      "menu_id": 1,
      "qty": 2
    },
    {
      "menu_id": 3,
      "qty": 1
    }
  ]
}
```

Response:
```json
{
  "success": true,
  "message": "Checkout successful",
  "data": {
    "transaction_id": 1,
    "total_amount": 85000.00,
    "items": [
      {
        "menu_id": 1,
        "qty": 2,
        "subtotal": 50000.00
      },
      {
        "menu_id": 3,
        "qty": 1,
        "subtotal": 35000.00
      }
    ]
  }
}
```

#### Get Transaction History
```http
GET /api/transactions
Authorization: Bearer <your-jwt-token>
```

Response:
```json
{
  "success": true,
  "message": "Transactions retrieved successfully",
  "data": [
    {
      "id": 1,
      "cashier_id": 1,
      "total_amount": 85000.00,
      "created_at": "2024-01-01T12:00:00Z",
      "details": [
        {
          "id": 1,
          "transaction_id": 1,
          "menu_id": 1,
          "qty": 2,
          "subtotal": 50000.00,
          "created_at": "2024-01-01T12:00:00Z"
        }
      ]
    }
  ]
}
```

## Project Structure

```
/cmd
  /server
    main.go                    # Application entry point
/config
  config.go                    # Viper configuration setup
/internal
  /database
    mysql.go                   # Database connection
  /middleware
    jwt.go                     # JWT authentication middleware
  /model
    user.go                    # User model
    menu.go                    # Menu model
    transaction.go             # Transaction models
  /repository
    user_repository.go         # User data access layer
    menu_repository.go         # Menu data access layer
    transaction_repository.go  # Transaction data access layer
  /service
    user_service.go            # User business logic
    menu_service.go            # Menu business logic
    transaction_service.go     # Transaction business logic (with concurrency)
  /handler
    auth_handler.go            # Authentication endpoints
    menu_handler.go            # Menu endpoints
    transaction_handler.go     # Transaction endpoints
  /router
    router.go                  # Route definitions
/pkg
  /utils
    response.go                # JSON response helpers
    jwt.go                     # JWT utilities
.env.example                   # Environment variables template
go.mod                         # Go module dependencies
README.md                      # This file
```

## Concurrency Design

The checkout process uses **goroutines** and **channels** for high-performance concurrent processing:

1. Each item in a checkout request is processed in a separate goroutine
2. Each goroutine validates stock availability and calculates the subtotal
3. Results are collected through a channel
4. A `sync.WaitGroup` ensures all goroutines complete before proceeding
5. Once all items are processed, the transaction is committed to the database

This design allows the API to handle multiple items in a single checkout simultaneously, reducing processing time and improving throughput under heavy load.

## Error Handling

All responses follow a consistent JSON format:

Success:
```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

Error:
```json
{
  "success": false,
  "message": "Error description",
  "data": null
}
```

## Security

- Passwords are hashed using **bcrypt** before storage
- JWT tokens are required for all protected endpoints
- JWT tokens contain cashier ID for request authorization
- Database queries use GORM's parameterized queries to prevent SQL injection

## Development

To add new features:

1. Define the model in `/internal/model`
2. Create repository methods in `/internal/repository`
3. Implement business logic in `/internal/service`
4. Create handlers in `/internal/handler`
5. Register routes in `/internal/router`

## Testing

To create a test user with password "password123":

```bash
# You can use this bcrypt hash in your INSERT statement:
# $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy
```

## License

MIT License

## Support

For issues and questions, please create an issue in the repository.
