package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type UserPermissionRepository interface {
	SetPermissionsToUser(ctx context.Context, data *SetUserPermissionsData) ([]entity.Permission, error)
	GetUserPermissions(ctx context.Context, id int64) ([]entity.Permission, error)
	DeleteUserPermissions(ctx context.Context, data *DeleteUserPermissionsData) error
}

type SetUserPermissionsData struct {
	UserID      int64
	Permissions []entity.Permission
}

type DeleteUserPermissionsData struct {
	UserID      int64
	Permissions []entity.Permission
}
