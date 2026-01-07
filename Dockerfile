# ---------- Build Stage ----------
FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Copy go mod files first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build ONLY the API binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o api-server ./cmd/api

# ---------- Runtime Stage ----------
FROM alpine:latest

# Create non-root user
RUN adduser -D appuser

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/api-server .

# Expose API port
EXPOSE 8080

USER appuser

# Start API
CMD ["./api-server"]
