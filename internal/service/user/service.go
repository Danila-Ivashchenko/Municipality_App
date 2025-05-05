package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"municipality_app/internal/domain/entity"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

func (svc userService) RegisterUser(ctx context.Context, data *service.CreateUserData) (*entity.User, error) {
	userExists, err := svc.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if userExists != nil {
		return nil, errors.New("this email is already registered")
	}

	hashedPassword, err := hashPassword(data.Password)
	if err != nil {
		return nil, err
	}

	crateData := &repository.CreateUserData{
		Name:     data.Name,
		LastName: data.LastName,
		Email:    data.Email,
		Password: hashedPassword,
	}

	user, err := svc.UserRepository.CreateUser(ctx, crateData)

	return user, err
}

func (svc userService) UpdateUser(ctx context.Context, data *service.UpdateUserData) (*entity.User, error) {
	userExists, err := svc.UserRepository.GetUserByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if userExists == nil {
		return nil, errors.New("user no found")
	}

	if data.Name != nil {
		userExists.Name = *data.Name
	}

	if data.LastName != nil {
		userExists.LastName = *data.LastName
	}

	if data.Email != nil && *data.Email != userExists.Email {
		userFull, err := svc.UserRepository.GetUserFullByEmail(ctx, *data.Email)
		if err != nil {
			return nil, err
		}

		if userFull != nil {
			return nil, errors.New("this email is already registered")
		}

		userExists.Email = *data.Email
	}

	user, err := svc.UserRepository.UpdateUser(ctx, userExists)

	return user, err
}

func (svc userService) Login(ctx context.Context, data *service.UserLoginData) (*entity.UserAuthToken, error) {
	userFull, err := svc.UserRepository.GetUserFullByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	compareErr := comparePassword(data.Password, userFull.Password)
	if compareErr != nil {
		return nil, compareErr
	}

	err = svc.UserAuthService.DeleteExtraTokens(ctx, userFull.ID)
	if err != nil {
		return nil, err
	}

	userToken, err := svc.UserAuthService.NewUserToken(ctx, userFull.ID)
	if err != nil {
		return nil, err
	}

	return userToken, nil
}

func (svc userService) Logout(ctx context.Context, userToken *entity.UserAuthToken) error {
	return svc.UserAuthService.DeleteUserToken(ctx, userToken)
}

func (svc userService) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	return svc.UserRepository.GetUserByEmail(ctx, email)
}

func (svc userService) GetUserByID(ctx context.Context, id int64) (*entity.User, error) {
	return svc.UserRepository.GetUserByID(ctx, id)
}

func (svc userService) BlockUserByID(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (svc userService) ChangeUserIsAdmin(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (svc userService) ChangeUserPassword(ctx context.Context, data *service.ChangeUserPasswordData) error {
	hashedPassword, err := hashPassword(data.Password)
	if err != nil {
		return err
	}

	return svc.UserRepository.ChangeUserPassword(ctx, data.UserID, hashedPassword)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func comparePassword(password, hashedPassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("password is incorrect")
	}

	return nil
}
