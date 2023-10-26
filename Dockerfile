# Use the official Go image as the base image
FROM golang:1.17 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the Working Directory inside the container
COPY . .

# Fetch the dependencies
RUN go mod init server && go get golang.org/x/net/http2

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

### Start a new stage from scratch ###
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/server .

# Expose the ports the app runs on
EXPOSE 8080
EXPOSE 8081
EXPOSE 8082

# Command to run the executable
CMD ["./server"]

