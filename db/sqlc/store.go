package sqlc

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
	q    Querier
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		pool: pool,
		q:    New(pool),
	}
}
