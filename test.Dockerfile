# Stage 1: Build the Go application
FROM golang:1.20 AS builder

# Set the working directory in the builder container
WORKDIR /app

# Copy the Go module definition and the go.sum file to the builder container
COPY go.mod go.sum ./

# Download and install Go module dependencies
RUN go mod download
RUN go mod tidy

# Copy the rest of the application code to the builder container
COPY . .
