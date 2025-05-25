package object_type

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createObjectQuery = `INSERT INTO municipality_object_type (name) VALUES ($1)`
	updateObjectQuery = `UPDATE municipality_object_type SET name = $1 WHERE id = $2`
	selectObjectQuery = `SELECT id, name FROM municipality_object_type `
	deleteObjectQuery = "DELETE FROM municipality_object_type WHERE id = $1"
)

func (r *objectRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *objectRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.ObjectType, error) {
	return r.fetchRow(ctx, selectObjectQuery+" WHERE "+condition, args...)
}

func (r *objectRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.ObjectType, error) {
	m := &modelObjectType{}

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

func (r *objectRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.ObjectType, error) {
	return r.fetchRows(ctx, selectObjectQuery+" WHERE "+condition, args...)
}

func (r *objectRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.ObjectType, error) {
	var (
		result []entity.ObjectType
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &modelObjectType{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *modelObjectType, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
	)
}
