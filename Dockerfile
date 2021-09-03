# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./
COPY localhost.crt ./
COPY localhost.key ./

RUN go build -o /server

EXPOSE 8000
EXPOSE 9000

CMD [ "/server" ]