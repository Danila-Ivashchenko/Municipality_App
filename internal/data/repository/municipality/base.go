package municipality

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createMunicipality     = `INSERT INTO municipality (name, region_id) VALUES ($1, $2)`
	selectMunicipality     = `SELECT id, name, region_id, is_hidden, created_at FROM municipality `
	updateMunicipalityByID = `UPDATE municipality SET name = $1, region_id = $2, is_hidden = $3 WHERE id = $4`
	deleteMunicipalityByID = "DELETE FROM municipality WHERE id = $1"
)

func (r *municipalityRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *municipalityRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Municipality, error) {
	return r.fetchRow(ctx, selectMunicipality+" WHERE "+condition, args...)
}

func (r *municipalityRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Municipality, error) {
	m := &municipalityModel{}

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

func (r *municipalityRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Municipality, error) {
	return r.fetchRows(ctx, selectMunicipality+" WHERE "+condition, args...)
}

func (r *municipalityRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Municipality, error) {
	var (
		result []entity.Municipality
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &municipalityModel{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *municipalityModel, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.RegionID,
		&m.IsHidden,
		&m.CreatedAt,
	)
}
