FROM golang:1.20-alpine

# Set environment variables for CGO support
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=arm64

# Install build dependencies for CGO and SQLite
RUN apk add --no-cache \
    gcc \
    g++ \
    musl-dev \
    make \
    sqlite \
    sqlite-dev \
    bash

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o task-service ./cmd/main.go

# Expose the port and run the binary
EXPOSE 8080
CMD ["./task-service"]

