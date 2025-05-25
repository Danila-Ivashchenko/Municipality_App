package user_auth_token

import (
	"context"
	"municipality_app/internal/common/sql_handler"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/infrastructure/db"
	"time"
)

type userAuthRepository struct {
	handler sql_handler.Handler
}

func New(m db.DataBaseManager) repository.UserAuthRepository {
	repo := &userAuthRepository{
		handler: sql_handler.NewHandler(m.GetDB()),
	}
	return repo
}

func (r *userAuthRepository) CreateUserToken(ctx context.Context, data *repository.CreateUserTokenData) (*entity.UserAuthToken, error) {
	model := newUserAuthTokenModelFromCreateData(data)

	err := r.execQuery(ctx, createUserAuthTokenQuery, model.UserID, model.Token, model.ExpiresAt)
	if err != nil {
		return nil, err
	}

	return r.GetByToken(ctx, data.Token)
}

func (r *userAuthRepository) DeleteUserTokenByID(ctx context.Context, id int64) error {
	return r.execQuery(ctx, deleteUserAuthToken+" WHERE id = $1", id)
}

func (r *userAuthRepository) DeleteAllUserTokens(ctx context.Context, userID string) error {
	return r.execQuery(ctx, deleteUserAuthToken+" WHERE user_id = $1", userID)

}

func (r *userAuthRepository) UpdateExpireTime(ctx context.Context, userID string, newExpireTime time.Time) error {
	//TODO implement me
	panic("implement me")
}

func (r *userAuthRepository) GetByToken(ctx context.Context, token string) (*entity.UserAuthToken, error) {
	return r.fetchRowWithCondition(ctx, " token = $1", token)
}

func (r *userAuthRepository) GetByUserID(ctx context.Context, userID int64) ([]entity.UserAuthToken, error) {
	return r.fetchRowsWithCondition(ctx, " user_id = $1 ORDER BY created_at ASC", userID)
}
