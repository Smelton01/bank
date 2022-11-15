package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/smelton01/bank/api"
	db "github.com/smelton01/bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5432/bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
