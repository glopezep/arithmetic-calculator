DB_URL=postgresql://test:test@localhost:5432/test?sslmode=disable

build:
	sam build

dev:
	sam local start-api

createdb:
	docker exec -it postgres createdb --username=test --owner=test test

dropdb:
	docker exec -it postgres dropdb --username=test test

migrateup:
	migrate -path internal/infrastructure/db/migrations -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/infrastructure/db/migrations -database "$(DB_URL)" -verbose down

createmigration:
	migrate create -ext sql -dir internal/infrastructure/db/migrations -seq action_name -database "$(DB_URL)" -verbose down

test:
	go test ./... --cover


.PHONY: build, dev, createdb, dropdb, migrateup, migratedown, createmigration, test
