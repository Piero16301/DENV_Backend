FROM golang:1.18-alpine

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY ${PWD} /app
RUN go mod download

# Build the binary.
RUN go build -o /denv-backend

# Run backend on port 8080
EXPOSE 8080

# Run the web service on container startup
CMD ["/dev-backend"]