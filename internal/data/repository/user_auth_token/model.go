package user_auth_token

import (
	"database/sql"
	sql_convertor "municipality_app/internal/common/data/sql"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
)

type userAuthTokenModel struct {
	ID     sql.NullInt64
	Token  sql.NullString
	UserID sql.NullInt64

	CreatedAt sql.NullTime
	ExpiresAt sql.NullTime
}

func (m *userAuthTokenModel) convert() *entity.UserAuthToken {
	return &entity.UserAuthToken{
		ID:        m.ID.Int64,
		Token:     m.Token.String,
		UserID:    m.UserID.Int64,
		CreatedAt: m.CreatedAt.Time,
		ExpiresAt: m.ExpiresAt.Time,
	}
}

func newUserAuthTokenModelFromCreateData(data *repository.CreateUserTokenData) *userAuthTokenModel {
	return &userAuthTokenModel{
		Token:  sql_convertor.NewNullString(data.Token),
		UserID: sql_convertor.NewNullInt64(data.UserID),

		ExpiresAt: sql_convertor.NewNullTime(data.ExpireAt),
	}
}

func newUserAuthTokenModel(i *entity.UserAuthToken) *userAuthTokenModel {
	return &userAuthTokenModel{
		ID:     sql_convertor.NewNullInt64(i.UserID),
		Token:  sql_convertor.NewNullString(i.Token),
		UserID: sql_convertor.NewNullInt64(i.UserID),

		ExpiresAt: sql_convertor.NewNullTime(i.ExpiresAt),
	}
}
