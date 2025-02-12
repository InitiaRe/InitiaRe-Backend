# syntax=docker/dockerfile:1
FROM golang:alpine

WORKDIR /code

RUN apk add --no-cache git && \
    go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
