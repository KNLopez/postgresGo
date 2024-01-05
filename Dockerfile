# write go dockerfile
FROM golang:latest

WORKDIR /app

# Copy the entire project
COPY . .

# Download the necessary modules
RUN go mod download

# Build the application
RUN go build -o main ./cmd/main

# Run the application
CMD ["/app/main"]