# Use the official Go image as a parent image
FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o /otel-prometheus-study cmd/api/main.go

# Expose the port the API runs on
EXPOSE 8000
# Expose the port for Prometheus metrics
EXPOSE 2112

# Command to run the executable
CMD [ "/otel-prometheus-study" ]