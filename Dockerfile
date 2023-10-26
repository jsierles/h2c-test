FROM golang:1.21 AS builder
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

### Start a new stage from scratch ###
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server .

# Expose the ports the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./server"]

