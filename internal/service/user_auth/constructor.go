package user_auth

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In
	UserAuthTokenRepository repository.UserAuthRepository
}

type userAuthService struct {
	ServiceParams
}

func New(params ServiceParams) service.UserAuthService {
	return &userAuthService{
		ServiceParams: params,
	}
}
