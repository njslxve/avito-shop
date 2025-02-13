.PHONY: up
up:
	@docker compose up --build -d

.PHONY: down
down:
	@docker compose down && docker rmi avito-shop-avito-shop-service:latest

.PHONY: load
load:
	@cd k6-config && k6 run loadtest.js

.PHONY: test
test:
	@go test ./internal/... -cover