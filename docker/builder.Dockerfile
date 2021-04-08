FROM golang:latest

WORKDIR /baumanfinder

COPY . .

RUN make build-app