.PHONY: build run migrate-up migrate-down generate clean

BINARY := bin/crewport
DB_PATH ?= ./crewport.db
MIGRATIONS_PATH := internal/db/migrations

build:
	@mkdir -p bin
	CGO_ENABLED=1 go build -o $(BINARY) ./cmd/crewport

run: build
	./$(BINARY)

migrate-up:
	@echo "Running migrations up..."
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1 \
		-path $(MIGRATIONS_PATH) \
		-database "sqlite3://$(DB_PATH)" \
		up

migrate-down:
	@echo "Running migrations down (1 step)..."
	go run -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1 \
		-path $(MIGRATIONS_PATH) \
		-database "sqlite3://$(DB_PATH)" \
		down 1

generate:
	@echo "Running sqlc generate..."
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0 generate

clean:
	rm -rf bin/
	rm -f *.db *.db-shm *.db-wal

tidy:
	go mod tidy
