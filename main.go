package main

import (
	"database/sql"
	"log"
	"sb/api"
	db "sb/db/sqlc"
	"sb/db/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil{
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, _ := api.NewServer(config, store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
