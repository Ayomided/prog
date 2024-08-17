ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm AS builder

# Install SQLite
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download && go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Generate templ files
RUN templ generate

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM debian:bookworm

RUN apt-get update && apt-get install -y ca-certificates sqlite3 libsqlite3-0 && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Copy the SQLite database into the container
COPY blog.db /root/data/blog.db

# Command to run the executable
CMD ["./main"]
