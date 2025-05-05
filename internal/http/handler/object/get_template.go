package object

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) GetTemplateByID(c *gin.Context) {
	ctx := parser.Context(c)

	template := context_paylod_parser.GetObjectTemplateFromContext(ctx)
	if template == nil {
		response.Error(c, errors.New("template is nil"))
		return
	}

	objectTemplate, err := h.Params.ObjectExService.GetByID(ctx, template.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	if objectTemplate == nil {
		response.ResponseNoContent(c)
		return
	}

	v := view.NewObjectTemplateExView(objectTemplate)
	response.Response(c, v)
}

func (h *Handler) GetTemplatesByMunicipality(c *gin.Context) {
	ctx := parser.Context(c)

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("municipality is nil"))
		return
	}

	objectTemplates, err := h.Params.ObjectExService.GetByMunicipalityID(ctx, municipality.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	if objectTemplates == nil || len(objectTemplates) == 0 {
		response.ResponseNoContent(c)
		return
	}

	views := make([]view.ObjectTemplateExView, 0)

	for _, objectTemplate := range objectTemplates {
		views = append(views, *view.NewObjectTemplateExView(&objectTemplate))
	}

	response.Response(c, views)
}
