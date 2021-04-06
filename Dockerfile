FROM golang:latest

COPY ./ ./
RUN go build -o .bin/docparser cmd/docparser/main.go && \
    .bin/docparser