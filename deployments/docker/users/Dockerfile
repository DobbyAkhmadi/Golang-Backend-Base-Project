# Use the official Golang image as the base image
FROM golang:1.16 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application inside the container
RUN go build -o service1 ./cmd/service1

# Create a minimal runtime image
FROM debian:buster-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/service1 .

# Expose the port on which the Go application will listen
EXPOSE 8080

# Define the command to run the application
CMD ["./service1"]
