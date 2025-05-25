package service

import (
	"context"
	"municipality_app/internal/common/validator/field"
	"municipality_app/internal/common/validator/validator"
	"municipality_app/internal/domain/entity"
)

type UserService interface {
	RegisterUser(ctx context.Context, data *CreateUserData) (*entity.User, error)
	Login(ctx context.Context, data *UserLoginData) (*entity.UserAuthToken, error)
	Logout(ctx context.Context, userToken *entity.UserAuthToken) error

	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)

	BlockUserByID(ctx context.Context, id int64, isBlocked bool) error

	ChangeUserIsAdmin(ctx context.Context, id int64, isAdmin bool) error
	ChangeUserPassword(ctx context.Context, data *ChangeUserPasswordData) error
	UpdateUser(ctx context.Context, data *UpdateUserData) (*entity.User, error)

	Me(ctx context.Context, id int64) (*entity.UserEx, error)

	GetAll(ctx context.Context) ([]entity.User, error)
}

type CreateUserData struct {
	Login    string
	Name     string
	LastName string
	Email    string
	IsAdmin  bool
	IsBlock  bool

	Password string
}

func (d *CreateUserData) Validate() error {
	v := validator.Validator{}

	v.AddField(
		field.NewStringField("Email", d.Email).Required().Bigger(4),
		field.NewStringField("Пароль", d.Password).Required().Bigger(4),
	)

	return v.Validate()
}

type UpdateUserData struct {
	ID       int64
	Name     *string
	LastName *string
	Email    *string
}

func (d *UpdateUserData) Validate() error {
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

type ChangeUserPasswordData struct {
	UserID int64

	Password string
}

func (d *ChangeUserPasswordData) Validate() error {
	v := validator.Validator{}

	v.AddField(
		field.NewInt64Field("Идентификатор пользователя", d.UserID).Required(),
		field.NewStringField("Пароль", d.Password).Required().Bigger(4),
	)

	return v.Validate()
}

type UserLoginData struct {
	Email    string
	Password string
}

func (d *UserLoginData) Validate() error {
	v := validator.Validator{}

	v.AddField(
		field.NewStringField("Email", d.Email).Required().Bigger(4),
		field.NewStringField("Пароль", d.Password).Required().Bigger(4),
	)

	return v.Validate()
}

type UserLogoutData struct {
	Token string
}
