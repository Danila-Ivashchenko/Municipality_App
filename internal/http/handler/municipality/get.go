package municipality

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) GetByID(c *gin.Context) {
	ctx := parser.Context(c)

	municipality := context_paylod_parser.GetMunicipalityFromContext(ctx)
	if municipality == nil {
		response.Error(c, errors.New("mun not found"))
		return
	}

	municipalityEx, err := h.Params.MunicipalityService.GetExById(ctx, municipality.ID)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newMunicipalityExView(municipalityEx)
	response.Response(c, view)
}

func (h *Handler) GetByParams(c *gin.Context) {
	req := &getByParamsRequest{}

	ctx, err := parser.Parse(c, &req)
	if err != nil {
		response.Error(c, err)
		return
	}

	data := req.Convert()

	municipalitiesEx, err := h.Params.MunicipalityService.GetExByParams(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newMunicipalityExViews(municipalitiesEx)
	response.Response(c, view)
}
