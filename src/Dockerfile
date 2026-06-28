# syntax=docker/dockerfile:1.7

FROM --platform=$BUILDPLATFORM golang:1.26.4-alpine AS builder
ARG TARGETOS
ARG TARGETARCH

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o server .

FROM alpine:3.21
RUN addgroup -S app && adduser -S app -G app
WORKDIR /app

COPY --from=builder /app/server /app/server

EXPOSE 8080
USER app
ENTRYPOINT ["/app/server"]
