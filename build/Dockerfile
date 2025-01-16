# Stage 1: Build the Go application using the Golang image
FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY configs/ /app/configs/
RUN go build -o ./cmd/main ./cmd

FROM ubuntu:22.04
WORKDIR /app
RUN mkdir -p /app/configs
COPY --from=builder /app/cmd/main ./cmd/main
COPY --from=builder /app/configs/ /app/configs/
EXPOSE 80

CMD ["./cmd/main"]