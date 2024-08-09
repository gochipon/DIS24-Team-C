.PHONY: create_migration setup

setup:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create_migration:
ifndef NAME
	@echo "Usage: make NAME=migration_name create_migration"
else
	migrate create -ext sql -dir db/migrations -seq $(NAME)
endif
