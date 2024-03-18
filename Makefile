.PHONY: run
run:
	go run cmd/filmLibrary/main.go

.PHONY: build
build:
	go build -o ./bin/app ./cmd/filmLibrary/main.go

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

.PHONY: rollback
rollback:
	dotenv -- tern migrate -m migrations/ -d -1

.PHONY: test
test:
	go test -coverpkg=./... -coverprofile=c.out.tmp ./...

.PHONY: cover
cover: test
	go tool cover -func=c.out.tmp
