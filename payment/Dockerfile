FROM golang:alpine AS builder
WORKDIR /app

COPY . .
WORKDIR /app/microservices/payment
RUN go mod tidy
RUN go build -o payment_app cmd/main.go

FROM alpine:latest
WORKDIR /root/

# Variáveis de ambiente
ENV DB_DRIVER=mysql
ENV APPLICATION_PORT=3001
ENV ENV=development

COPY --from=builder /app/microservices/payment/payment_app .
CMD ["./payment_app"]