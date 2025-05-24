package sqlc

import (
	"context"
	"fmt"
)

var txContextKey = struct{}{}

func (s *Store) execTx(ctx context.Context, fn func(context.Context) error) error {
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

func (s *Store) querier(ctx context.Context) Querier {
	querier, ok := ctx.Value(txContextKey).(Querier)
	if ok {
		return querier
	}
	return s.q
}
