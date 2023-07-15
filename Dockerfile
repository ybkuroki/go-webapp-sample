# Base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o app

# Port exposed by the application
EXPOSE 8080

# Command to run the application
CMD ["./app"]