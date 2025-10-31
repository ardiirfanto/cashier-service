# API Testing Guide

This document provides examples for testing all API endpoints using `curl`.

## Setup

1. Make sure the server is running on `http://localhost:8080`
2. Create a `.env` file based on `.env.example`
3. Run the database setup script: `mysql -u root -p < database_setup.sql`
4. Start the server: `go run ./cmd/server`

## Test Credentials

Default test users (password: `password123`):
- `cashier1`
- `cashier2`
- `admin`

---

## 1. Health Check

```bash
curl -X GET http://localhost:8080/health
```

Expected Response:
```json
{
  "message": "Cashier API is running",
  "status": "ok"
}
```

---

## 2. Login

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "cashier1",
    "password": "password123"
  }'
```

Expected Response:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Save the token for subsequent requests!**

---

## 3. Get All Menus (Protected)

Replace `YOUR_JWT_TOKEN` with the token from the login response.

```bash
curl -X GET http://localhost:8080/api/menus \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

Expected Response:
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
    },
    {
      "id": 2,
      "name": "Cappuccino",
      "price": 30000.00,
      "stock": 100,
      "image": "https://images.unsplash.com/photo-1572442388796-11668a67e53d?w=400",
      "created_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

---

## 4. Checkout (Protected)

Replace `YOUR_JWT_TOKEN` with the token from the login response.

```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "items": [
      {
        "menu_id": 1,
        "qty": 2
      },
      {
        "menu_id": 3,
        "qty": 1
      },
      {
        "menu_id": 7,
        "qty": 3
      }
    ]
  }'
```

Expected Response:
```json
{
  "success": true,
  "message": "Checkout successful",
  "data": {
    "transaction_id": 1,
    "total_amount": 187000.00,
    "items": [
      {
        "menu_id": 1,
        "qty": 2,
        "subtotal": 50000.00
      },
      {
        "menu_id": 3,
        "qty": 1,
        "subtotal": 32000.00
      },
      {
        "menu_id": 7,
        "qty": 3,
        "subtotal": 105000.00
      }
    ]
  }
}
```

---

## 5. Get Transaction History (Protected)

Replace `YOUR_JWT_TOKEN` with the token from the login response.

```bash
curl -X GET http://localhost:8080/api/transactions \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

Expected Response:
```json
{
  "success": true,
  "message": "Transactions retrieved successfully",
  "data": [
    {
      "id": 1,
      "cashier_id": 1,
      "total_amount": 187000.00,
      "created_at": "2024-01-01T12:00:00Z",
      "details": [
        {
          "id": 1,
          "transaction_id": 1,
          "menu_id": 1,
          "qty": 2,
          "subtotal": 50000.00,
          "created_at": "2024-01-01T12:00:00Z",
          "menu": {
            "id": 1,
            "name": "Espresso",
            "price": 25000.00,
            "stock": 98,
            "created_at": "2024-01-01T00:00:00Z"
          }
        },
        ...
      ]
    }
  ]
}
```

---

## Error Scenarios

### Invalid Login Credentials
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "cashier1",
    "password": "wrongpassword"
  }'
```

Response:
```json
{
  "success": false,
  "message": "invalid username or password",
  "data": null
}
```

### Missing Authorization Header
```bash
curl -X GET http://localhost:8080/api/menus
```

Response:
```json
{
  "success": false,
  "message": "Authorization header is required",
  "data": null
}
```

### Insufficient Stock
```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "items": [
      {
        "menu_id": 1,
        "qty": 1000
      }
    ]
  }'
```

Response:
```json
{
  "success": false,
  "message": "insufficient stock for menu item 'Espresso' (available: 100, requested: 1000)",
  "data": null
}
```

### Invalid Menu ID
```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "items": [
      {
        "menu_id": 999,
        "qty": 1
      }
    ]
  }'
```

Response:
```json
{
  "success": false,
  "message": "menu item with ID 999 not found",
  "data": null
}
```

---

## Postman Collection

You can also import these requests into Postman:

1. Create a new collection
2. Add the login request and use the "Tests" tab to save the token:
```javascript
var jsonData = JSON.parse(responseBody);
pm.environment.set("jwt_token", jsonData.data.token);
```
3. For protected endpoints, add the header:
   - Key: `Authorization`
   - Value: `Bearer {{jwt_token}}`

---

## Testing Concurrency

To test the concurrent checkout processing, you can use a tool like Apache Bench or create multiple simultaneous requests:

```bash
# Install Apache Bench (if not already installed)
# macOS: comes pre-installed
# Ubuntu: sudo apt-get install apache2-utils

# Run 100 requests with 10 concurrent connections
ab -n 100 -c 10 -T 'application/json' \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -p checkout.json \
  http://localhost:8080/api/checkout
```

Create a `checkout.json` file:
```json
{
  "items": [
    {
      "menu_id": 1,
      "qty": 1
    }
  ]
}
```

This will test the concurrent processing capability of the checkout endpoint using goroutines and channels.
