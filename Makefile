postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root postgres12

dropdb:
	docker exec -it postgres12 dropdb postgres12

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres12?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres12?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres12?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/postgres12?sslmode=disable" -verbose down 1

makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

psql:
	docker exec -it postgres12 psql -U root -d postgres12

server:
	go run main.go

start:
	docker start postgres12

stop:
	docker stop postgres12

check:
	docker ps

mock:
	mockgen -package mockdb -destination db/mock/store.go backend_masterclass/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test psql server start check mock stop migrateup1 migratedown1
