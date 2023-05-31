# Use an official Golang image as the base image
FROM golang:1.16-alpine

# Set the working directory in the container
WORKDIR /server

# Copy the server application code to the working directory
COPY . .

# Build the server application
RUN go build -o server .

# Expose the port on which the server will listen (change if needed)
EXPOSE 8000

# Set the command to run the server
CMD ["./server"]
