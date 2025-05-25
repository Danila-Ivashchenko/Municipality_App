package service

import (
	"context"
	"municipality_app/internal/common/validator/field"
	"municipality_app/internal/common/validator/validator"
	"municipality_app/internal/domain/entity"
)

type UserAdminService interface {
	GetAll(ctx context.Context) ([]entity.UserEx, error)
	GetByID(ctx context.Context, userID int64) (*entity.UserEx, error)
	Update(ctx context.Context, data *UpdateUserByAdminData) (*entity.UserEx, error)
}

type UpdateUserByAdminData struct {
	ID        int64
	Name      *string
	LastName  *string
	Email     *string
	IsAdmin   *bool
	IsBlocked *bool

	Permissions *[]entity.Permission
}

func (d *UpdateUserByAdminData) Validate() error {
	v := validator.Validator{}

	v.AddField(
		field.NewInt64Field("Идентификатор пользователя", d.ID).Required(),
	)

	if d.Name != nil {
		v.AddField(
			field.NewStringField("Имя", *d.Name).Required().Bigger(4),
		)
	}

	if d.LastName != nil {
		v.AddField(
			field.NewStringField("Фамилия", *d.LastName).Required().Bigger(4),
		)
	}

	if d.Email != nil {
		v.AddField(
			field.NewStringField("Email", *d.Email).Required().Bigger(4),
		)
	}

	return v.Validate()
}
