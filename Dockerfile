FROM golang:latest AS builder

ENV GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
WORKDIR /build/cmd/toky
RUN go build -o toky .

FROM alpine:latest

RUN apk add --no-cache bash

WORKDIR /root

COPY --from=builder /build/cmd/toky/toky .

ARG DB_PASSWORD
ENV DB_PASSWORD=${DB_PASSWORD}

ENTRYPOINT ["./toky"]
CMD ["c=./config/local.yml"]
