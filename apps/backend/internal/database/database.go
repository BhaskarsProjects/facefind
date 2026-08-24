package database

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/bhaskarvb28/facefind/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

const DatabasePingTimeout = 10

type Database struct {
	Pool *pgxpool.Pool
}

func New(cfg *config.Config) (*Database, error) {
	hostPort := net.JoinHostPort(
		cfg.Database.Host,
		strconv.Itoa(cfg.Database.Port),
	)

	encodedPassword := url.QueryEscape(cfg.Database.Password)

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=%s",
		cfg.Database.User,
		encodedPassword,
		hostPort,
		cfg.Database.Name,
		cfg.Database.SSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse database config: %w",
			err,
		)
	}

	pool, err := pgxpool.NewWithConfig(
		context.Background(),
		poolConfig,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to create database pool: %w",
			err,
		)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		DatabasePingTimeout*time.Second,
	)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()

		return nil, fmt.Errorf(
			"failed to ping database: %w",
			err,
		)
	}

	return &Database{
		Pool: pool,
	}, nil
}

func (db *Database) Close() {
	db.Pool.Close()
}
