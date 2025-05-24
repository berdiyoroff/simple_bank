package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Option -.
type Option func(*pgxpool.Config)

// MaxPoolSize -.
func MaxPoolSize(size int) Option {
	return func(c *pgxpool.Config) {
		c.MaxConns = int32(size)
	}
}

// MaxConnIdleTime -.
func MaxConnIdleTime(time time.Duration) Option {
	return func(c *pgxpool.Config) {
		c.MaxConnIdleTime = time
	}
}

// MaxConnLifeTime -.
func MaxConnLifeTime(time time.Duration) Option {
	return func(c *pgxpool.Config) {
		c.MaxConnLifetime = time
	}
}

// AfterConnect -.
func AfterConnect(f func(ctx context.Context, conn *pgx.Conn) error) Option {
	return func(c *pgxpool.Config) {
		c.AfterConnect = f
	}
}

// Other ...
