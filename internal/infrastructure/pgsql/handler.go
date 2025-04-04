package pgsql

import (
	"context"
	"database/sql"
	"municipality_app/internal/common/config"
	"municipality_app/internal/infrastructure/db"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type dbManager struct {
	db *sql.DB
}

func NewPsqlManager(config *config.Config) (db.DataBaseManager, error) {
	db, err := sql.Open("pgx", config.GetPsqlURL())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(5)
	return &dbManager{
		db: db,
	}, nil
}

func (dbm dbManager) GetConnection(ctx context.Context) (*sql.Conn, error) {
	return dbm.db.Conn(ctx)
}

func (dbm dbManager) GetDB() *sql.DB {
	return dbm.db
}

func GetConnectionsPull(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}
	_, err = pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
