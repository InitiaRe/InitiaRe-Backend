# syntax=docker/dockerfile:1
FROM golang:alpine

WORKDIR /code
COPY . .
RUN go install