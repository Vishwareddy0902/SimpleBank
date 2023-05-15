package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Vishwareddy0902/Simple-Bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Unable to load the config file")
	}
	testDB, err = sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("Unable to connect to DB", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
