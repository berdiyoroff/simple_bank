package db

import (
	"log"
	"os"
	"testing"

	"github.com/berdiyoroff/simple_bank/config"
	"github.com/berdiyoroff/simple_bank/pkg/database/postgres"
)

var testStore Store

func TestMain(m *testing.M) {

	config, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testPool, err := postgres.NewPool(config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(testPool)

	os.Exit(m.Run())
}
