package sql_handler

import (
	"context"
	"database/sql"
	"municipality_app/internal/common/context_paylod_parser"
)

type Handler interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) Handler {
	return &handler{db: db}
}

func (h *handler) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	tx := context_paylod_parser.GetTransactionFromContext(ctx)
	if tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}

	return h.db.ExecContext(ctx, query, args...)
}

func (h *handler) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	tx := context_paylod_parser.GetTransactionFromContext(ctx)
	if tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return h.db.QueryRowContext(ctx, query, args...)
}

func (h *handler) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	tx := context_paylod_parser.GetTransactionFromContext(ctx)
	if tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}

	return h.db.QueryContext(ctx, query, args...)
}
