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

type updateUserRequest struct {
	Email    *string `json:"email"`
	Name     *string `json:"name"`
	LastName *string `json:"last_name"`
}

func (r *updateUserRequest) Validate() error {
	var (
		updated bool
	)

	if r.Email != nil {
		updated = true
		if len(*r.Email) < 3 {
			return errors.New("email too small")
		}
	}
	if r.Name != nil {
		updated = true
		if len(*r.Name) < 3 {
			return errors.New("email too small")
		}
	}
	if r.LastName != nil {
		updated = true
		if len(*r.LastName) < 3 {
			return errors.New("email too small")
		}
	}

	if !updated {
		return errors.New("no update")
	}

	return nil
}

func (r *updateUserRequest) Convert(userID int64) *service.UpdateUserData {
	return &service.UpdateUserData{
		ID:       userID,
		Email:    r.Email,
		Name:     r.Name,
		LastName: r.LastName,
	}
}
