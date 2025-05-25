package user_admin

import (
	"go.uber.org/fx"
	"municipality_app/internal/common/transactor"
	"municipality_app/internal/domain/service"
)

type ServiceParams struct {
	fx.In

	UserService           service.UserService
	UserPermissionService service.UserPermissionService

	Transactor transactor.Transactor
}

type userAdminService struct {
	ServiceParams
}

func New(params ServiceParams) service.UserAdminService {
	return &userAdminService{
		ServiceParams: params,
	}
}
