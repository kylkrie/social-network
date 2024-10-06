# Start from the official Golang image for a specific version
FROM golang:1.23 as local
WORKDIR /app
# Install air for hot reloading
RUN go install github.com/air-verse/air@latest
# Copy go.mod and go.sum to download the dependencies
COPY go.mod .
COPY go.sum .
# Download the dependencies - this can be cached unless
# go.mod or go.sum change
RUN go mod download
# Copy the source code into the container
COPY . .
# Run air for hot reloading
CMD ["air"]

# Build stage
FROM golang:1.23 as build
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Copy the source code
COPY . .
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Production stage
FROM alpine:latest as prod
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy the pre-built binary file from the build stage
COPY --from=build /app/main .
# Expose the port the app runs on
EXPOSE 8080
# Command to run the executable
CMD ["./main"]
