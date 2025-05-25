package sql_transactor

import (
	"context"
	"database/sql"
	"errors"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/common/transactor"
	"municipality_app/internal/infrastructure/db"
)

type sqlTransactor struct {
	db *sql.DB
}

func NewTransactor(m db.DataBaseManager) transactor.Transactor {
	return &sqlTransactor{
		db: m.GetDB(),
	}
}

func (t *sqlTransactor) Execute(ctx context.Context, fn func(tx context.Context) error) error {
	var (
		ctxWithTx context.Context
		isChild   = false
		tx        *sql.Tx
		err       error
	)

	tx = context_paylod_parser.GetTransactionFromContext(ctx)
	if tx == nil {
		tx, err = t.db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		ctxWithTx = context_paylod_parser.SetTransactionToContext(ctx, tx)
	} else {
		isChild = true
		ctxWithTx = ctx
	}

	err = fn(ctxWithTx)
	if err != nil {
		if isChild {
			return err
		}
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.New("rollback failed: " + rbErr.Error() + " (original core_errors: " + err.Error() + ")")
		}
		return err
	}

	if !isChild {
		if err = tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}
