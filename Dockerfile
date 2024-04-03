# Use an official Golang runtime as a parent image
FROM golang:1.20.5 as base

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Install Deps
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Expose port 8080 for incoming traffic
EXPOSE 8080

# Define the command to run the app when the container starts
CMD ["/app/main"]