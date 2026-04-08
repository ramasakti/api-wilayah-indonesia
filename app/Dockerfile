# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Set environment variables
ENV PORT=8080
ENV DATA_DIR=/app/data
ENV GIN_MODE=release

# Expose port
EXPOSE 8080

# Command to run
CMD ["./main"]
