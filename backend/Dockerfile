# Etapa de compilación
ARG TARGETPLATFORM
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Instalar 'file' en la etapa de compilación
RUN apk add --no-cache file

# Mostrar la versión de Go y las variables de entorno
RUN go version
RUN go env

# Establecer variables de entorno para compilación cruzada
ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

ENV CGO_ENABLED=0
ENV GOOS=${TARGETOS}
ENV GOARCH=${TARGETARCH}
ENV GOARM=${TARGETVARIANT#v}

# Copiar los archivos go.mod y go.sum
COPY go.mod go.sum ./

# Descargar las dependencias
RUN go mod download

# Copiar el resto del código
COPY . .

# Compilar la aplicación
RUN go build -o newsletter-app cmd/api/main.go

# Verificar que el binario se ha creado y su arquitectura
RUN ls -l /app
RUN file /app/newsletter-app

# Etapa de producción
ARG TARGETPLATFORM
FROM alpine:latest

WORKDIR /app

# Instalar 'file' para verificar la arquitectura
RUN apk add --no-cache file

# Copiar el binario compilado desde la etapa anterior
COPY --from=builder /app/newsletter-app .

# Asegurar que el binario tiene permisos de ejecución
RUN chmod +x newsletter-app

# Verificar que el binario existe y su arquitectura
RUN ls -l /app
RUN file /app/newsletter-app

# Exponer el puerto en el que escucha la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./newsletter-app"]