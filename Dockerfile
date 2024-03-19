# Используем образ Golang в качестве базового образа
FROM golang:latest AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем содержимое текущей директории внутрь контейнера
COPY . .

# Собираем приложение
RUN go build -o build/adamenblog-api ./cmd/adamenblog-api

# Второй этап сборки
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем бинарник из предыдущего этапа
COPY --from=builder /app/build/adamenblog-api /app/adamenblog-api

# Указываем порт, который будет слушать приложение
EXPOSE 8080

# Запускаем приложение при старте контейнера
CMD ["/app/adamenblog-api"]
