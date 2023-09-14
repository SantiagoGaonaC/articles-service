# Etapa de compilación
FROM golang:1.20-alpine3.18 AS builder

# Instalar git para la descarga de paquetes
RUN apk --no-cache add git

# Establecer el directorio de trabajo
WORKDIR /src

# Copiar y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o /build/api ./cmd/api

# Etapa de ejecución
FROM alpine:3.18

# Variables de entorno
ENV JWT_SECRET="secret" \
    JWT_EXPITATION_DAYS=7 \
    DATABASE_NAME="products" \
    DATABASE_USER="postgres" \
    DATABASE_PASSWORD="postgres" \
    DATABASE_HOST="postgresql" \
    DATABASE_PORT="5432" \
    DATABASE_SSLMODE="disable" \
    DATABASE_LOGMODE="false"

# Copiar el binario construido desde la etapa builder
COPY --from=builder /build/api /service-bin/api

# Exponer el puerto
EXPOSE 8080

# Ejecutar la aplicación
CMD ["/service-bin/api"]
