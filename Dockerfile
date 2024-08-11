FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o go-ip

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/go-ip .

CMD ["./go-ip"]

EXPOSE 8080
