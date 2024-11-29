FROM golang:1.23.1-bullseye AS build

WORKDIR /post_service

#RUN apt-get install protobuf-compiler

RUN apt-get update && apt-get install -y protobuf-compiler

COPY . .

COPY go.mod go.sum ./

#ENV PATH="$PATH:$(go env GOPATH)/bin"

RUN go mod download

RUN go build ./cmd/main.go

EXPOSE 8080

CMD ["./main"]