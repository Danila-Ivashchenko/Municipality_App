package user

import (
	"errors"
	"municipality_app/internal/domain/service"
)

type registerUserRequest struct {
	Email    *string `json:"email"`
	Name     string  `json:"name"`
	LastName string  `json:"last_name"`

	Password *string `json:"password"`
}

func (r *registerUserRequest) Validate() error {
	if r.Email == nil {
		return errors.New("no email")
	}
	if r.Password == nil {
		return errors.New("no password")
	}

	return nil
}

func (r *registerUserRequest) Convert() *service.CreateUserData {
	return &service.CreateUserData{
		Email:    *r.Email,
		Name:     r.Name,
		LastName: r.LastName,
		Password: *r.Password,
	}
}

type loginUserRequest struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (r *loginUserRequest) Validate() error {
	if r.Email == nil {
		return errors.New("no email")
	}
	if r.Password == nil {
		return errors.New("no password")
	}

	return nil
}

func (r *loginUserRequest) Convert() *service.UserLoginData {
	return &service.UserLoginData{
		Email:    *r.Email,
		Password: *r.Password,
	}
}
