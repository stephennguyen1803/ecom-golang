#Make the default Value
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING="shopdev:shopdev@tcp(127.0.0.1:3306)/shopdevgo?parseTime=true"
GOOSE_DIR=sql/schema
GOOSE_MIGRATION_DIR=sql/schema

up_by_one:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_DIR) up-by-one
create_migration:
	@goose -dir=$(GOOSE_MIGRATION_DIR) create $(name) sql
upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_DIR) reset

sqlc_generate:
	@sqlc generate
swagger_generate:
	@swag init -g main.go -o  ./cmd/swag/docs --dir ./cmd/server

.PHONY: upse downse resetse create_migration up_by_one