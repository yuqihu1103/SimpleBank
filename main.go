package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yuqihu1103/SimpleBank/api"
	db "github.com/yuqihu1103/SimpleBank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot connect to server:", err)
	}
}
