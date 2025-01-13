FROM golang:1.23.4-alpine

# Install dependencies
RUN apk add --no-cache bash git

# Install Air for live reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory
WORKDIR /BooksAPI

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code (Ensure everything is copied, including `cmd/BooksAPI/main.go`)
COPY . .  

# Check if the Go files are present in the container
RUN ls -R /BooksAPI

# Build the Go application and place the binary in the root of the project
RUN go build -o /BooksAPI/main /BooksAPI/cmd/BooksAPI/main.go

# Set the entrypoint to Air (for live reloading)
CMD ["air"]

