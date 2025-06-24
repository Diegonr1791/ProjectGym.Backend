# Build stage
FROM golang:1.23-alpine AS builder

# Instalar dependencias del sistema
RUN apk add --no-cache git ca-certificates tzdata

# Establecer directorio de trabajo
WORKDIR /app

# Copiar go mod files
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar código fuente
COPY . .

# Build de la aplicación (con build tag para Swagger)
RUN go build -tags swag -o main cmd/main.go

# Production stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copiar el binario desde el stage de build
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"] 