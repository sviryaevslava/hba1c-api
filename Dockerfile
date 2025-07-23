FROM golang:1.21 AS builder
WORKDIR /app

COPY . .
RUN go build -o app ./cmd/app

FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/config.yml .

EXPOSE 8080
CMD ["./app"]
