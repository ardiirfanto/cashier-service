# Quick Start Guide

Get the Cashier API up and running in 5 minutes!

## Prerequisites

- Go 1.21 or higher installed
- MySQL 5.7+ running locally
- Git (optional)

## Step 1: Install Dependencies

```bash
go mod download
```

Or using the Makefile:
```bash
make install
```

## Step 2: Configure Environment

Create a `.env` file:
```bash
cp .env.example .env
```

Edit `.env` with your MySQL credentials:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=cashier_db
JWT_SECRET=supersecretkey
SERVER_PORT=8080
```

## Step 3: Setup Database

Run the database setup script:
```bash
mysql -u root -p < database_setup.sql
```

This will:
- Create the `cashier_db` database
- Create all required tables
- Insert sample users (username: `cashier1`, `cashier2`, `admin` - password: `password123`)
- Insert 15 sample menu items

## Step 4: Run the Server

```bash
go run ./cmd/server
```

Or using the Makefile:
```bash
make run
```

You should see:
```
Configuration loaded successfully
Database connection established successfully
Running database migrations...
Database migrations completed successfully
Starting server on :8080
```

## Step 5: Test the API

### Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "cashier1",
    "password": "password123"
  }'
```

Save the token from the response!

### Get Menus
```bash
curl -X GET http://localhost:8080/api/menus \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### Checkout
```bash
curl -X POST http://localhost:8080/api/checkout \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -d '{
    "items": [
      {"menu_id": 1, "qty": 2},
      {"menu_id": 3, "qty": 1}
    ]
  }'
```

### Get Transaction History
```bash
curl -X GET http://localhost:8080/api/transactions \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## Next Steps

- Read [README.md](README.md) for detailed documentation
- Check [API_TESTING.md](API_TESTING.md) for comprehensive testing examples
- Explore the code structure in `/internal` and `/pkg` directories

## Troubleshooting

### "Failed to connect to database"
- Verify MySQL is running: `mysql -u root -p`
- Check your `.env` credentials
- Ensure `cashier_db` database exists

### "Failed to load configuration"
- Make sure `.env` file exists in the project root
- Verify all required environment variables are set

### "Invalid token"
- Token expires after 24 hours
- Login again to get a fresh token

### Port already in use
- Change `SERVER_PORT` in `.env` to a different port (e.g., 8081)

## Building for Production

```bash
# Build binary
make build

# Run binary
./bin/cashier-api
```

## Docker (Optional)

Build and run with Docker:
```bash
# Build image
docker build -t cashier-api .

# Run container
docker run -p 8080:8080 --env-file .env cashier-api
```

---

**Happy Coding! 🚀**
