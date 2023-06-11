# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go dependencies
RUN go mod download

# Copy the remaining source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8000
EXPOSE 8000

# Set the entry point of the container
CMD ["./main"]

