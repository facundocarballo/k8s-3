# Etapa 1: build del binario
FROM golang:1.24-alpine AS builder

COPY go.mod /opt/app/
COPY go.sum /opt/app/

WORKDIR /opt/app

RUN go mod download

COPY src /opt/app/src

# Compilamos el binario estatico
RUN go build -o server ./src

# Etapa 2: imagen final minimalista
FROM alpine:latest

WORKDIR /usr/local/bin

# Copiamos solo el binario desde la etapa anterior
COPY --from=builder /opt/app/server .

# Comando para ejecutar el servidor
CMD ["./server"]
