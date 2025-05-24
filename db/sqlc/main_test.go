package sqlc

import (
	"log"
	"os"
	"testing"

	"github.com/berdiyoroff/simple_bank/pkg/database/postgres"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testStore *Store

func TestMain(m *testing.M) {
	testPool, err := postgres.NewPool(dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testStore = NewStore(testPool)

	os.Exit(m.Run())
}
