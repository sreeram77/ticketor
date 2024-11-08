# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy Go source code
COPY . .

# Build the Go application
RUN go build -o main cmd/server/main.go

# Stage 2: Create the final runtime image
FROM alpine:3.20

# Copy the built binary from the builder stage
COPY --from=builder /app/main /app/

# Set the working directory
WORKDIR /app

# Expose the port your application listens on
EXPOSE 8000

# Command to run the application
CMD ["/app/main"]