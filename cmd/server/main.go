package server

import (
	"database/sql"
	"log"

	"github.com/smelton01/bank/api"
	db "github.com/smelton01/bank/db/sqlc"
	"github.com/smelton01/bank/util"
)

func main() {
	StartServer()
}

func StartServer() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
