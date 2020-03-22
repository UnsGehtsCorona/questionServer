FROM golang:alpine as builder
MAINTAINER "Benedict Burckhart" <burckhartb@gmail.com>

ENV GO111MODULE=on

RUN mkdir -p /app

WORKDIR /app
COPY . .

RUN go build -o main

FROM alpine as prod
RUN mkdir -p /app

COPY --from=builder /app /app
CMD ["/app/main"]
