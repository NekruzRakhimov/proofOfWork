# Используем официальный образ Golang
FROM golang:1.18-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код
COPY server.go .

# Сборка приложения
RUN go build -o server .

# Запуск сервера
CMD ["./server"]
