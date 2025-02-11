.PHONY: up
up:
	@docker compose up --build -d

.PHONY: down
down:
	@docker compose down && docker rmi deploy-avito-shop-service:latest