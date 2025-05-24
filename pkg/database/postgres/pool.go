package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const timeout = 5 * time.Second

func NewPool(psqlURI string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return pgxpool.New(ctx, psqlURI)
}

func NewPoolWithConfig(psqlURI string, opts ...Option) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(psqlURI)
	if err != nil {
		return nil, err
	}

	// Custom options
	for _, opt := range opts {
		opt(config)
	}

	return pgxpool.NewWithConfig(context.Background(), config)
}
