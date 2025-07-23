# 1. Используем официальный образ Golang
FROM golang:1.21 AS builder

# 2. Устанавливаем рабочую директорию
WORKDIR /app

# 3. Копируем все файлы
COPY . .

# 4. Сборка бинарника
RUN go build -o app ./cmd/app

# ----------------------------
# 5. Минимальный образ для запуска
FROM debian:bookworm-slim

# 6. Создаем рабочую папку
WORKDIR /app

# 7. Копируем бинарник и конфиг
COPY --from=builder /app/app .
COPY --from=builder /app/config.yml .

# 8. Указываем порт (если нужно)
EXPOSE 8080

# 9. Команда запуска
CMD ["./app"]
