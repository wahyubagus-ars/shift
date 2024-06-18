FROM golang:1.22.4-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go install github.com/air-verse/air@latest
RUN go install github.com/google/wire/cmd/wire@latest

COPY . .

# Add step to execute generate wire in the /app/cmd/app/provider directory
# RUN cd ./cmd/app/provider && wire

RUN go build -o main .

EXPOSE 8080

CMD ["air"]