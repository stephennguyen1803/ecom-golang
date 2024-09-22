#Make the default Value
GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING="root:admin@tcp(127.0.0.1:3306)/shopDevGo?parseTime=true"
GOOSE_DIR=sql/schema

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_DIR) reset

.PHONY: upse downse resetse