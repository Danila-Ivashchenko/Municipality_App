package municipality

import (
	"errors"
	"github.com/gin-gonic/gin"
	"municipality_app/internal/common/context_paylod_parser"
	"municipality_app/internal/http/utils/parser"
	"municipality_app/internal/http/utils/response"
)

func (h *Handler) Update(c *gin.Context) {
	req := &updateMunicipalityRequest{}

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
		response.Error(c, errors.New("municipality not found"))
		return
	}

	data := req.Convert(municipality.ID)

	municipalityEx, err := h.Params.MunicipalityService.Update(ctx, data)
	if err != nil {
		response.Error(c, err)
		return
	}

	view := newMunicipalityExView(municipalityEx)
	response.Response(c, view)
}
