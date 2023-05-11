package main

import (
	"database/sql"
	"log"

	"github.com/Vishwareddy0902/Simple-Bank/api"
	db "github.com/Vishwareddy0902/Simple-Bank/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable"
	serverAddres = "0.0.0.0:8080"
)

func main() {
	dB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Unable to connect to DB", err)
	}
	store := db.NewStore(dB)
	server := api.NewServer(store)
	err = server.StartServer(serverAddres)
	if err != nil {
		log.Fatal("Unable to start the server", err)
	}
}
