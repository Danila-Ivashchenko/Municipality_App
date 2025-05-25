package object

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createObjectQuery = `INSERT INTO municipality_object (name, municipality_object_template_id, location_id, description) VALUES ($1, $2, $3, $4);`
	updateObjectQuery = `UPDATE municipality_object SET name = $1, description = $2, location_id = $3 WHERE id = $4;`
	selectObjectQuery = `SELECT id, name, municipality_object_template_id, location_id, description FROM municipality_object `
	deleteObjectQuery = "DELETE FROM municipality_object WHERE id = $1"
)

func (r *objectTemplateRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *objectTemplateRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Object, error) {
	return r.fetchRow(ctx, selectObjectQuery+" WHERE "+condition, args...)
}

func (r *objectTemplateRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Object, error) {
	m := &modelObject{}

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

func (r *objectTemplateRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Object, error) {
	return r.fetchRows(ctx, selectObjectQuery+" WHERE "+condition, args...)
}

func (r *objectTemplateRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Object, error) {
	var (
		result []entity.Object
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &modelObject{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *modelObject, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.ObjectTemplateID,
		&m.LocationID,
		&m.Description,
	)
}
