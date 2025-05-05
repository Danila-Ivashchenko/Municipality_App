package passport

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createPassportQuery          = `INSERT INTO municipality_passport (name, municipality_id, description, year, revision_code, is_main) VALUES ($1, $2, $3, $4, $5, $6)`
	selectPassportQuery          = `SELECT id, name, municipality_id, description, year, revision_code, is_main, is_hidden, updated_at, created_at  FROM municipality_passport `
	updatePassportQuery          = `UPDATE municipality_passport SET name = $1, description = $2, year = $3, is_hidden = $4, updated_at = $5 WHERE id = $6`
	updateUpdatedAtPassportQuery = `UPDATE municipality_passport SET updated_at = $1 WHERE id = $2`
	updateIsMainPassportQuery    = `UPDATE municipality_passport SET is_main = $1, updated_at = $2 `
	deletePassportQuery          = "DELETE FROM municipality_passport WHERE id = $1"
)

func (r *passportRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *passportRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Passport, error) {
	return r.fetchRow(ctx, selectPassportQuery+" WHERE "+condition, args...)
}

func (r *passportRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Passport, error) {
	m := &passportModel{}

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

func (r *passportRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Passport, error) {
	return r.fetchRows(ctx, selectPassportQuery+" WHERE "+condition, args...)
}

func (r *passportRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Passport, error) {
	var (
		result []entity.Passport
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &passportModel{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *passportModel, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Name,
		&m.MunicipalityID,
		&m.Description,
		&m.Year,
		&m.RevisionCode,
		&m.IsMain,
		&m.IsHidden,
		&m.UpdatedAt,
		&m.CreatedAt,
	)
}
