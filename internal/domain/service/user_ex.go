package service

import (
	"context"
	"municipality_app/internal/domain/entity"
)

type UserExService interface {
	AuthUserByToken(ctx context.Context, token string) (*entity.User, error)

	LogoutUserByToken(ctx context.Context, token string) error
	LogoutUserByID(ctx context.Context, userID string) error
}
