package user_admin

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type HandlerParams struct {
	fx.In

	UserAdminService service.UserAdminService
}

type Handler struct {
	Params HandlerParams
}

func New(params HandlerParams) Handler {
	return Handler{
		Params: params,
	}
}
