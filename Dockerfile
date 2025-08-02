FROM golang:latest AS builder

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

# Build the Go application
# CGO_ENABLED=0 disables cgo for static linking, resulting in a smaller binary
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o app ./cmd/web

# Use a minimal base image for the final stage
FROM alpine:latest as run

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]