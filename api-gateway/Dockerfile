FROM golang:1.23 AS builder

# Рабочая директория для сборки
WORKDIR /app

# Копируем go.mod и go.sum для скачивания зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Сборка бинарника (чистый statically linked бинарь)
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./cmd/api-gateway/main.go

# ===============================
# Stage 2: Minimal image for production
# ===============================
FROM alpine:latest

# Минимальный runtime (если нужно)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Копируем только собранный бинарник
COPY --from=builder /app/api-gateway .

# Открываем порты
EXPOSE 8080
EXPOSE 40000