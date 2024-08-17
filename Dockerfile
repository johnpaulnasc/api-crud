# 1. Fase de construção
FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin/api-crud ./cmd/api-crud/main.go

# 2. Fase final, utilizando a imagem scratch para minimizar o tamanho
FROM scratch

WORKDIR /app

# Copiando o binário construído
COPY --from=builder /bin/api-crud /app/api-crud

# Copiando os arquivos estáticos e templates necessários
COPY --from=builder /app/frontend /app/frontend

EXPOSE 8080

ENTRYPOINT ["/app/api-crud"]
