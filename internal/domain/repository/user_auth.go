package repository

import (
	"context"
	"municipality_app/internal/domain/entity"
	"time"
)

type UserAuthRepository interface {
	CreateUserToken(ctx context.Context, data *CreateUserTokenData) (*entity.UserAuthToken, error)

	DeleteUserTokenByID(ctx context.Context, id int64) error
	DeleteAllUserTokens(ctx context.Context, userID string) error

	UpdateExpireTime(ctx context.Context, userID string, newExpireTime time.Time) error

	GetByToken(ctx context.Context, token string) (*entity.UserAuthToken, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.UserAuthToken, error)
}

type CreateUserTokenData struct {
	Token  string
	UserID int64

	ExpireAt time.Time
}
