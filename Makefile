BINARY=bin/engine

DB := grpc-demo-server
DB_DIR := db
POSTGRES_USER := root
POSTGRES_HOST := localhost
POSTGRES_PASSWORD := password
POSTGRES_PORT := 5432


migrate-prepare:
	@echo "Installing golang-migrate"
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create-db:
	postgres create database --username=${POSTGRES_USER} --password=${POSTGRES_PASSWORD} --owner=${POSTGRES_USER} ${DB};

migrate-up:
	@-migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ up

migrate-down:
	@-migrate -database 'postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(DB)?sslmode=disable' -path $(DB_DIR)/migrations/ down

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

.PHONY: migrate-prepare db-migrateup db-migratedown proto