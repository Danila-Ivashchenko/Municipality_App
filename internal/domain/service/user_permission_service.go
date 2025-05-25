package service

import (
	"context"
	"municipality_app/internal/common/validator/field"
	"municipality_app/internal/common/validator/validator"
	"municipality_app/internal/domain/entity"
)

type UserPermissionService interface {
	SetPermissionsToUser(ctx context.Context, data *SetUserPermissionsData) ([]entity.Permission, error)
	GetUserPermissions(ctx context.Context, id int64) ([]entity.Permission, error)
}

type SetUserPermissionsData struct {
	UserID      int64
	Permissions []entity.Permission
}

func (d *SetUserPermissionsData) Validate() error {
	val := validator.Validator{}

	val.AddField(
		field.NewInt64Field("Идентификатор пользователя", d.UserID).Required(),
	)

	return val.Validate()
}
