# Etapa 1: Build
FROM golang:1.24.1-alpine AS builder
WORKDIR /app

# Copia os arquivos de dependências e baixa os módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código e compila a aplicação
COPY . .
RUN go build -o main ./cmd/main.go

# Etapa 2: Imagem final mais enxuta
FROM alpine:latest
WORKDIR /root/

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/main .

# Expõe a porta (ajuste conforme sua aplicação)
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./main"]
