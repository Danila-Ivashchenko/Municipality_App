package object_template

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createObjectQuery = `INSERT INTO municipality_object_template (name, object_type_id, municipality_id) VALUES ($1, $2, $3);`
	updateObjectQuery = `UPDATE municipality_object_template SET name = $1, object_type_id = $2 WHERE id = $3`
	selectObjectQuery = `SELECT id, name, object_type_id, municipality_id FROM municipality_object_template `
	deleteObjectQuery = "DELETE FROM municipality_object_template WHERE id = $1"
)

func (r *objectTemplateRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *objectTemplateRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.ObjectTemplate, error) {
	return r.fetchRow(ctx, selectObjectQuery+" WHERE "+condition, args...)
}

func (r *objectTemplateRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.ObjectTemplate, error) {
	m := &modelObjectTemplate{}

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

func (r *objectTemplateRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.ObjectTemplate, error) {
	return r.fetchRows(ctx, selectObjectQuery+" WHERE "+condition, args...)
}

func (r *objectTemplateRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.ObjectTemplate, error) {
	var (
		result []entity.ObjectTemplate
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &modelObjectTemplate{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *modelObjectTemplate, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.ObjectTypeID,
		&m.MunicipalityID,
	)
}
