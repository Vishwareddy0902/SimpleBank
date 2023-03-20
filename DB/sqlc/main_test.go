package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Unable to connect to DB", err)
	}
	testQueries = New(con)
	os.Exit(m.Run())
}
