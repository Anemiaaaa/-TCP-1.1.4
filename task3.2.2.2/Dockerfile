# Используем официальный образ Golang на базе Alpine
FROM golang:alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем все зависимости
RUN go mod download

# Копируем исходный код в рабочую директорию внутри контейнера
COPY . .

# Устанавливаем пакет dotenv
RUN go get github.com/joho/godotenv

# Собираем Go приложение
RUN go build -o main .

# Открываем порт 8080
EXPOSE 8080

# Команда для запуска исполняемого файла
CMD ["./main"]
