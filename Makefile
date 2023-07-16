postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root postgres12

dropdb:
	docker exec -it postgres12 dropdb postgres12

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres12?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres12?sslmode=disable" -verbose down

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

psql:
	docker exec -it postgres12 psql -U root -d postgres12

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test psql server
