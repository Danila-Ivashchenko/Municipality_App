package object

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type HandlerParams struct {
	fx.In

	ObjectService         service.ObjectService
	ObjectTemplateService service.ObjectTemplateService
	ObjectExService       service.ObjectExService
}

type Handler struct {
	Params HandlerParams
}

func New(params HandlerParams) Handler {
	return Handler{
		Params: params,
	}
}
