package auth

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type MiddlewareParams struct {
	fx.In

	UserService           service.UserService
	UserAuthService       service.UserAuthService
	UserPermissionService service.UserPermissionService
}

type Middleware struct {
	Params MiddlewareParams
}

func New(params MiddlewareParams) Middleware {
	return Middleware{
		Params: params,
	}
}
