package entity_template

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createEntityQuery = `INSERT INTO municipality_entity_template (name, entity_type_id, municipality_id) VALUES ($1, $2, $3);`
	updateEntityQuery = `UPDATE municipality_entity_template SET name = $1, entity_type_id = $2 WHERE id = $3`
	selectEntityQuery = `SELECT id, name, entity_type_id, municipality_id FROM municipality_entity_template `
	deleteEntityQuery = "DELETE FROM municipality_entity_template WHERE id = $1"
)

func (r *entityTemplateRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *entityTemplateRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.EntityTemplate, error) {
	return r.fetchRow(ctx, selectEntityQuery+" WHERE "+condition, args...)
}

func (r *entityTemplateRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.EntityTemplate, error) {
	m := &modelEntityTemplate{}

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

func (r *entityTemplateRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.EntityTemplate, error) {
	return r.fetchRows(ctx, selectEntityQuery+" WHERE "+condition, args...)
}

func (r *entityTemplateRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.EntityTemplate, error) {
	var (
		result []entity.EntityTemplate
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &modelEntityTemplate{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *modelEntityTemplate, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.EntityTypeID,
		&m.MunicipalityID,
	)
}
