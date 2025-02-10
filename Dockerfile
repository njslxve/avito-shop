FROM golang:1.22-alpine3.19 AS builder

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /app
COPY . .

RUN go build -o ./bin/avito-shop ./cmd/avito-shop/main.go \
    && go clean -cache -modcache

FROM alpine:latest

COPY --from=builder app/bin/avito-shop /avito-shop
    
EXPOSE 8080

CMD ["/avito-shop"]