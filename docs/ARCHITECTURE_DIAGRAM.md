# Architecture Diagram

Visual representation of the Cashier API architecture and data flow.

## System Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLIENT LAYER                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │ React/Next.js│  │   Flutter    │  │   Postman    │          │
│  │  Web Client  │  │ Mobile Client│  │  API Testing │          │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘          │
└─────────┼──────────────────┼──────────────────┼─────────────────┘
          │                  │                  │
          └──────────────────┴──────────────────┘
                             │
                    HTTP/HTTPS (JSON)
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                     API SERVER (Gin)                             │
│                    Port: 8080 (default)                          │
└─────────────────────────────────────────────────────────────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      ROUTER LAYER                                │
│  ┌───────────────────────────────────────────────────────┐      │
│  │  Routes:                                              │      │
│  │  • POST /api/login           (Public)                 │      │
│  │  • GET  /api/menus           (Protected - JWT)        │      │
│  │  • POST /api/checkout        (Protected - JWT)        │      │
│  │  • GET  /api/transactions    (Protected - JWT)        │      │
│  │  • GET  /health              (Public)                 │      │
│  └───────────────────────────────────────────────────────┘      │
└─────────────────────────────────────────────────────────────────┘
                             │
                    ┌────────┴────────┐
                    │                 │
                    ▼                 ▼
          ┌─────────────────┐  ┌──────────────┐
          │   JWT Middleware │  │   Public     │
          │   (Protected)    │  │   Routes     │
          └────────┬─────────┘  └──────┬───────┘
                   │                   │
                   └─────────┬─────────┘
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      HANDLER LAYER                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │    Auth      │  │     Menu     │  │ Transaction  │          │
│  │   Handler    │  │   Handler    │  │   Handler    │          │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘          │
└─────────┼──────────────────┼──────────────────┼─────────────────┘
          │                  │                  │
          ▼                  ▼                  ▼
┌─────────────────────────────────────────────────────────────────┐
│                      SERVICE LAYER                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │    User      │  │     Menu     │  │ Transaction  │          │
│  │   Service    │  │   Service    │  │   Service    │          │
│  │              │  │              │  │              │          │
│  │ • Login      │  │ • GetAll     │  │ • Checkout   │          │
│  │ • Hash PWD   │  │ • GetByID    │  │ • GetHistory │          │
│  │ • Gen JWT    │  │              │  │ • Goroutines │          │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘          │
└─────────┼──────────────────┼──────────────────┼─────────────────┘
          │                  │                  │
          ▼                  ▼                  ▼
┌─────────────────────────────────────────────────────────────────┐
│                    REPOSITORY LAYER                              │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐          │
│  │    User      │  │     Menu     │  │ Transaction  │          │
│  │  Repository  │  │  Repository  │  │  Repository  │          │
│  │              │  │              │  │              │          │
│  │ • FindByUser │  │ • GetAll     │  │ • Create     │          │
│  │ • FindByID   │  │ • FindByID   │  │ • GetByCash  │          │
│  │ • Create     │  │ • LockForUpd │  │ • BeginTx    │          │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘          │
└─────────┼──────────────────┼──────────────────┼─────────────────┘
          │                  │                  │
          └──────────────────┴──────────────────┘
                             │
                    GORM (ORM Layer)
                             │
                             ▼
┌─────────────────────────────────────────────────────────────────┐
│                      DATABASE LAYER                              │
│                         MySQL 8.0                                │
│  ┌────────────┐  ┌─────────┐  ┌──────────────┐  ┌────────────┐ │
│  │   users    │  │  menus  │  │ transactions │  │ trans_det  │ │
│  ├────────────┤  ├─────────┤  ├──────────────┤  ├────────────┤ │
│  │ id         │  │ id      │  │ id           │  │ id         │ │
│  │ username   │  │ name    │  │ cashier_id ─┼──┤ trans_id   │ │
│  │ pass_hash  │  │ price   │  │ total_amount │  │ menu_id  ──┼─┐│
│  │ created_at │  │ stock   │  │ created_at   │  │ qty        │ ││
│  └────────────┘  │created  │  └──────────────┘  │ subtotal   │ ││
│                  └─────────┘                     │ created_at │ ││
│                       │                          └────────────┘ ││
│                       └───────────────────────────────────────┘│
└─────────────────────────────────────────────────────────────────┘
```

## Checkout Flow with Concurrency

```
┌────────────────────────────────────────────────────────────────────┐
│                    CONCURRENT CHECKOUT FLOW                         │
└────────────────────────────────────────────────────────────────────┘

Client Request
     │
     ▼
POST /api/checkout
{ items: [{menu_id:1, qty:2}, {menu_id:3, qty:1}, {menu_id:7, qty:3}] }
     │
     ▼
