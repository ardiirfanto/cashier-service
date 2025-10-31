# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o cashier-api ./cmd/server

# Runtime stage
FROM alpine:latest

# Install ca-certificates and tzdata
RUN apk --no-cache add ca-certificates tzdata

# Set timezone to Jakarta
ENV TZ=Asia/Jakarta

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/cashier-api .

# Copy config file if exists (optional, can use env vars instead)
COPY --from=builder /app/config.yaml* ./

EXPOSE 8080

# Run as non-root user for security
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /app

USER appuser

CMD ["./cashier-api"]