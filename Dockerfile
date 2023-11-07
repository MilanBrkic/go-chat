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

# Build the Go application inside the builder container
RUN CGO_ENABLED=0 go build -o main .

# Stage 2: Create a clean container
FROM alpine:3.14

# Create a directory to store the binary
WORKDIR /app

# Copy only the binary from the builder container to the clean container
COPY --from=builder /app/main .

# Expose the port that your Go application will listen on
EXPOSE 8080

# Define the command to run your Go application
CMD ["./main"]
