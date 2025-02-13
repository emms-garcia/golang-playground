FROM golang:1.23.5 AS builder
WORKDIR /go/src/app
COPY . .
CMD ["go", "run", "."]
