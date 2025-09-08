# Etapa 1: Construcción
FROM golang:1.24.1 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./
RUN go build -o main .

# Etapa 2: Imagen final liviana
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
