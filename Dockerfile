FROM golang:1.23.0-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin/api-crud ./cmd/api-crud/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /bin/api-crud /app/api-crud

COPY --from=builder /app/frontend /app/frontend

EXPOSE 8080

ENTRYPOINT ["/app/api-crud"]
