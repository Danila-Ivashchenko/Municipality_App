package user_auth_token

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createUserAuthTokenQuery = `INSERT INTO user_auth_token (user_id, token, expires_at) VALUES ($1, $2, $3)`
	selectUserAuthTokenQuery = `SELECT id, user_id, token, expires_at, created_at FROM user_auth_token `
	deleteUserAuthToken      = "DELETE FROM user_auth_token"
)

func (r *userAuthRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *userAuthRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.UserAuthToken, error) {
	return r.fetchRow(ctx, selectUserAuthTokenQuery+" WHERE "+condition, args...)
}

func (r *userAuthRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.UserAuthToken, error) {
	m := &userAuthTokenModel{}

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

func (r *userAuthRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.UserAuthToken, error) {
	return r.fetchRows(ctx, selectUserAuthTokenQuery+" WHERE "+condition, args...)
}

func (r *userAuthRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.UserAuthToken, error) {
	var (
		result []entity.UserAuthToken
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &userAuthTokenModel{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, *m.convert())
	}

	return result, nil
}

func scan(m *userAuthTokenModel, row sql_common.RowScanner) error {
	return row.Scan(
		&m.ID,
		&m.UserID,
		&m.Token,
		&m.ExpiresAt,
		&m.CreatedAt,
	)
}
