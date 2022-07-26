postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=reginapost -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:reginapost@localhost:5432/simple_bank?sslmode=disable" -verbose up	

migrateup1:
	migrate -path db/migration -database "postgresql://root:reginapost@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:reginapost@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:reginapost@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
    docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate

server: 
	go run main.go

test:
	go test -v -cover ./...

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/rewhatsmore/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock