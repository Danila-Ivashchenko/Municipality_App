package municipality_passport

import (
	"go.uber.org/fx"
	"municipality_app/internal/domain/service"
)

type MiddlewareParams struct {
	fx.In

	PassportService       service.PassportService
	ChapterService        service.ChapterService
	PartitionService      service.PartitionService
	ObjectTemplateService service.ObjectTemplateService
	EntityTemplateService service.EntityTemplateService
}

type Middleware struct {
	Params MiddlewareParams
}

func New(params MiddlewareParams) Middleware {
	return Middleware{
		Params: params,
	}
}
