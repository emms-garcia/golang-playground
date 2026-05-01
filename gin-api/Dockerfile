FROM golang:1.26.2-alpine AS builder
WORKDIR /go/src/app
COPY . .

FROM builder AS dev
RUN go install github.com/air-verse/air@v1.65.1
CMD ["air", "-c", ".air.toml"]
