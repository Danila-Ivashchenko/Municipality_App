package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type UserService interface {
	RegisterUser(ctx context.Context, data *CreateUserData) (*entity.User, error)
	Login(ctx context.Context, data *UserLoginData) (*entity.UserAuthToken, error)
	Logout(ctx context.Context, userToken *entity.UserAuthToken) error

	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByID(ctx context.Context, id int64) (*entity.User, error)

	BlockUserByID(ctx context.Context, id int) error

	ChangeUserIsAdmin(ctx context.Context, id int) error
	ChangeUserPassword(ctx context.Context, data ChangeUserPasswordData) error
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

type ChangeUserPasswordData struct {
	Login string
	Email string

	Password        string
	PasswordConfirm string
}

type UserLoginData struct {
	Email    string
	Password string
}

type UserLogoutData struct {
	Token string
}
