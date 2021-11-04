FROM golang:1.17 as builder

WORKDIR /app

ENV CGO_ENABLED=1

COPY go.mod go.sum ./

RUN go mod download
RUN apt-get -y update && apt-get install -y sqlite3

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o goapi
