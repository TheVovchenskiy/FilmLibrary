.PHONY: run
run:
	go run cmd/filmLibrary/main.go

.PHONY: compose-up
compose-up:
	docker compose up -d

.PHONY: create-migration
create-migration:
	tern new -m migrations/ $(name)

.PHONY: install-dotenv
install-dotenv:
	npm install -g dotenv-cli

.PHONY: migrate
migrate:
	dotenv -- tern migrate -m migrations/
