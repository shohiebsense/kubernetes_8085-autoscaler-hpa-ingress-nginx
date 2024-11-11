# Step 1: Use an official Golang image as the build environment
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application with CGO disabled
RUN CGO_ENABLED=0 go build -o networkTrafficMonitor .

# Step 2: Use a minimal image to run the built binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/networkTrafficMonitor .
RUN chmod +x networkTrafficMonitor

EXPOSE 8085

# Run the application
CMD ["./networkTrafficMonitor"]
