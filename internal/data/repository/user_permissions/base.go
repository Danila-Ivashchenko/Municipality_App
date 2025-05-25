package user_permissions

import (
	"context"
	"database/sql"
	"errors"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

const (
	createUserPermissionQuery = `INSERT INTO user_permission (user_id, permission) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	selectUserPermissionQuery = `SELECT user_id, permission FROM user_permission `
	deleteUserPermissionQuery = "DELETE FROM user_permission WHERE user_id = $1 AND permission = ANY($2)"
)

func (r *userPermissionRepository) execQuery(ctx context.Context, sqlQuery string, args ...any) error {
	_, err := r.handler.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userPermissionRepository) fetchRowWithCondition(ctx context.Context, condition string, args ...any) (*entity.Permission, error) {
	return r.fetchRow(ctx, selectUserPermissionQuery+" WHERE "+condition, args...)
}

func (r *userPermissionRepository) fetchRow(ctx context.Context, sqlQuery string, args ...any) (*entity.Permission, error) {
	m := &userPermissionModel{}

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

	val := m.convert()

	return &val, nil
}

func (r *userPermissionRepository) fetchRowsWithCondition(ctx context.Context, condition string, args ...any) ([]entity.Permission, error) {
	return r.fetchRows(ctx, selectUserPermissionQuery+" WHERE "+condition, args...)
}

func (r *userPermissionRepository) fetchRows(ctx context.Context, sqlQuery string, args ...any) ([]entity.Permission, error) {
	var (
		result []entity.Permission
	)

	rows, err := r.handler.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			m = &userPermissionModel{}
		)

		err = scan(m, rows)
		if err != nil {
			return nil, err
		}

		result = append(result, m.convert())
	}

	return result, nil
}

func scan(m *userPermissionModel, row sql_common.RowScanner) error {
	return row.Scan(
		&m.UserID,
		&m.Permission,
	)
}
