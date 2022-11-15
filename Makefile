createdb:
	docker compose exec -it postgres createdb --username=postgres --owner=postgres bank

dropdb:
	docker compose exec -it postgres dropdb bank

postgres:
	docker-compose up -d 

migrateup:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:root@localhost:5432/bank?sslmode=disable" -verbose down

generate:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc