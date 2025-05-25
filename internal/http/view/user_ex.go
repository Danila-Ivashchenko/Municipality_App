package view

import (
	"municipality_app/internal/domain/entity"
	"sort"
	"time"
)

type UserExView struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`

	IsAdmin   bool `json:"is_admin"`
	IsBlocked bool `json:"is_blocked"`

	CreatedAt   time.Time `json:"created_at"`
	Permissions []int     `json:"permissions"`
}

func NewUserExViews(data []entity.UserEx) []UserExView {
	var (
		result = make([]UserExView, 0)
	)

	for _, v := range data {
		result = append(result, *NewUserExView(&v))
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func NewUserExView(i *entity.UserEx) *UserExView {
	var (
		permissions []int
	)

	for _, p := range i.Permissions {
		permissions = append(permissions, int(p.ToUint8()))
	}

	return &UserExView{
		ID:       i.User.ID,
		Email:    i.User.Email,
		Name:     i.User.Name,
		LastName: i.User.LastName,

		IsAdmin:   i.User.IsAdmin,
		IsBlocked: i.User.IsBlocked,

		CreatedAt:   i.User.CreatedAt,
		Permissions: permissions,
	}
}
