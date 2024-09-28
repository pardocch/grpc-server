# Use the official Go base image
FROM golang:1.19

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all files from the current directory to the working directory in the container
COPY . .

# Build the gRPC server
RUN go build -o grpc-server

# Expose the port for the gRPC server (e.g., 50051)
EXPOSE 50051

# Run the compiled server
CMD ["./grpc-server"]

