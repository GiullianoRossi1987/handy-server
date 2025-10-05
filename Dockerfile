FROM golang:latest AS builder
COPY . /app
WORKDIR /app

# COPY go.mod go.sum ./
RUN go mod download
RUN go build main

FROM alpine:latest AS final
RUN apk add libc6-compat gcompat

COPY --from=builder /app /app
WORKDIR /app
