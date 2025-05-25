package user_permissions

import (
	"context"
	sql_common "municipality_app/internal/common/data/sql"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
)

type userPermissionRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.UserPermissionRepository {
	repo := &userPermissionRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *userPermissionRepository) SetPermissionsToUser(ctx context.Context, data *repository.SetUserPermissionsData) ([]entity.Permission, error) {
	if len(data.Permissions) == 0 {
		return nil, nil
	}

	permissionModels := make([]userPermissionModel, 0)

	for _, p := range data.Permissions {
		permissionModels = append(permissionModels, *newUserPermissionModel(data.UserID, p))
	}

	for _, p := range permissionModels {
		err := r.execQuery(ctx, createUserPermissionQuery, p.UserID, p.Permission)
		if err != nil {
			return nil, err
		}
	}

	return data.Permissions, nil
}

func (r *userPermissionRepository) GetUserPermissions(ctx context.Context, id int64) ([]entity.Permission, error) {
	return r.fetchRowsWithCondition(ctx, "user_id = $1", id)
}

func (r *userPermissionRepository) DeleteUserPermissions(ctx context.Context, data *repository.DeleteUserPermissionsData) error {
	if len(data.Permissions) == 0 {
		return nil
	}

	permissions := make([]int64, 0)

	for _, p := range data.Permissions {
		permissions = append(permissions, int64(p.ToUint8()))
	}

	return r.execQuery(ctx, deleteUserPermissionQuery, data.UserID, sql_common.NewNullInt64Array(permissions))
}
