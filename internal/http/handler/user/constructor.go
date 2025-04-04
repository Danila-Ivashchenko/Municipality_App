package user

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type HandlerParams struct {
	fx.In

	UserService service.UserService
}

type Handler struct {
	Params HandlerParams
}

func New(params HandlerParams) Handler {
	return Handler{
		Params: params,
	}
}
