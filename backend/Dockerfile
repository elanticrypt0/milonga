FROM golang:1.23-alpine

WORKDIR /app

# Instalar herramientas necesarias
RUN apk add --no-cache git && \
    go install github.com/air-verse/air@latest

# Copiar archivos de configuración
COPY . .

# Inicializar módulo si no existe go.mod
RUN if [ ! -f go.mod ]; then \
    go mod init milonga && \
    go mod tidy; \
    fi

# Exponer el puerto que usará tu aplicación
EXPOSE 8921

# Usar air para hot-reload
CMD ["air"]