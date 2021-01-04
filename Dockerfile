# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /src
COPY . .
RUN go build -o bin/wombagd cmd/wombagd/wombagd.go
RUN go build -o bin/wombagcli cmd/wombagcli/wombagcli.go
