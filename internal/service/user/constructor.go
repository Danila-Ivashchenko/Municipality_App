package user

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	UserAuthService service.UserAuthService
	UserRepository  repository.UserRepository
}

type userService struct {
	ServiceParams
}

func New(params ServiceParams) service.UserService {
	return &userService{
		ServiceParams: params,
	}
}
