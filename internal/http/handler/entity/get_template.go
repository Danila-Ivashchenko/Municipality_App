package entity

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

	template := context_paylod_parser.GetEntityTemplateFromContext(ctx)
	if template == nil {
		response.Error(c, errors.New("template is nil"))
		return
	}

	objectTemplate, err := h.Params.EntityExService.GetByID(ctx, template.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	if objectTemplate == nil {
		response.ResponseNoContent(c)
		return
	}

	v := view.NewEntityTemplateExView(objectTemplate)
	response.Response(c, v)
}

func (h *Handler) GetTemplatesByMunicipality(c *gin.Context) {
	ctx := parser.Context(c)

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("municipality is nil"))
		return
	}

	objectTemplates, err := h.Params.EntityExService.GetByMunicipalityID(ctx, municipality.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	if objectTemplates == nil || len(objectTemplates) == 0 {
		response.ResponseNoContent(c)
		return
	}

	views := make([]view.EntityTemplateExView, 0)

	for _, objectTemplate := range objectTemplates {
		views = append(views, *view.NewEntityTemplateExView(&objectTemplate))
	}

	response.Response(c, views)
}
