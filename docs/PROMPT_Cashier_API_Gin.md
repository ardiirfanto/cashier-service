# 🧩 Prompt for Claude Code: Cashier API Backend (Go + Gin)

## 🧠 Overview

Create a **RESTful API backend** using **Go (Golang)** and the **Gin** framework for a **Cashier Application**.  
This backend will support a **React/Next.js** web frontend and a **Flutter** mobile client.

The backend must handle:
- Authentication for cashiers.
- Menu listing and checkout process.
- Transaction history per cashier.

It should be designed with performance and concurrency in mind using **goroutines and channels**.

---

## 🔧 Tech Stack

- **Gin** → HTTP web framework  
- **GORM** → ORM for MySQL  
- **MySQL** → Primary database  
- **JWT** → Authentication middleware  
- **Viper** → Environment configuration  
- **Goroutines & Channels** → Handle concurrent transaction processing  
- **Go Modules** → Dependency management

---

## 📁 Folder Structure

Use a clean, beginner-friendly folder layout that’s still scalable for future growth:

```
/cmd
  /server
    main.go
/config
  config.go
/internal
  /database
    mysql.go
  /middleware
    jwt.go
  /model
    user.go
    menu.go
    transaction.go
  /repository
    user_repository.go
    menu_repository.go
    transaction_repository.go
  /service
    user_service.go
    menu_service.go
    transaction_service.go
  /handler
    auth_handler.go
    menu_handler.go
    transaction_handler.go
  /router
    router.go
/pkg
  /utils
    response.go
    jwt.go
.env.example
go.mod
go.sum
README.md
```

---

## 🧩 Feature Specification

### 1. **Cashier Login**
- **Endpoint:** `POST /api/login`
- **Input:** `username`, `password`
- **Output:** JWT token
- Passwords must be hashed using **bcrypt**.
- Authentication logic handled inside `user_service.go`.

---

### 2. **Menu List**
- **Endpoint:** `GET /api/menus`
- Requires a valid JWT token.
- Returns all menu items from the `menus` table.

---

### 3. **Checkout**
- **Endpoint:** `POST /api/checkout`
- **Input:** array of items (`menu_id`, `qty`)
- Each item is processed concurrently using **goroutines**.
- Use **channels** to gather results and a **WaitGroup** to synchronize completion.
- Write results into `transactions` and `transaction_details` tables once all processing is done.

---

### 4. **Transaction History**
- **Endpoint:** `GET /api/transactions`
- Requires JWT authentication.
- Returns transactions filtered by cashier (based on JWT claims).

---

## ⚙️ Concurrency Design

Use **goroutines** and **channels** during checkout for parallel processing:

**Flow Example:**
1. Each item in a checkout request starts a goroutine.
2. Each goroutine validates stock and calculates subtotal.
3. Each result is sent back via a channel.
4. A `sync.WaitGroup` ensures all goroutines complete.
5. Once all are done, aggregate totals and commit to MySQL.

This structure helps achieve high concurrency and low response times even under heavy load.

---

## 🔐 Middleware

JWT Middleware must:
- Validate token.
- Parse claims.
- Attach `cashier_id` to the Gin context.
- Restrict access to protected endpoints.

---

## ⚙️ Configuration

Use **Viper** to read `.env` configuration.

Example `.env.example`:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=cashier_db
JWT_SECRET=supersecretkey
SERVER_PORT=8080
```

---

## 🧰 Database Schema (MySQL DDL)

Use the following schema when setting up the MySQL database.

```sql
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

---

## 🚀 Running the Project

After generation, you should be able to start the server with:

```bash
go run ./cmd/server
```

The application should:
1. Load environment variables with Viper.
2. Connect to the MySQL database.
3. Run GORM auto-migrations.
4. Serve all routes at the configured port (default: `8080`).

---

## 🎯 Final Deliverables

Generate a **complete working Go project** that includes:
- Full CRUD for users, menus, and transactions.
- JWT authentication middleware.
- Concurrency in checkout logic with goroutines and channels.
- Clear code separation following the provided folder structure.
- `.env` configuration via Viper.
- Proper error handling and clean responses (JSON format).

The result should be a **fully runnable backend** that can serve as the API foundation for a React/Next.js web client and a Flutter mobile client.

---
