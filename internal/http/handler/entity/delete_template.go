package object

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) DeleteTemplate(c *gin.Context) {
	ctx := parser.Context(c)

	template := context_paylod_parser.GetEntityTemplateFromContext(ctx)
	if template == nil {
		response.Error(c, errors.New("template is nil"))
		return
	}

	err := h.Params.EntityTemplateService.DeleteByIDAndMunicipalityID(ctx, template.ID, template.MunicipalityID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
