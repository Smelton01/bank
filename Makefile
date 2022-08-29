createdb:
	docker exec -it postgres14 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres14 dropdb bank

postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14.5-alpine

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down

