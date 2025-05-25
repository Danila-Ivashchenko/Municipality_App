package entity_attribute_value

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createQuery = `INSERT INTO municipality_entity_attribute_value (entity_attribute_id, entity_id, value) VALUES ($1, $2, $3) returning id;`
	updateQuery = `UPDATE municipality_entity_attribute_value SET value = $1 WHERE id = $2`
	selectQuery = `SELECT id, entity_attribute_id, entity_id, value FROM municipality_entity_attribute_value `
	deleteQuery = "DELETE FROM municipality_entity_attribute WHERE id = $1"
)

func (r *repo) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *repo) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.EntityAttributeValue, error) {
	return r.fetchRow(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *repo) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.EntityAttributeValue, error) {
	m := &model{}

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

func (r *repo) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.EntityAttributeValue, error) {
	return r.fetchRows(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *repo) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.EntityAttributeValue, error) {
	var (
		result []entity.EntityAttributeValue
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &model{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *model, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.EntityAttributeID,
		&m.EntityID,
		&m.Value,
	)
}
