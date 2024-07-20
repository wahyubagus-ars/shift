# Use the golang:1.22.4-alpine image as the builder stage
FROM golang:1.22.4-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary


# Use alpine image forstage two prepare final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/binary /app/binary

COPY ./build.sh /app/build.sh

RUN chmod +x /app/build.sh

EXPOSE 8080

ARG ENV
ENV ENV=$ENV

ENTRYPOINT ["/bin/sh", "/app/build.sh"]