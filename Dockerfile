# --- Stage 1: Сборка приложения ---
FROM golang:1.21-alpine AS builder

# Устанавливаем необходимые пакеты, затем удаляем кеши
RUN apk add --no-cache git

WORKDIR /app

# Копируем файлы зависимостей отдельно для кэширования
COPY app/go.mod app/go.sum ./
RUN go mod download && go mod verify

# Копируем исходный код приложения
COPY app/*.go ./

# Собираем приложение (статический бинарник, отключаем CGO)
RUN CGO_ENABLED=0 GOOS=linux go build -o ownapp

# --- Stage 2: Финальный образ для запуска ---
FROM alpine:3.18

# Создаем непривилегированного пользователя
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

# Копируем собранный бинарник
COPY --from=builder /app/ownapp ./

# Копируем статику (она также монтируется как volume)
COPY static ./static

# Передаем права непривилегированному пользователю
RUN chown -R appuser:appgroup /app

# Переключаемся на непривилегированного пользователя
USER appuser

# Открываем порт (по умолчанию 8080)
EXPOSE 8080

# Запускаем приложение
CMD ["./ownapp"]