# Deployment Guide

This guide covers deploying the Cashier API to production environments.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Local Development](#local-development)
3. [Production Build](#production-build)
4. [Docker Deployment](#docker-deployment)
5. [Cloud Deployment](#cloud-deployment)
6. [Environment Variables](#environment-variables)
7. [Database Migration](#database-migration)
8. [Monitoring](#monitoring)

---

## Prerequisites

- Go 1.21+
- MySQL 5.7+ or 8.0+
- Git
- Make (optional, for using Makefile commands)

---

## Local Development

### 1. Clone and Setup

```bash
git clone <repository-url>
cd service-cashier
cp .env.example .env
```

### 2. Configure Environment

Edit `.env`:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=cashier_db
JWT_SECRET=your-secret-key-change-in-production
SERVER_PORT=8080
```

### 3. Setup Database

```bash
mysql -u root -p < database_setup.sql
```

### 4. Install Dependencies

```bash
go mod download
go mod tidy
```

### 5. Run Development Server

```bash
go run ./cmd/server
```

Or with auto-reload (requires [Air](https://github.com/cosmtrek/air)):
```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with auto-reload
air
```

---

## Production Build

### 1. Build Binary

```bash
# Build for current platform
go build -o bin/cashier-api ./cmd/server

# Build for Linux (from macOS/Windows)
GOOS=linux GOARCH=amd64 go build -o bin/cashier-api-linux ./cmd/server

# Build with optimizations
go build -ldflags="-s -w" -o bin/cashier-api ./cmd/server
```

### 2. Run Binary

```bash
./bin/cashier-api
```

### 3. Create Systemd Service (Linux)

Create `/etc/systemd/system/cashier-api.service`:

```ini
[Unit]
Description=Cashier API Service
After=network.target mysql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/cashier-api
EnvironmentFile=/opt/cashier-api/.env
ExecStart=/opt/cashier-api/bin/cashier-api
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Enable and start:
```bash
sudo systemctl enable cashier-api
sudo systemctl start cashier-api
sudo systemctl status cashier-api
```

---

## Docker Deployment

### 1. Create Dockerfile

Create `Dockerfile` in project root:

```dockerfile
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o cashier-api ./cmd/server

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/cashier-api .

# Copy .env file (or use docker-compose environment)
COPY .env .env

EXPOSE 8080

CMD ["./cashier-api"]
```

### 2. Create docker-compose.yml

```yaml
version: '3.8'

services:
  mysql:
    image: mysql:8.0
    container_name: cashier_mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: cashier_db
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql
      - ./database_setup.sql:/docker-entrypoint-initdb.d/setup.sql
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      timeout: 20s
      retries: 10

  api:
    build: .
    container_name: cashier_api
    restart: always
    ports:
      - "8080:8080"
    environment:
      DB_HOST: mysql
      DB_PORT: 3306
      DB_USER: root
      DB_PASS: rootpassword
      DB_NAME: cashier_db
      JWT_SECRET: production-secret-key-change-me
      SERVER_PORT: 8080
    depends_on:
      mysql:
        condition: service_healthy

volumes:
  mysql_data:
```

### 3. Run with Docker Compose

```bash
# Build and start
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

---

## Cloud Deployment

### AWS EC2

1. **Launch EC2 Instance** (Ubuntu 22.04)
2. **Install Go and MySQL**:
```bash
sudo apt update
sudo apt install -y golang-go mysql-server
```

3. **Clone and Setup**:
```bash
git clone <repository-url>
cd service-cashier
cp .env.example .env
# Edit .env with production values
```

4. **Setup Database**:
```bash
sudo mysql < database_setup.sql
```

5. **Build and Run**:
```bash
go build -o cashier-api ./cmd/server
nohup ./cashier-api &
```

6. **Setup Nginx Reverse Proxy**:
```nginx
server {
    listen 80;
    server_name api.yourdomain.com;

    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

### Google Cloud Platform (Cloud Run)

1. **Create Dockerfile** (see Docker section)

2. **Build and Push**:
```bash
gcloud builds submit --tag gcr.io/PROJECT_ID/cashier-api
```

3. **Deploy**:
```bash
gcloud run deploy cashier-api \
  --image gcr.io/PROJECT_ID/cashier-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars DB_HOST=DB_IP,DB_NAME=cashier_db
```

### Heroku

1. **Create Procfile**:
```
web: ./bin/cashier-api
```

2. **Add buildpack**:
```bash
heroku buildpacks:set heroku/go
```

3. **Add MySQL addon**:
```bash
heroku addons:create cleardb:ignite
```

4. **Deploy**:
```bash
git push heroku main
```

---

## Environment Variables

### Required Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `DB_HOST` | MySQL host | `localhost` |
| `DB_PORT` | MySQL port | `3306` |
| `DB_USER` | MySQL username | `root` |
| `DB_PASS` | MySQL password | `secretpassword` |
| `DB_NAME` | Database name | `cashier_db` |
| `JWT_SECRET` | JWT signing key | `random-secret-key` |
| `SERVER_PORT` | API server port | `8080` |

### Production Recommendations

```env
# Use strong, random JWT secret (minimum 32 characters)
JWT_SECRET=your-very-long-random-secret-key-here-change-in-production

# Use environment-specific database
DB_HOST=production-db-host.com
DB_NAME=cashier_production

# Consider using connection pooling
DB_MAX_IDLE_CONNS=10
DB_MAX_OPEN_CONNS=100
```

---

## Database Migration

### Manual Migration

```bash
# Backup existing database
mysqldump -u root -p cashier_db > backup.sql

# Run new migration
mysql -u root -p cashier_db < migration.sql
```

### Using golang-migrate

Install:
```bash
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Create migrations:
```bash
migrate create -ext sql -dir db/migrations -seq create_users_table
```

Run migrations:
```bash
migrate -path db/migrations -database "mysql://user:pass@tcp(localhost:3306)/cashier_db" up
```

---

## Monitoring

### Health Check Endpoint

```bash
curl http://localhost:8080/health
```

### Logging

Add structured logging (e.g., [zap](https://github.com/uber-go/zap)):
```bash
go get -u go.uber.org/zap
```

### Metrics

Consider adding Prometheus metrics:
```bash
go get github.com/prometheus/client_golang/prometheus
```

### Application Performance Monitoring (APM)

Options:
- New Relic
- DataDog
- Elastic APM
- Sentry (error tracking)

---

## Security Checklist

- [ ] Change default JWT secret
- [ ] Use strong database passwords
- [ ] Enable HTTPS/TLS
- [ ] Set up firewall rules
- [ ] Regular security updates
- [ ] Database backups
- [ ] Rate limiting (e.g., using middleware)
- [ ] Input validation
- [ ] CORS configuration
- [ ] SQL injection protection (GORM handles this)

---

## Performance Tuning

### Database

```sql
-- Add indexes for better query performance
CREATE INDEX idx_transactions_cashier_created ON transactions(cashier_id, created_at);
CREATE INDEX idx_transaction_details_transaction ON transaction_details(transaction_id);
CREATE INDEX idx_menus_price ON menus(price);
```

### Go Application

```go
// Configure GORM connection pool
db.DB().SetMaxIdleConns(10)
db.DB().SetMaxOpenConns(100)
db.DB().SetConnMaxLifetime(time.Hour)
```

### Server

```bash
# Increase file descriptors
ulimit -n 65536
```

---

## Backup Strategy

### Database Backup

Daily automated backups:
```bash
#!/bin/bash
# backup.sh
DATE=$(date +%Y%m%d_%H%M%S)
mysqldump -u root -p$DB_PASS cashier_db > /backups/cashier_db_$DATE.sql
gzip /backups/cashier_db_$DATE.sql

# Keep only last 30 days
find /backups -name "*.sql.gz" -mtime +30 -delete
```

Add to crontab:
```bash
0 2 * * * /path/to/backup.sh
```

---

## Troubleshooting

### Connection Issues

```bash
# Test MySQL connection
mysql -h DB_HOST -u DB_USER -p DB_NAME

# Check if port is open
netstat -tuln | grep 8080

# Check service status
systemctl status cashier-api
```

### High Memory Usage

```bash
# Check memory usage
free -h

# Monitor Go process
top -p $(pgrep cashier-api)
```

### Database Deadlocks

Enable query logging:
```sql
SET GLOBAL general_log = 'ON';
SET GLOBAL log_output = 'TABLE';
SELECT * FROM mysql.general_log ORDER BY event_time DESC LIMIT 100;
```

---

## Support

For issues and questions:
- Check logs: `journalctl -u cashier-api -f`
- Review documentation
- Create GitHub issue

---

**Production Deployment Checklist**

- [ ] Environment variables configured
- [ ] Database setup and migrated
- [ ] SSL/TLS certificate installed
- [ ] Firewall configured
- [ ] Monitoring set up
- [ ] Backup strategy implemented
- [ ] Load testing completed
- [ ] Documentation updated
- [ ] Team trained on deployment process
