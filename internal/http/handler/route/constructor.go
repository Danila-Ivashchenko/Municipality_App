package route

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type HandlerParams struct {
	fx.In

	RouteService service.RouteService
}

type Handler struct {
	Params HandlerParams
}

func New(params HandlerParams) Handler {
	return Handler{
		Params: params,
	}
}
