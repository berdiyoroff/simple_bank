package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	TranferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SqlStore struct {
	*Queries
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) Store {
	return &SqlStore{
		Queries: New(pool),
		pool:    pool,
	}
}

var txContextKey = struct{}{}

func (s *SqlStore) execTx(ctx context.Context, fn func(context.Context) error) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}

	txQuery := New(tx)
	vCtx := context.WithValue(ctx, txContextKey, txQuery)

	err = fn(vCtx)
	if err != nil {
		if rbErr := tx.Rollback(vCtx); rbErr != nil {
			return fmt.Errorf("tx error: %w; tx rollback error: %w", err, rbErr)
		}
		return fmt.Errorf("tx error: %w", err)
	}

	return tx.Commit(vCtx)
}

func (s *SqlStore) querier(ctx context.Context) Querier {
	querier, ok := ctx.Value(txContextKey).(Querier)
	if ok {
		return querier
	}
	return s.Queries
}
