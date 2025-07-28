# Build stage
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for dependency caching
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o app main.go

# Production stage
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose port (change if your app uses a different port)
EXPOSE 8080

# Command to run the executable
CMD ["./app"]