package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yuqihu1103/SimpleBank/api"
	db "github.com/yuqihu1103/SimpleBank/db/sqlc"
	"github.com/yuqihu1103/SimpleBank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot connect to server:", err)
	}
}
