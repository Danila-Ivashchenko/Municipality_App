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

	template := context_paylod_parser.GetObjectTemplateFromContext(ctx)
	if template == nil {
		response.Error(c, errors.New("template is nil"))
		return
	}

	err := h.Params.ObjectTemplateService.DeleteByIDAndMunicipalityID(ctx, template.ID, template.MunicipalityID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.ResponseNoContent(c)
}
