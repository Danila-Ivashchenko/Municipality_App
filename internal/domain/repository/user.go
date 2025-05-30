package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, data *CreateUserData) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)

	GetUserByID(ctx context.Context, id int64) (*entity.User, error)
	GetUserFullByID(ctx context.Context, id int64) (*entity.UserFull, error)

	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserFullByEmail(ctx context.Context, email string) (*entity.UserFull, error)

	ChangeUserPassword(ctx context.Context, userID int64, password string) error

	GetAllUsers(ctx context.Context) ([]entity.User, error)

	ChangeUserBlocked(ctx context.Context, userID int64, isBlocked bool) error
	ChangeUserAdmin(ctx context.Context, userID int64, isBlocked bool) error
}

type CreateUserData struct {
	Name     string
	LastName string
	Email    string

	IsAdmin   bool
	IsBlocked bool

	Password string
}
