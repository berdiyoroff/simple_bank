package main

import (
	"log"

	"github.com/berdiyoroff/simple_bank/api"
	db "github.com/berdiyoroff/simple_bank/db/sqlc"
	"github.com/berdiyoroff/simple_bank/pkg/database/postgres"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := postgres.NewPool(dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	err = server.Start(":8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
