package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:root@localhost:5432/bank?sslmode=disable"
)

type MainSuite struct {
	suite.Suite

	queries *Queries
	db      *sql.DB
}

func (s *MainSuite) SetupSuite() {
	testDB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("could not connect to database: %s", err)
	}

	s.db = testDB
	s.queries = New(testDB)
	// os.Exit(m.Run())

}

func (s *MainSuite) TeardownSuite() {
	log.Println("done!!!")
}

func TestMain(t *testing.T) {
	suite.Run(t, &MainSuite{})
}