┌─────────────────────────────────────────────────────────────────┐
│                  Transaction Service                             │
│  1. Begin Database Transaction                                  │
│  2. Create Result Channel (buffered)                            │
│  3. Create WaitGroup                                            │
└─────────────────────────────────────────────────────────────────┘
     │
     ├───────────────────┬───────────────────┬────────────────────┐
     ▼                   ▼                   ▼                    ▼
┌──────────┐      ┌──────────┐      ┌──────────┐         ┌──────────┐
│Goroutine1│      │Goroutine2│      │Goroutine3│   ...   │GoroutineN│
│          │      │          │      │          │         │          │
│ Item 1   │      │ Item 2   │      │ Item 3   │         │ Item N   │
│menu_id:1 │      │menu_id:3 │      │menu_id:7 │         │          │
│qty: 2    │      │qty: 1    │      │qty: 3    │         │          │
└────┬─────┘      └────┬─────┘      └────┬─────┘         └────┬─────┘
     │                 │                 │                     │
     │  1. Lock Row    │  1. Lock Row    │  1. Lock Row        │
     │  2. Check Stock │  2. Check Stock │  2. Check Stock     │
     │  3. Calc Price  │  3. Calc Price  │  3. Calc Price      │
     │                 │                 │                     │
     ├─────────────────┴─────────────────┴─────────────────────┤
     │                                                          │
     ▼                    Channel (buffered)                    │
┌─────────────────────────────────────────────────────────────────┐
│  Result Channel                                                  │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐      ┌──────────┐   │
│  │ Result 1 │  │ Result 2 │  │ Result 3 │ ...  │ Result N │   │
│  └──────────┘  └──────────┘  └──────────┘      └──────────┘   │
└─────────────────────────────────────────────────────────────────┘
     │
     │   WaitGroup.Wait() - All goroutines complete
     │   Channel closed
     ▼
┌─────────────────────────────────────────────────────────────────┐
│  Aggregate Results                                               │
│  • Check for errors                                              │
│  • Calculate total amount                                        │
│  • Create transaction record                                     │
│  • Create transaction details                                    │
│  • Update stock for each item                                    │
│  • Commit database transaction                                   │
└─────────────────────────────────────────────────────────────────┘
     │
     ▼
JSON Response
{
  "transaction_id": 123,
  "total_amount": 187000.00,
  "items": [...]
}
```

## Authentication Flow

```
┌──────────────────────────────────────────────────────────────┐
│                     LOGIN FLOW                                │
└──────────────────────────────────────────────────────────────┘

POST /api/login
{ username: "cashier1", password: "password123" }
     │
     ▼
┌────────────────┐
│  Auth Handler  │
└────┬───────────┘
     │
     ▼
┌────────────────┐
│  User Service  │
│  1. Find User  │
│  2. Verify PWD │───► bcrypt.Compare()
│  3. Gen Token  │───► JWT Sign (HS256)
└────┬───────────┘
     │
     ▼
Return JWT Token
{ "token": "eyJhbGc..." }


┌──────────────────────────────────────────────────────────────┐
│                  PROTECTED ROUTE FLOW                         │
└──────────────────────────────────────────────────────────────┘

GET /api/menus
Header: Authorization: Bearer eyJhbGc...
     │
     ▼
┌─────────────────┐
│ JWT Middleware  │
│ 1. Extract Token│
│ 2. Validate     │───► JWT Parse & Verify
│ 3. Parse Claims │
│ 4. Set Context  │───► c.Set("user_id", claims.UserID)
└────┬────────────┘
     │
     │ Token Valid ✓
     ▼
┌─────────────────┐
│  Menu Handler   │
│  Access Granted │
└─────────────────┘
```

## Data Flow: Complete Request Cycle

```
┌─────────┐
│ Client  │
└────┬────┘
     │ 1. HTTP Request (JSON)
     ▼
┌──────────────┐
│  Gin Router  │
└────┬─────────┘
     │ 2. Route Match
     ▼
┌─────────────────┐
│   Middleware    │ (JWT Auth for protected routes)
└────┬────────────┘
     │ 3. Authorized
     ▼
┌─────────────────┐
│    Handler      │ (Parse request, validate input)
└────┬────────────┘
     │ 4. Call Service
     ▼
┌─────────────────┐
│    Service      │ (Business logic, concurrency)
└────┬────────────┘
     │ 5. Call Repository
     ▼
┌─────────────────┐
│   Repository    │ (Data access)
└────┬────────────┘
     │ 6. GORM Query
     ▼
┌─────────────────┐
│     MySQL       │ (Database operations)
└────┬────────────┘
     │ 7. Return Data
     ▼
┌─────────────────┐
│   Repository    │ (Map to models)
└────┬────────────┘
     │ 8. Return Models
     ▼
┌─────────────────┐
│    Service      │ (Process, transform)
└────┬────────────┘
     │ 9. Return DTOs
     ▼
┌─────────────────┐
│    Handler      │ (Format response)
└────┬────────────┘
     │ 10. JSON Response
     ▼
