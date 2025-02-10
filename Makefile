.PHONY: up
up:
	@cd ./deploy && docker compose up -d

.PHONY: down
down:
	@cd ./deploy && docker compose down