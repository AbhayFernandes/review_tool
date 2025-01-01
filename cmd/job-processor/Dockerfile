# Use a lightweight base image with Go
FROM golang:1.23 as builder

# Set working directory
WORKDIR /app

# Copy the Job Processor code
COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

# Copy the go.mod and  files and download dependencies
COPY go.mod go.work ./
RUN go mod download

# Build the Job Processor binary
RUN go build -o job-processor ./cmd/job-processor

# Use a minimal base image for running the binary
FROM gcr.io/distroless/base-debian10

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/job-processor .

# Expose the application port (if applicable)
# EXPOSE 8081 # Uncomment if the job processor exposes a port

# Run the binary
CMD ["./job-processor"]

