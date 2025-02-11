.PHONY: up
up:
	@cd ./deploy && docker compose up --build -d

.PHONY: down
down:
	@cd ./deploy && docker compose down && docker rmi deploy-avito-shop-service:latest