# Используем Dockerfile.system как базовый образ
FROM system AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы зависимостей и загружаем их
COPY app/go.mod app/go.sum ./
RUN go mod download && go mod verify

# Копируем исходный код приложения
COPY app/*.go ./

# Собираем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o ownapp

# Создаем директорию для бинарника
RUN mkdir /output && cp ownapp /output/