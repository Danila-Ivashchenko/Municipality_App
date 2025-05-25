package region

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createRegionQuery = `INSERT INTO region (name, code) VALUES ($1, $2)`
	selectRegionQuery = `SELECT id, name, code FROM region `
	deleteRegionQuery = "DELETE FROM region"
)

func (r *regionRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *regionRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Region, error) {
	return r.fetchRow(ctx, selectRegionQuery+" WHERE "+condition, args...)
}

func (r *regionRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Region, error) {
	m := &regionModel{}

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

func (r *regionRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Region, error) {
	return r.fetchRows(ctx, selectRegionQuery+" WHERE "+condition, args...)
}

func (r *regionRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Region, error) {
	var (
		result []entity.Region
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &regionModel{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *regionModel, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.Code,
	)
}
