# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

WORKDIR /app

COPY ./src ./src

WORKDIR /app/src

RUN go mod download
RUN go build -o /app/exec ./cmd

EXPOSE 80

CMD [ "/app/exec" ]
