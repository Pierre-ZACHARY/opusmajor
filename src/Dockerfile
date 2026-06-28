# syntax=docker/dockerfile:1.7

FROM golang:1.26.4-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:3.21
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8080
USER app
ENTRYPOINT ["/app/server"]
