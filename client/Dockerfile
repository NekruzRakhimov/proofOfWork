# Используем официальный образ Golang
FROM golang:1.18-alpine

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем исходный код
COPY client.go .

# Сборка приложения
RUN go build -o client .

# Запуск клиента
CMD ["./client"]
