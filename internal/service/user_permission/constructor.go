package user_permission

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/transactor"
	"municipality_app/internal/domain/repository"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In
	UserPermissionRepository repository.UserPermissionRepository

	Transactor transactor.Transactor
}

type userPermissionService struct {
	ServiceParams
}

func New(params ServiceParams) service.UserPermissionService {
	return &userPermissionService{
		ServiceParams: params,
	}
}
