package user_permissions

import (
	"database/sql"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
)

type userPermissionModel struct {
	UserID     sql.NullInt64
	Permission sql.NullInt64
}

func newUserPermissionModel(userID int64, permission entity.Permission) *userPermissionModel {
	return &userPermissionModel{
		UserID:     sql_common.NewNullInt64(userID),
		Permission: sql_common.NewNullInt64(int64(permission.ToUint8())),
	}
}

func (m *userPermissionModel) convert() entity.Permission {
	return entity.PermissionFromUint8(uint8(m.Permission.Int64))
}
