FROM golang:1.23-alpine

WORKDIR /app

# Instalar herramientas de desarrollo
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest

# Copiar archivos de configuración
COPY app/go.mod app/go.sum ./
RUN go mod download

# El código se montará como volumen