# syntax = docker/dockerfile:experimental
FROM golang:latest

WORKDIR /app

# Force the go compiler to use modules
ENV GO111MODULE=on
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .