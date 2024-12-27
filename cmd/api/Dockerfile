# Use a lightweight base image with Go
FROM golang:1.23 as builder

# Set working directory
WORKDIR /app

# Copy the API service code
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

# Copy the go.mod files and download dependencies
COPY go.mod go.work ./
RUN go mod download

# Build the API binary
RUN go build -o api ./cmd/api

# Use a minimal base image for running the binary
FROM gcr.io/distroless/base-debian12

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api .

# Expose the application port
EXPOSE 8080

# Run the binary
CMD ["./api"]

