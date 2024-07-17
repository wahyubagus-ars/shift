FROM golang:1.22.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go build -o main .

EXPOSE 8080

COPY ./build.sh /app/build.sh

# Make the entrypoint script executable
RUN chmod +x /app/build.sh

# Set the build-time ARG as an ENV variable inside the Docker image
ARG ENV
ENV ENV=$ENV

# Use the shell to run the script
ENTRYPOINT ["/bin/sh", "/app/build.sh"]
