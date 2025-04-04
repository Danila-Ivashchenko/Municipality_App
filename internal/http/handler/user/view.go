package user

import (
	"municipality_app/internal/domain/entity"
	"time"
)

type userView struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`

	IsAdmin   bool `json:"is_admin"`
	IsBlocked bool `json:"is_blocked"`

	CreatedAt time.Time `json:"created_at"`
}

func newUserView(i *entity.User) *userView {
	return &userView{
		ID:       i.ID,
		Email:    i.Email,
		Name:     i.Name,
		LastName: i.LastName,

		IsAdmin:   i.IsAdmin,
		IsBlocked: i.IsBlocked,

		CreatedAt: i.CreatedAt,
	}
}

type userTokenView struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}

func newUserTokenView(i *entity.UserAuthToken) *userTokenView {
	return &userTokenView{
		Token:    i.Token,
		ExpireAt: i.ExpiresAt,
	}
}
