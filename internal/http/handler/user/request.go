package user

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type changeUserPasswordRequest struct {
	Password *string `json:"password"`
}

func (r *changeUserPasswordRequest) Validate() error {
	if r.Password == nil {
		if r.Password == nil {
			return errors.New("no password")
		}
	}

	return nil
}

func (r *changeUserPasswordRequest) Convert(userID int64) *service.ChangeUserPasswordData {
	return &service.ChangeUserPasswordData{
		UserID:   userID,
		Password: *r.Password,
	}
}

type registerUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`

	Password string `json:"password"`
}

func (r *registerUserRequest) Convert() *service.CreateUserData {
	return &service.CreateUserData{
		Email:    r.Email,
		Name:     r.Name,
		LastName: r.LastName,
		Password: r.Password,
	}
}

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *loginUserRequest) Convert() *service.UserLoginData {
	return &service.UserLoginData{
		Email:    r.Email,
		Password: r.Password,
	}
}

type updateUserRequest struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
}

func (r *updateUserRequest) Convert(userID int64) *service.UpdateUserData {
	return &service.UpdateUserData{
		ID:       userID,
		Email:    r.Email,
		Name:     r.Name,
		LastName: r.LastName,
	}
}
