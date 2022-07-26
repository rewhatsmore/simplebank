package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rewhatsmore/simplebank/api"
	db "github.com/rewhatsmore/simplebank/db/sqlc"
	"github.com/rewhatsmore/simplebank/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server:", err)
	}

}
