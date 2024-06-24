# Use the official Golang image as a base image
FROM golang:1.22.2-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Set environment variables
ENV MONGO_URI="mongodb://mongo:27017"
ENV DB_NAME="rentfield_gc"
ENV PORT=8080

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
