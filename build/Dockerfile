# --- Stage 1: Builder ---
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container for the builder stage
WORKDIR /app

# Copy go.mod and go.sum files to the working directory to leverage build cache for dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application binary
RUN CGO_ENABLED=0 GOOS=linux go build -o go-server

# --- Stage 2: Runtime ---
# Use a minimal base image, like scratch or alpine, for the final image
FROM scratch

# Set the working directory for the runtime stage (optional but good practice)
WORKDIR /root/

# Copy the built binary from the 'builder' stage to the new stage
COPY --from=builder /app/go-server .

# Command to run the application when the container starts
CMD ["./go-server"]