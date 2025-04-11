FROM golang:1.24.2-alpine3.21 AS builder
WORKDIR /go/src/app
COPY . .

FROM builder AS dev
RUN go install github.com/air-verse/air@latest
CMD ["air", "-c", ".air.toml"]
