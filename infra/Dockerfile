FROM golang:alpine3.16 AS builder

WORKDIR /app

COPY . ./

EXPOSE 8080

RUN go mod tidy

CMD go run cmd/main.go