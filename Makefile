createdb:
	docker exec -it postgres14 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres14 dropdb bank

postgres:
	# docker start postgres14 || docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14.5-alpine
	docker-compose up -d 

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down

generate:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb postgres migrateup migratedown sqlc