package passport_file

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createPassportFileQuery = `INSERT INTO municipality_passport_file (path, passport_id, file_name, created_at) VALUES ($1, $2, $3, $4) returning id;`
	selectPassportFileQuery = `SELECT id, path, passport_id, file_name, created_at FROM municipality_passport_file `
	deletePassportFileQuery = "DELETE FROM municipality_passport_file WHERE id = $1"
)

func (r *passportFileRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
	_, err := r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *passportFileRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.PassportFile, error) {
	return r.fetchRow(ctx, selectPassportFileQuery+" WHERE "+condition, args...)
}

func (r *passportFileRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.PassportFile, error) {
	m := &passportFileModel{}

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

func (r *passportFileRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.PassportFile, error) {
	return r.fetchRows(ctx, selectPassportFileQuery+" WHERE "+condition, args...)
}

func (r *passportFileRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.PassportFile, error) {
	var (
		result []entity.PassportFile
	)

	rows, err := r.db.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &passportFileModel{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *passportFileModel, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.Path,
		&m.PassportID,
		&m.FileName,
		&m.CreateAt,
	)
}
