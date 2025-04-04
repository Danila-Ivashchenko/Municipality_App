package user

import (
	"context"
	"database/sql"
	"errors"
	"municipality_app/internal/domain/entity"
)

const (
	createUserQuery     = `INSERT INTO users (email, name, last_name, password) VALUES ($1, $2, $3, $4) RETURNING id`
	selectUserQuery     = `SELECT id, email, name, last_name, is_admin, is_blocked, created_at FROM users `
	selectUserFullQuery = `SELECT id, email, name, last_name, is_admin, is_blocked, password, created_at FROM users `
)

func (r *userRepository) exexUserQuery(ctx context.Context, sqlQuery string, args ...any) error {
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

func (r *userRepository) fetchUserWithСondition(ctx context.Context, condition string, args ...any) (*entity.User, error) {
	return r.fetchUser(ctx, selectUserQuery+" WHERE "+condition, args...)
}

func (r *userRepository) fetchUser(ctx context.Context, sqlQuery string, args ...any) (*entity.User, error) {
	m := &userModel{}

	row := r.db.QueryRowContext(ctx, sqlQuery, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&m.ID,
		&m.Email,
		&m.Name,
		&m.LastName,
		&m.IsAdmin,
		&m.IsBlocked,
		&m.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return m.convert(), nil
}

func (r *userRepository) fetchUserFullWithСondition(ctx context.Context, condition string, args ...any) (*entity.UserFull, error) {
	return r.fetchUserFull(ctx, selectUserFullQuery+" WHERE "+condition, args...)
}

func (r *userRepository) fetchUserFull(ctx context.Context, sqlQuery string, args ...any) (*entity.UserFull, error) {
	m := &userFullModel{}

	row := r.db.QueryRowContext(ctx, sqlQuery, args...)
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(
		&m.ID,
		&m.Email,
		&m.Name,
		&m.LastName,
		&m.IsAdmin,
		&m.IsBlocked,
		&m.Password,
		&m.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return m.convert(), nil
}
