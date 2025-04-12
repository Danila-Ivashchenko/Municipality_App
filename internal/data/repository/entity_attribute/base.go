package entity_attribute

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createQuery = `INSERT INTO municipality_entity_attribute ( entity_template_id, name, default_value, to_show) VALUES ($1, $2, $3, $4)`
	updateQuery = `UPDATE municipality_entity_attribute SET name = $1, default_value = $2, to_show = $3 WHERE id = $4`
	selectQuery = `SELECT id, entity_template_id, name, default_value, to_show FROM municipality_entity_attribute `
	deleteQuery = "DELETE FROM municipality_entity_attribute WHERE id = $1"
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

func (r *repo) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.EntityAttribute, error) {
	return r.fetchRow(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *repo) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.EntityAttribute, error) {
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

func (r *repo) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.EntityAttribute, error) {
	return r.fetchRows(ctx, selectQuery+" WHERE "+condition, args...)
}

func (r *repo) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.EntityAttribute, error) {
	var (
		result []entity.EntityAttribute
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
		&m.EntityTemplateID,
		&m.Name,
		&m.DefaultValue,
		&m.ToShow,
	)
}
