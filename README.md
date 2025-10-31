# Cashier API Backend

A RESTful API backend built with **Go (Golang)** and the **Gin** framework for a Cashier Application.

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)](https://gin-gonic.com)
[![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=flat&logo=mysql&logoColor=white)](https://www.mysql.com)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 🚀 Quick Start

```bash
# 1. Install dependencies
go mod download

# 2. Configure environment
cp .env.example .env
# Edit .env with your MySQL credentials

# 3. Setup database and run
mysql -u root -p < database_setup.sql
go run ./cmd/server
```

**Server will start on:** `http://localhost:8080`

**Test Credentials:**
- Username: `cashier1`
- Password: `password123`

## ✨ Features

- ✅ **JWT Authentication** - Secure token-based authentication
- ✅ **Concurrent Processing** - Goroutines & channels for high performance
- ✅ **Clean Architecture** - Scalable and maintainable code structure
- ✅ **RESTful API** - Industry-standard REST endpoints
- ✅ **MySQL Database** - Reliable data persistence
- ✅ **Auto-Migration** - Database schema auto-setup with GORM
- ✅ **Comprehensive Docs** - Complete documentation in `/docs` folder

## 📚 Documentation

All documentation is available in the [`/docs`](docs/) folder:

| Document | Description |
|----------|-------------|
| **[📖 INDEX](docs/INDEX.md)** | **Start here!** Complete navigation guide |
| **[⚡ QUICKSTART](docs/QUICKSTART.md)** | Get running in 5 minutes |
| **[📘 Full Documentation](docs/README.md)** | Complete API documentation |
| **[🧪 API Testing](docs/API_TESTING.md)** | Testing examples with curl |
| **[🏗️ Architecture](docs/ARCHITECTURE_DIAGRAM.md)** | Visual diagrams and flows |
| **[📁 Project Structure](docs/PROJECT_STRUCTURE.md)** | Code organization guide |
| **[🚀 Deployment](docs/DEPLOYMENT.md)** | Production deployment guide |
| **[✅ Implementation](docs/IMPLEMENTATION_SUMMARY.md)** | What was built |
| **[📋 Manifest](docs/PROJECT_MANIFEST.txt)** | Complete file checklist |

## 🎯 API Endpoints

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| `POST` | `/api/login` | ❌ | Login and get JWT token |
| `GET` | `/api/menus` | ✅ | Get all menu items |
| `POST` | `/api/checkout` | ✅ | Process checkout (concurrent) |
| `GET` | `/api/transactions` | ✅ | Get transaction history |
| `GET` | `/health` | ❌ | Health check |

**Full API examples:** [docs/API_TESTING.md](docs/API_TESTING.md)

## 🛠️ Tech Stack

- **[Go 1.21+](https://golang.org)** - Programming language
- **[Gin](https://gin-gonic.com)** - HTTP web framework
- **[GORM](https://gorm.io)** - ORM for MySQL
- **[MySQL 8.0](https://www.mysql.com)** - Database
- **[JWT](https://github.com/golang-jwt/jwt)** - Authentication
- **[Viper](https://github.com/spf13/viper)** - Configuration
- **[bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)** - Password hashing

## 📁 Project Structure

```
service-cashier/
├── cmd/server/          # Application entry point
├── config/              # Configuration management
├── internal/            # Private application code
│   ├── database/        # Database connection
│   ├── handler/         # HTTP handlers
│   ├── middleware/      # JWT middleware
│   ├── model/           # Data models
│   ├── repository/      # Data access layer
│   ├── router/          # Route definitions
│   └── service/         # Business logic (with concurrency)
├── pkg/utils/           # Utility packages
├── docs/                # 📚 All documentation
├── database_setup.sql   # Database schema + sample data
├── Makefile             # Common commands
└── go.mod               # Dependencies
```

**Detailed structure:** [docs/PROJECT_STRUCTURE.md](docs/PROJECT_STRUCTURE.md)

## 💡 Key Features

### 🔐 Secure Authentication
- bcrypt password hashing
- JWT tokens with 24-hour expiration
- Protected endpoints with middleware

### ⚡ Concurrent Checkout
The checkout process uses **goroutines and channels** for high-performance parallel processing:
- Each item processed concurrently
- Results collected via buffered channels
- WaitGroup synchronization
- Row-level database locking
- Transaction atomicity

**See implementation:** [internal/service/transaction_service.go](internal/service/transaction_service.go#L45-L134)

### 🏗️ Clean Architecture
- Handler → Service → Repository → Model
- Dependency injection
- Testable and maintainable
- Separation of concerns

## 🧪 Testing

```bash
# Health check
curl http://localhost:8080/health

# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "cashier1", "password": "password123"}'

# Get menus (with token)
curl http://localhost:8080/api/menus \
  -H "Authorization: Bearer YOUR_TOKEN"
```

**More examples:** [docs/API_TESTING.md](docs/API_TESTING.md)

## 🔧 Common Commands

```bash
make install        # Install dependencies
make run           # Run development server
make build         # Build production binary
make test          # Run tests
make clean         # Clean build artifacts
```

**See all commands:** `make help`

## 🚀 Deployment

The project is ready for deployment to:
- Traditional servers (Ubuntu/CentOS)
- Docker containers
- Cloud platforms (AWS, GCP, Azure, Heroku)
- Kubernetes

**Full deployment guide:** [docs/DEPLOYMENT.md](docs/DEPLOYMENT.md)

## 📊 Database Schema

- **users** - Cashier accounts with bcrypt passwords
- **menus** - Available items with stock tracking
- **transactions** - Checkout records
- **transaction_details** - Individual items per transaction

All tables include `created_at` timestamp.

**Setup database:**
```bash
mysql -u root -p < database_setup.sql
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 📝 License

This project is licensed under the MIT License.

## 💬 Support

For help and questions:
- 📖 Read the [documentation](docs/)
- 🔍 Check [API examples](docs/API_TESTING.md)
- 🚀 Follow [Quick Start](docs/QUICKSTART.md)

## 🎓 Learning Resources

This project demonstrates:
- ✅ Goroutines and Channels (concurrent processing)
- ✅ Clean Architecture pattern
- ✅ RESTful API design
- ✅ JWT authentication
- ✅ Database design with relationships
- ✅ Error handling and validation
- ✅ Configuration management

**Perfect for learning Go backend development!**

## 📈 Project Status

✅ **Production Ready**

- All features implemented
- Comprehensive documentation
- Security best practices
- Clean, maintainable code
- Ready for immediate use

---

**Built with ❤️ using Go and Gin Framework**

**[📚 View Full Documentation →](docs/INDEX.md)**