┌─────────┐
│ Client  │
└─────────┘
```

## Technology Stack Diagram

```
┌────────────────────────────────────────────────────────────┐
│                      TECH STACK                             │
├────────────────────────────────────────────────────────────┤
│                                                             │
│  Frontend (Not in this project)                            │
│  ┌────────────┐  ┌────────────┐                           │
│  │   React/   │  │  Flutter   │                           │
│  │  Next.js   │  │   Mobile   │                           │
│  └────────────┘  └────────────┘                           │
│         │                │                                  │
│         └────────────────┘                                  │
│                  │                                          │
│         HTTP/HTTPS (JSON)                                   │
│                  │                                          │
│  ────────────────▼──────────────────────────────────────   │
│                                                             │
│  Backend (This Project)                                     │
│  ┌──────────────────────────────────────────────┐         │
│  │  Go 1.21+                                    │         │
│  │  ┌────────────┐  ┌────────────┐             │         │
│  │  │    Gin     │  │   GORM     │             │         │
│  │  │ Framework  │  │    ORM     │             │         │
│  │  └────────────┘  └────────────┘             │         │
│  │  ┌────────────┐  ┌────────────┐             │         │
│  │  │    JWT     │  │   Viper    │             │         │
│  │  │   Auth     │  │   Config   │             │         │
│  │  └────────────┘  └────────────┘             │         │
│  │  ┌────────────┐  ┌────────────┐             │         │
│  │  │ Goroutines │  │  Channels  │             │         │
│  │  │ Concurrent │  │   Sync     │             │         │
│  │  └────────────┘  └────────────┘             │         │
│  │  ┌────────────┐                             │         │
│  │  │  bcrypt    │                             │         │
│  │  │  Password  │                             │         │
│  │  └────────────┘                             │         │
│  └──────────────────────────────────────────────┘         │
│                  │                                          │
│         MySQL Driver (go-sql-driver)                        │
│                  │                                          │
│  ────────────────▼──────────────────────────────────────   │
│                                                             │
│  Database                                                   │
│  ┌──────────────────────────────────────────────┐         │
│  │  MySQL 8.0                                   │         │
│  │  • Users                                     │         │
│  │  • Menus                                     │         │
│  │  • Transactions                              │         │
│  │  • Transaction Details                       │         │
│  └──────────────────────────────────────────────┘         │
│                                                             │
└────────────────────────────────────────────────────────────┘
```

## Deployment Architecture

```
┌──────────────────────────────────────────────────────────────┐
│                    DEPLOYMENT OPTIONS                         │
└──────────────────────────────────────────────────────────────┘

Option 1: Traditional Server
┌────────────────────────────────────────┐
│  Linux Server (Ubuntu/CentOS)          │
│  ┌──────────────────────────────────┐  │
│  │  Nginx (Reverse Proxy)           │  │
│  │  :80/:443 → :8080                │  │
│  └─────────┬────────────────────────┘  │
│            │                            │
│  ┌─────────▼────────────────────────┐  │
│  │  Cashier API (:8080)             │  │
│  │  (systemd service)               │  │
│  └─────────┬────────────────────────┘  │
│            │                            │
│  ┌─────────▼────────────────────────┐  │
│  │  MySQL Database                  │  │
│  └──────────────────────────────────┘  │
└────────────────────────────────────────┘

Option 2: Docker Compose
┌────────────────────────────────────────┐
│  Docker Host                           │
│  ┌──────────────────────────────────┐  │
│  │  api container                   │  │
│  │  (service-cashier:latest)        │  │
│  │  Port: 8080                      │  │
│  └─────────┬────────────────────────┘  │
│            │ network bridge            │
│  ┌─────────▼────────────────────────┐  │
│  │  mysql container                 │  │
│  │  (mysql:8.0)                     │  │
│  │  Volume: mysql_data              │  │
│  └──────────────────────────────────┘  │
└────────────────────────────────────────┘

Option 3: Cloud (AWS/GCP/Azure)
┌────────────────────────────────────────┐
│  Cloud Provider                        │
│  ┌──────────────────────────────────┐  │
│  │  Load Balancer                   │  │
│  └─────────┬────────────────────────┘  │
│            │                            │
│  ┌─────────▼────────────────────────┐  │
│  │  API Instances (Auto-scaling)    │  │
│  │  • Instance 1                    │  │
│  │  • Instance 2                    │  │
│  │  • Instance N                    │  │
│  └─────────┬────────────────────────┘  │
│            │                            │
│  ┌─────────▼────────────────────────┐  │
│  │  RDS MySQL (Managed)             │  │
│  │  • Multi-AZ                      │  │
│  │  • Automated Backups             │  │
│  └──────────────────────────────────┘  │
└────────────────────────────────────────┘
```

---

**Legend:**
- `│` : Vertical flow
- `▼` : Direction of flow
- `┌─┐` : Component boundary
- `───` : Connection/relationship

This architecture provides a scalable, maintainable, and performant API backend for the Cashier Application.
