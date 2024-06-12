# Use an official Golang runtime as a base image
FROM golang:1.22.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the local package files to the container's workspace
COPY . .

# Download and install any required dependencies
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
