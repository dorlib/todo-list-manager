# Base image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code
COPY . .

# Install dependencies
RUN go get -d -v ./...
RUN go install -v ./...

# Build the application
ARG ENV_FILE
ENV MY_ENV_VAR=value
RUN if [ -f "$ENV_FILE" ]; then . "$ENV_FILE"; fi && \
    go build -o app

# Expose port 8080 for the application
EXPOSE 8080

# Start the application
CMD ["./app", "go run main.go"]