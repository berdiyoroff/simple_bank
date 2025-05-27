package main

import (
	"log"

	"github.com/berdiyoroff/simple_bank/api"
	"github.com/berdiyoroff/simple_bank/config"
	db "github.com/berdiyoroff/simple_bank/db/sqlc"
	"github.com/berdiyoroff/simple_bank/pkg/database/postgres"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := postgres.NewPool(config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
