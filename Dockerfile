# Stage 1: Build the Go application
FROM golang:1.20.5-alpine AS builder

# Set necessary environment variables for cross-compilation
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o /app/paper-rocket .

# Stage 2: Create the final image with a specific version of Alpine
FROM alpine:3.17

# Set the working directory inside the container
WORKDIR /root/

# Copy the built Go application from the builder stage
COPY --from=builder /app/paper-rocket .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the Go application
CMD ["./paper-rocket"]
