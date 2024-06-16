FROM golang:1.22.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/air-verse/air@latest

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["air"]