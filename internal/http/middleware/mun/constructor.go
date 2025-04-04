package mun

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type MiddlewareParams struct {
	fx.In

	MunicipalityService service.MunicipalityService
}

type Middleware struct {
	Params MiddlewareParams
}

func New(params MiddlewareParams) Middleware {
	return Middleware{
		Params: params,
	}
}
