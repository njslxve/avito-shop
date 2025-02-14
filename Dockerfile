FROM golang:1.24-alpine3.21 AS builder

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /app
COPY . .

RUN go build -o ./bin/avito-shop ./cmd/avito-shop/main.go

FROM alpine:latest

COPY --from=builder app/bin/avito-shop /avito-shop
    
EXPOSE 8080

CMD ["/avito-shop"]