package object

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
	"municipality_app/internal/http/view"
)

func (h *Handler) CreateTemplate(c *gin.Context) {
	req := &createEntityTemplateReq{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	if validateErr := req.Validate(); validateErr != nil {
		response.Error(c, validateErr)
		return
	}

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("municipality is nil"))
		return
	}

	data := req.Convert(municipality.ID)

	objectTemplate, err := h.Params.EntityTemplateService.Create(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	v := view.NewEntityTemplateExView(objectTemplate)
	response.Response(c, v)
}
