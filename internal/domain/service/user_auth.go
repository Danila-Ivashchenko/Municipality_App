package service

import (
	"context"
	"municipality_app/internal/domain/entity"
	"time"
)

type UserAuthService interface {
	NewUserToken(ctx context.Context, userID int64) (*entity.UserAuthToken, error)

	DeleteUserToken(ctx context.Context, userToken *entity.UserAuthToken) error
	DeleteAllUserTokens(ctx context.Context, userID string) error

	UpdateExpireTime(ctx context.Context, userID string, timeDelta time.Duration) error

	GetByToken(ctx context.Context, token string) (*entity.UserAuthToken, error)
	GetByTokenWithValidation(ctx context.Context, token string) (*entity.UserAuthToken, error)
	GetByUserID(ctx context.Context, userID int64) ([]entity.UserAuthToken, error)

	DeleteExtraTokens(ctx context.Context, userID int64) error
}
