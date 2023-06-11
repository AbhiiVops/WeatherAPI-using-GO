# Using  the official Golang image as the base image for the container
FROM golang:1.20-alpine

# Setting the current working directory inside the container to /app
WORKDIR /app

# Copying the Go module files to the current working directory
COPY go.mod go.sum ./

# Downloading the Go dependencies for the project
RUN go mod download

# Copying the remaining source code to the current working directory
COPY . .

# Building the Go application inside the container
RUN go build -o main .

# Exposing port 8000 to the host machine
EXPOSE 8000

# Setting the entry point of the container to the executable we just built
CMD ["./main"]

