package object

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) UpdateTemplate(c *gin.Context) {
	req := &updateObjectTemplateReq{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	if validateErr := req.Validate(); validateErr != nil {
		response.Error(c, validateErr)
		return
	}

	template := context_paylod_parser.GetObjectTemplateFromContext(ctx)
	if template == nil {
		response.Error(c, errors.New("template is nil"))
		return
	}

	data := req.Convert(template.ID, template.MunicipalityID)

	objectTemplate, err := h.Params.ObjectTemplateService.Update(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewObjectTemplateExView(objectTemplate)
	response.Response(c, v)
}
