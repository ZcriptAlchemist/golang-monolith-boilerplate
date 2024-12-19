# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /server

# Install Air for hot-reloading (used for development)
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files to leverage Docker cache
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if go.mod and go.sum are unchanged
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o /server/app ./cmd/main.go

# Stage 2: Final stage for production
FROM alpine:latest AS production

# Set the working directory inside the container
WORKDIR /server

# Copy the built executable from the builder stage
COPY --from=builder /server/app .

# Copy environment variables file to /server
COPY .env /server/.env

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the built application
CMD ["./app"]

# Stage 3: Development stage with Air
FROM golang:1.23-alpine AS development

# Set the working directory inside the container
WORKDIR /server

# Install Air for hot-reloading
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Expose port 8080 for development
EXPOSE 8080

# Start Air
CMD ["air", "-c", ".air.toml"]