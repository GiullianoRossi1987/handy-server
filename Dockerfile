FROM golang:latest
COPY . /app
WORKDIR /app

# COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN go build