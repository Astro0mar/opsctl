FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o opsctl

FROM debian:stable-slim

WORKDIR /root/

COPY --from=builder /app/opsctl .

CMD ["./opsctl", "version"]
