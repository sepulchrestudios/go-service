# --- Stage 1: Builder ---
FROM golang:1.25-alpine AS builder

# Copy any custom CA certificate(s) into the image
# https://ko.build/advanced/root-ca-certificates/#example_1
COPY build/certs/*.crt /usr/local/share/ca-certificates/

# Install the ca-certificates package and update the certificate store
# https://ko.build/advanced/root-ca-certificates/#example_1
RUN apk update && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

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

# Copy the CA certificate bundle from the builder stage
# https://medium.com/on-docker/use-multi-stage-builds-to-inject-ca-certs-ad1e8f01de1b
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Command to run the application when the container starts
CMD ["./go-server"]