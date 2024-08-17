FROM golang:1.22.4-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o /bin/api-crud ./cmd/api-crud/main.go

FROM scratch

WORKDIR /app

COPY --from=build /app/api-crud /app/api-crud

COPY --from=build /app/frontend /app/frontend

EXPOSE 8080

ENTRYPOINT ["./api-crud"]
