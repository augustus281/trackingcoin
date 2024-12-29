# Use an official Golang image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go mod and sum files
COPY go.mod go.sum ./

# Download dependencies (to cache them in the Docker image)
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Expose port (assuming your app runs on port 8080)
EXPOSE 5000

# Command to run the app
CMD ["./main"]
