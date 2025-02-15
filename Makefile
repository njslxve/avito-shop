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
	@go test -coverpkg=./internal/... -coverprofile=coverage.out ./internal/... && go tool cover -func=coverage.out

.PHONY: e2eup
e2eup:
	@cd tests && docker compose up -d && go test ./e2e/...

.PHONY: e2edown
e2edown:
	@cd tests && docker compose down