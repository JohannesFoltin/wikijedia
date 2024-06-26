# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Johannes Foltin <Johannes.foltin@gmx.de>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Enable cgo
RUN go env -w CGO_ENABLED=1

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux go build -o /main

# Expose port 18143 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/main"]