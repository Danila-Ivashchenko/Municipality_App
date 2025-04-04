package location

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createLocationQuery = `INSERT INTO location (address, latitude, longitude, geometry) VALUES ($1, $2, $3, $4) RETURNING id`
	updateLocationQuery = `UPDATE location SET address = $1, latitude = $2, longitude = $3, geometry = $4 WHERE id = $5`
	selectLocationQuery = `SELECT id, address, latitude, longitude, geometry FROM location `
	deleteLocationQuery = "DELETE FROM location WHERE id = $1"
)

func (r *locationRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *locationRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Location, error) {
	return r.fetchRow(ctx, selectLocationQuery+" WHERE "+condition, args...)
}

func (r *locationRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Location, error) {
	m := &modelLocation{}

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

func (r *locationRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Location, error) {
	return r.fetchRows(ctx, selectLocationQuery+" WHERE "+condition, args...)
}

func (r *locationRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Location, error) {
	var (
		result []entity.Location
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &modelLocation{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *modelLocation, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Address,
		&m.Latitude,
		&m.Longitude,
		&m.Geometry,
	)
}
