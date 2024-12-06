# Stage 1: Build the Go application
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire application source
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Stage 2: Create a lightweight container
FROM alpine:latest

# Set working directory in the container
WORKDIR /app

# Copy the compiled Go binary from the builder
COPY --from=builder /app/main .

# Expose application port
EXPOSE 8080

# Command to run the binary
CMD ["./main"]

