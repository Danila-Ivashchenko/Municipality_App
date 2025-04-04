package object_attribute_value

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createQuery = `INSERT INTO municipality_object_attribute_value (object_attribute_id, object_id, value) VALUES ($1, $2, $3) returning id;`
	updateQuery = `UPDATE municipality_object_attribute_value SET value = $1 WHERE id = $2`
	selectQuery = `SELECT id, object_attribute_id, object_id, value FROM municipality_object_attribute_value `
	deleteQuery = "DELETE FROM municipality_object_attribute WHERE id = $1"
)

func (r *repo) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
	res, err := r.db.ExecContext(ctx, sqlQuery, args...)
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

func (r *repo) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.ObjectAttributeValue, error) {
	return r.fetchRow(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *repo) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.ObjectAttributeValue, error) {
	m := &model{}

	row := r.db.QueryRowContext(ctx, sqlQuery, args...)
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

func (r *repo) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.ObjectAttributeValue, error) {
	return r.fetchRows(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *repo) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.ObjectAttributeValue, error) {
	var (
		result []entity.ObjectAttributeValue
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
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
		&m.ObjectAttributeID,
		&m.ObjectID,
		&m.Value,
	)
}
