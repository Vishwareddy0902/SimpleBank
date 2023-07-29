package main

import (
	"database/sql"
	"log"

	"github.com/Vishwareddy0902/Simple-Bank/api"
	db "github.com/Vishwareddy0902/Simple-Bank/db/sqlc"
	"github.com/Vishwareddy0902/Simple-Bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to load the config file")
	}
	dB, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Fatal("Unable to connect to DB", err)
	}
	store := db.NewStore(dB)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	err = server.StartServer(config.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("Unable to start the server", err)
	}
}
