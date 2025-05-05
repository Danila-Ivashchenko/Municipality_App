package entity

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type HandlerParams struct {
	fx.In

	EntityService         service.EntityService
	EntityTemplateService service.EntityTemplateService
	EntityExService       service.EntityExService
}

type Handler struct {
	Params HandlerParams
}

func New(params HandlerParams) Handler {
	return Handler{
		Params: params,
	}
}
