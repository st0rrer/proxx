# syntax=docker/dockerfile:1.3
FROM golang:1.19-buster as base

RUN apt-get update
RUN apt-get install -y libreadline-dev

FROM base as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
go build -o ./bin/proxx cmd/*.go

FROM debian:buster-slim

RUN apt-get update
RUN apt-get install -y libreadline-dev

WORKDIR /app
COPY --from=builder /app/bin/proxx .

CMD ["./proxx"]