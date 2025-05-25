package entity_type

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createEntityQuery = `INSERT INTO municipality_entity_type (name) VALUES ($1)`
	updateEntityQuery = `UPDATE municipality_entity_type SET name = $1 WHERE id = $2`
	selectEntityQuery = `SELECT id, name FROM municipality_entity_type `
	deleteEntityQuery = "DELETE FROM municipality_entity_type WHERE id = $1"
)

func (r *entityRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
	res, err := r.handler.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("now rows affected")
	}

	return nil
}

func (r *entityRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.EntityType, error) {
	return r.fetchRow(ctx, selectEntityQuery+" WHERE "+condition, args...)
}

func (r *entityRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.EntityType, error) {
	m := &modelEntityType{}

	row := r.handler.QueryRowContext(ctx, sqlQuery, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := scan(m, row)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return m.convert(), nil
}

func (r *entityRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.EntityType, error) {
	return r.fetchRows(ctx, selectEntityQuery+" WHERE "+condition, args...)
}

func (r *entityRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.EntityType, error) {
	var (
		result []entity.EntityType
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &modelEntityType{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *modelEntityType, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
	)
}
